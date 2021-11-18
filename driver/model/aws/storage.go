package aws

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
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

func (s storage) List(ctx context.Context, options ...model.ListOption) ([]model.Object, []model.ListOption, error) {
	objects := []model.Object{}

	o := model.NewListOptions(options...)

	bucket := o.URL.Bucket()
	prefix := o.URL.Prefix()
	delimiter := "/"
	params := &s3.ListObjectsV2Input{
		Bucket:    &bucket,
		Prefix:    &prefix,
		Delimiter: &delimiter,
	}

	prefixes := []string{}
	p := s3.NewListObjectsV2Paginator(s.client, params)
	for p.HasMorePages() {
		page, err := p.NextPage(ctx)
		if err != nil {
			return objects, nil, fmt.Errorf("failed to get a page: %w", err)
		}
		for _, p := range page.CommonPrefixes {
			prefixes = append(prefixes, *p.Prefix)
		}
	}

	sort.SliceStable(prefixes, func(i, j int) bool {
		if o.Order != nil && *o.Order == model.OrderAsc {
			return prefixes[i] < prefixes[j]
		} else {
			return prefixes[i] > prefixes[j]
		}
	})

	if s.verbose {
		log.Printf("AWS S3 List prefixes ( bucket: %s, prefix: %s ) %v", bucket, prefix, prefixes)
	}

	return objects, nil, nil
}

func (s storage) Delete(ctx context.Context, objects []model.Object) error {
	// s.client.DeleteObjects(ctx)

	return nil
}
