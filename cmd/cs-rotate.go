package main

import (
	"context"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	domain_model "github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/usecase"
	driver_model "github.com/t-mutaguchi-10antz/cs-rotate/driver/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

var args struct {
	Platform string
	URL      string `short:"u" long:"url" required:"true" validate:"url" description:"削除対象の始点となる URL ( protocol://bucket/prefix )"`
	Quantity uint   `short:"q" long:"quantity" required:"true" validate:"gt=0" description:"ローテートを行わない確保量"`
	Order    string `short:"o" long:"order" choice:"desc" choice:"asc" default:"desc" description:"確保するにあたって降順か昇順どちらで並べ替えるか"`
	Verbose  bool   `short:"v" long:"verbose" description:"詳細ログを出力するか"`
}

func init() {
	// コマンドラインからの入力を解析・検証する
	args.Platform = os.Args[1]
	if _, err := domain_model.Platform(args.Platform); err != nil {
		log.Fatal(err)
	}
	if _, err := flags.Parse(&args); err != nil {
		log.Fatal(err)
	}
	if err := validator.CheckStruct(&args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	// プラットフォーム用のストレージ構造体を生成する
	storage, err := driver_model.NewStorage(ctx, args.Verbose, args.Platform)
	if err != nil {
		log.Fatal(err)
	}

	// クラウドストレージ上のリソースを条件に合わせて削除する
	model := domain_model.NewModel(storage)
	usecase := usecase.NewUsecase(model)
	if err := usecase.RotateCloudStorage(ctx, args.URL, args.Quantity, args.Order); err != nil {
		log.Fatal(err)
	}
}
