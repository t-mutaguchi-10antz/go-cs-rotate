package main

import (
	"log"

	"github.com/jessevdk/go-flags"

	"github.com/t-mutaguchi-10antz/cs-rotate/adapter/model/aws"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/usecase"
	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

var args struct {
	Verbose  bool   `short:"v" long:"verbose" description:"verbose"`
	Quantity uint   `short:"q" long:"quantity" description:"quantity" required:"true" validate:"gt=0"`
	Order    string `short:"o" long:"order" description:"order"`
}

func init() {
	if _, err := flags.Parse(&args); err != nil {
		log.Fatal(err)
	}

	if err := validator.Check(&args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	storage, err := aws.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	model := model.NewModel(storage)
	usecase := usecase.NewUsecase(model)

	if err := usecase.RotateStorage(); err != nil {
		log.Fatal(err)
	}
}
