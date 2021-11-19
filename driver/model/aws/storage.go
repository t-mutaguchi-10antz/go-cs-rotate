package aws

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/t-mutaguchi-10antz/go/cs-rotate/domain/model"
)

var _ model.Storage = storage{}

type storage struct {
	verbose bool
	client  *s3.Client
}

func NewStorage(ctx context.Context, verbose bool, profile string) (model.Storage, error) {
	s := storage{verbose: verbose}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	if err != nil {
		return s, fmt.Errorf("failed to create storage struct: %w", err)
	}

	s.client = s3.NewFromConfig(cfg)

	return s, nil
}

func (s storage) Rotate(ctx context.Context, params ...model.RotateParam) error {
	p := model.NewRotateParams(params...)

	if s.verbose {
		log.Printf("AWS S3 Rotate ( %s )", p.URL)
	}

	deleted, _, err := s.prefixes(ctx, p)
	if err != nil {
		return fmt.Errorf("failed to rotate: %w", err)
	}

	bucket, err := p.URL.Bucket()
	if err != nil {
		return fmt.Errorf("failed to rotate: %w", err)
	}
	if err := s.delete(ctx, bucket, deleted); err != nil {
		return fmt.Errorf("failed to rotate: %w", err)
	}

	return nil
}

func (s storage) prefixes(ctx context.Context, params model.RotateParams) (deleted []string, kept []string, err error) {
	bucket, err := params.URL.Bucket()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get prefixes: %w", err)
	}
	prefix, err := params.URL.Prefix()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get prefixes: %w", err)
	}

	prefixes := []string{}
	if err := s.paginate(ctx, bucket, prefix, "/", func(page *s3.ListObjectsV2Output) error {
		for _, c := range page.CommonPrefixes {
			prefixes = append(prefixes, *c.Prefix)
		}
		return nil
	}); err != nil {
		return nil, nil, fmt.Errorf("failed to get prefixes: %w", err)
	}

	sort.SliceStable(prefixes, func(i, j int) bool {
		if params.Order == model.OrderAsc {
			return prefixes[i] < prefixes[j]
		} else {
			return prefixes[i] > prefixes[j]
		}
	})

	qty := int(params.Quantity)
	len := len(prefixes)
	if len < qty {
		qty = len
	}
	deleted = prefixes[qty:]
	kept = prefixes[0:qty]

	if s.verbose {
		log.Printf("AWS S3 Keep prefixes ( %v )", kept)
	}

	return deleted, kept, nil
}

func (s storage) delete(ctx context.Context, bucket string, prefixes []string) error {
	for _, prefix := range prefixes {
		if err := s.paginate(ctx, bucket, prefix, "", func(page *s3.ListObjectsV2Output) error {
			objects := []types.ObjectIdentifier{}
			for _, content := range page.Contents {
				if s.verbose {
					log.Printf("AWS S3 Delete ( %s )", *content.Key)
				}
				objects = append(objects, types.ObjectIdentifier{
					Key: content.Key,
				})
			}
			_, err := s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
				Bucket: &bucket,
				Delete: &types.Delete{
					Objects: objects,
				},
			})
			if err != nil {
				return fmt.Errorf("failed to delete: %w", err)
			}

			return nil
		}); err != nil {
			return fmt.Errorf("failed to delete: %w", err)
		}
	}

	return nil
}

func (s storage) paginate(
	ctx context.Context,
	bucket string,
	prefix string,
	delimiter string,
	pageFunc func(page *s3.ListObjectsV2Output) error,
) error {
	if s.verbose {
		log.Printf("AWS S3 Paginate ( Bucket: %s, Prefix: %s )", bucket, prefix)
	}

	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket:    &bucket,
		Prefix:    &prefix,
		Delimiter: &delimiter,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to get a page: %w", err)
		}
		if err := pageFunc(page); err != nil {
			return fmt.Errorf("failed to process a page: %w", err)
		}
	}

	return nil
}
