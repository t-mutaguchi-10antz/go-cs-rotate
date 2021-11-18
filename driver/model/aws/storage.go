package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/davecgh/go-spew/spew"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
)

var _ model.Storage = storage{}

type storage struct {
	verbose bool
	client  *s3.Client
}

func NewStorage(ctx context.Context, verbose bool) (model.Storage, error) {
	s := storage{verbose: verbose}

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-2"))
	if err != nil {
		return s, fmt.Errorf("failed to create storage struct: %w", err)
	}

	s.client = s3.NewFromConfig(cfg)

	return s, nil
}

func (s storage) List(ctx context.Context, options ...model.ListOption) ([]model.Object, []model.ListOption, error) {
	o := model.NewListOptions(options...)

	if s.verbose {
		log.Printf("AWS S3 List objects %s", spew.Sdump(o))
	}

	bucket := o.URL.Bucket()
	prefix := o.URL.Prefix()
	params := &s3.ListObjectsInput{
		Bucket: &bucket,
		Prefix: &prefix,
	}
	optFns := []func(*s3.Options){}
	output, err := s.client.ListObjects(ctx, params, optFns...)
	if err != nil {
		return []model.Object{}, nil, fmt.Errorf("failed to list AWS S3 objects: %w", err)
	}

	objects := []model.Object{}
	for _, content := range output.Contents {
		url := ""
		if output.Prefix != nil {
			url = fmt.Sprintf("s3://%s/%s/%s", *output.Name, *output.Prefix, *content.Key)
		} else {
			url = fmt.Sprintf("s3://%s/%s", *output.Name, *content.Key)
		}
		object, err := model.NewObject(url)
		if err != nil {
			return []model.Object{}, nil, fmt.Errorf("failed to list AWS S3 objects: %w", err)
		}
		objects = append(objects, object)
	}

	return objects, nil, nil
}

func (s storage) Delete(ctx context.Context, objects []model.Object) error {
	// s.client.DeleteObjects(ctx)

	return nil
}
