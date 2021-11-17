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
	storage := aws.NewStorage()
	model := model.NewModel(storage)

	usecase := usecase.NewUsecase(model)
	usecase.RotateStorage()
}
