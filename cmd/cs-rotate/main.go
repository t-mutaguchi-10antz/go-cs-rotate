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
	Verbose  bool   `short:"v" long:"verbose" description:"verbose"`
	Quantity uint   `short:"q" long:"quantity" description:"quantity" required:"true" validate:"gt=0"`
	Order    string `short:"o" long:"order" default:"desc" description:"order"`
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

	if err := usecase.RotateStorage(ctx, args.Order, args.Quantity); err != nil {
		log.Fatal(err)
	}
}
