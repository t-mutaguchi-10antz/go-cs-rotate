# Requirements Definition

- クラウドストレージ上の指定したパス直下を名前順でソートした上で上位 n 件以外を削除する
- 当初は AWS のみ対応, 将来的には GCP も対応

```bash
# プロファイルで指定できそうなら
$ cmd aws natural-number --quantity 5 --order desc --path s3://${BUCKET}/${PREFIX} --aws-profile ${AWS_PROFILE_NAME}

# プロファイルでの指定が難しそうなら
$ AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=xxx REGION=ap-northeast-1 cmd aws natural-number --quantity 5 --order desc --path s3://${BUCKET}/${PREFIX}
```
