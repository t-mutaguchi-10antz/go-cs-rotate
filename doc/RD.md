# Requirements Definition

## Goal

- クラウドストレージ上のリソースを自然数を元に ( 小さい or 大きい ) ソートした上で上位 n 件以外を削除する処理
- 当初は AWS のみ対応, 将来的には GCP も対応

```bash
# プロファイルで指定できそうなら
$ cs-rotate aws natural-number --quantity 5 --order desc --path s3://${BUCKET}/${PREFIX} --aws-profile ${AWS_PROFILE_NAME}

# プロファイルでの指定が難しそうなら
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=xxx REGION=ap-northeast-1 cs-rotate aws natural-number --quantity 5 --order desc --path s3://${BUCKET}/${PREFIX}
```
