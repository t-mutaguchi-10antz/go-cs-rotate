# Requirements Definition

## Goal

自然数を元に順番 ( 小さい or 大きい ) でソートした上で上位 n 件を残す処理

```bash
# プロファイルで指定できそうなら
$ s3-rotate natural-number --quantity 5 --order desc --path s3://${BUCKET}/${PREFIX} --profile ${AWS_PROFILE_NAME}

# プロファイルでの指定が難しそうなら
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=xxx REGION=ap-northeast-1 s3-rotate natural-number --quantity 5 --order desc --path s3://${BUCKET}/${PREFIX}
```
