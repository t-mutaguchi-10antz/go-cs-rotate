package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
)

var _ model.Storage = storage{}

type storage struct {
	client *s3.Client
}

func NewStorage() (model.Storage, error) {
	s := storage{}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		return s, fmt.Errorf("Failed to create storage struct: %w", err)
	}

	s.client = s3.NewFromConfig(cfg)

	return s, nil
}

func (s storage) List(ctx context.Context) ([]model.Resource, error) {
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
