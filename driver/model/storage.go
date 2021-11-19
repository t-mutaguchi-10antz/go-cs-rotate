package model

import (
	"context"
	"fmt"

	"github.com/t-mutaguchi-10antz/go/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/go/cs-rotate/driver/model/aws"
)

func NewStorage(ctx context.Context, verbose bool, platform string, options ...Option) (model.Storage, error) {
	var storage model.Storage

	p, err := model.WithPlatform(platform)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage struct: %w", err)
	}

	o := NewOptions(options...)

	switch p {
	case model.PlatformAWS:
		if o.AWSProfile == nil {
			return nil, fmt.Errorf("failed to create storage struct: AWS profile is must needed")
		}
		storage, err = aws.NewStorage(ctx, verbose, *o.AWSProfile)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create storage struct: %w", err)
	}

	return storage, nil
}

type Option interface {
	Apply(options *Options)
}

type Options struct {
	AWSProfile *string
}

func NewOptions(options ...Option) Options {
	opts := Options{}
	for _, option := range options {
		if option == nil {
			continue
		}
		option.Apply(&opts)
	}
	return opts
}

type awsProfile string

func (a awsProfile) Apply(options *Options) {
	v := string(a)
	options.AWSProfile = &v
}

func WithAWSProfile(v string) awsProfile {
	return awsProfile(v)
}
