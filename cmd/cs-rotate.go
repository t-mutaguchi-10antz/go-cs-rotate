package main

import (
	"context"
	"log"
	"os"

	"github.com/jessevdk/go-flags"

	domain "github.com/t-mutaguchi-10antz/cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/domain/usecase"
	driver "github.com/t-mutaguchi-10antz/cs-rotate/driver/model"
	"github.com/t-mutaguchi-10antz/cs-rotate/validator"
)

var args struct {
	Platform   string
	URL        string `short:"u" long:"url" required:"true" validate:"url" description:"削除対象の始点となる URL ( protocol://bucket/prefix )"`
	Quantity   uint   `short:"q" long:"quantity" required:"true" validate:"gt=0" description:"ローテートせずに残す量"`
	Order      string `short:"o" long:"order" choice:"desc" choice:"asc" default:"desc" description:"ローテートせずに残すにあたって降順・昇順どちらで並べ替えるか"`
	Verbose    bool   `short:"v" long:"verbose" description:"詳細ログを出力するか"`
	AWSProfile string `long:"aws-profile" description:"ストレージが AWS S3 の場合はプロファイルを指定する"`
}

func init() {
	// コマンドラインからの入力を解析・検証する
	args.Platform = os.Args[1]
	if _, err := domain.WithPlatform(args.Platform); err != nil {
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

	// プラットフォームに合ったストレージ構造体を生成する
	options := []driver.Option{}
	if args.AWSProfile != "" {
		options = append(options, driver.WithAWSProfile(args.AWSProfile))
	}
	storage, err := driver.NewStorage(ctx, args.Verbose, args.Platform, options...)
	if err != nil {
		log.Fatal(err)
	}

	// クラウドストレージ上のリソースを条件に合わせて削除する
	model := domain.NewModel(storage)
	usecase := usecase.NewUsecase(model)
	if err := usecase.RotateStorage(ctx, args.URL, args.Quantity, args.Order); err != nil {
		log.Fatal(err)
	}
}
