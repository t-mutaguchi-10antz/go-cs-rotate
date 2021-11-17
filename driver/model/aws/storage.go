package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/davecgh/go-spew/spew"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/primitive"
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
		return s, fmt.Errorf("Failed to create storage struct: %w", err)
	}

	s.client = s3.NewFromConfig(cfg)

	return s, nil
}

func (s storage) List(ctx context.Context, options ...primitive.ListOption) ([]model.Resource, error) {
	o := primitive.NewListOptions(options...)

	if s.verbose {
		log.Printf("AWS S3 List resources %s", spew.Sdump(o))
	}

	params := &s3.ListObjectsInput{}
	optFns := []func(*s3.Options){}
	output, err := s.client.ListObjects(ctx, params, optFns...)
	if err != nil {
		return []model.Resource{}, fmt.Errorf("Failed to list AWS S3 resources: %w", err)
	}

	resources := []model.Resource{}
	for _, content := range output.Contents {
		url := ""
		if output.Prefix != nil {
			url = fmt.Sprintf("s3://%s/%s/%s", *output.Name, *output.Prefix, *content.Key)
		} else {
			url = fmt.Sprintf("s3://%s/%s", *output.Name, *content.Key)
		}
		resource, err := model.NewResource(url)
		if err != nil {
			return []model.Resource{}, fmt.Errorf("Failed to list AWS S3 resources: %w", err)
		}
		resources = append(resources, resource)
	}

	return resources, nil
}

func (s storage) Delete(ctx context.Context, resources []model.Resource) error {
	// s.client.DeleteObjects(ctx)

	return nil
}
