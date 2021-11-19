# Cloud Storage Rotate

[![Lint & UnitTest](https://github.com/t-mutaguchi-10antz/go-cs-rotate/actions/workflows/lint-and-unittest.yaml/badge.svg)](https://github.com/t-mutaguchi-10antz/go-cs-rotate/actions/workflows/lint-and-unittest.yaml)

## Descriptions

条件に合わせてクラウドストレージ上のリソースを削除する

## Installation

```bash
$ go install github.com/t-mutaguchi-10antz/go-cs-rotate@latest
```

## Usage

```
Usage:
  cs-rotate PLATFORM[aws] [OPTIONS]

Application Options:
  -u, --url=             削除対象の基点となる URL ( protocol://bucket/prefix )
  -q, --quantity=        ローテートせずに残す量
  -o, --order=[desc|asc] ローテートせずに残すにあたって降順・昇順どちらで並べ替えるか
  -v, --verbose          詳細ログを出力するか
      --aws-profile=     ストレージが AWS S3 の場合は必須, プロファイルを指定する
```