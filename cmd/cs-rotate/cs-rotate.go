package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/jessevdk/go-flags"

	domain "github.com/t-mutaguchi-10antz/go-cs-rotate/domain/model"
	"github.com/t-mutaguchi-10antz/go-cs-rotate/domain/usecase"
	driver "github.com/t-mutaguchi-10antz/go-cs-rotate/driver/model"
	"github.com/t-mutaguchi-10antz/go-cs-rotate/validator"
)

var Version string

var args struct {
	Platform   string
	URL        string `short:"u" long:"url" required:"true" validate:"url" description:"削除対象の基点となる URL ( protocol://bucket/prefix )"`
	Quantity   uint   `short:"q" long:"quantity" required:"true" validate:"gt=0" description:"ローテートせずに残す量"`
	Order      string `short:"o" long:"order" choice:"desc" choice:"asc" default:"desc" description:"ローテートせずに残すにあたって降順・昇順どちらで並べ替えるか"`
	Verbose    bool   `short:"v" long:"verbose" description:"詳細ログを出力するか"`
	AWSProfile string `long:"aws-profile" description:"ストレージが AWS S3 の場合は必須, プロファイルを指定する"`
	Version    bool   `long:"version" description:"バージョン"`
}

func init() {
	// 引数が無ければヘルプを表示する
	parser := flags.NewParser(&args, flags.HelpFlag)
	parser.Name = "cs-rotate"
	parser.Usage = "PLATFORM[aws] [OPTIONS]"
	if len(os.Args) == 1 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	// 引数を解析する
	if _, err := parser.Parse(); err != nil {
		if args.Version {
			// for "go build" ( go build -ldflags="-X main.Version=$(git describe --tags)" )
			if Version != "" {
				fmt.Println(Version)
				os.Exit(0)
			}

			// for "go install"
			if buildInfo, ok := debug.ReadBuildInfo(); ok {
				fmt.Println(buildInfo.Main.Version)
				os.Exit(0)
			}
		}
		log.Fatal(err)
	}

	// プラットフォームを解析する
	args.Platform = os.Args[1]
	if _, err := domain.WithPlatform(args.Platform); err != nil {
		log.Printf("invalid platform: %s", args.Platform)
		os.Exit(1)
	}

	// 引数を検証する
	if err := validator.CheckStruct(&args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()

	// プラットフォームに合ったストレージ構造体を生成する
	opts := []driver.Option{}
	if args.AWSProfile != "" {
		opts = append(opts, driver.WithAWSProfile(args.AWSProfile))
	}
	storage, err := driver.NewStorage(ctx, args.Verbose, args.Platform, opts...)
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
