package main

import (
	"context"
	"log"

	"github.com/jessevdk/go-flags"

	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/usecase"
	"github.com/t-mutaguchi-10antz/cs-rotate/driver/model/aws"
	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

var args struct {
	URL      string `short:"u" long:"url" description:"url" required:"true" validate:"url"`
	Quantity uint   `short:"q" long:"quantity" description:"quantity" required:"true" validate:"gt=0"`
	Order    string `short:"o" long:"order" default:"desc" description:"order"`
	Verbose  bool   `short:"v" long:"verbose" description:"verbose"`
}

func init() {
	if _, err := flags.Parse(&args); err != nil {
		log.Fatal(err)
	}

	if err := validator.CheckStruct(&args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	storage, err := aws.NewStorage(ctx, args.Verbose)
	if err != nil {
		log.Fatal(err)
	}

	model := model.NewModel(storage)
	usecase := usecase.NewUsecase(model)

	if err := usecase.RotateStorage(ctx, args.URL, args.Quantity, args.Order); err != nil {
		log.Fatal(err)
	}
}
