package main

import (
	"log"

	"github.com/jessevdk/go-flags"

	"github.com/t-mutaguchi-10antz/cs-rotate/adapter/model/aws"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/usecase"
)

var args struct {
	Quantity uint   `short:"q" long:"quantity" description:"quantity" required:"true"`
	Order    string `short:"o" long:"order" description:"order"`
}

func init() {
	_, err := flags.Parse(&args)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	storage := aws.NewStorage()
	model := model.NewModel(storage)

	usecase := usecase.NewUsecase(model)
	usecase.RotateStorage()
}
