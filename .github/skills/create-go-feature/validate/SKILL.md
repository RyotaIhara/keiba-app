---
name: validate-go-feature
description: 対象モデルの変更をgofmt、Goテスト、ビルドで検証する。
---

# Go機能の検証

対象モデル`<MODEL>`と変更ファイルを指定する。まず`git diff --check`で空白エラーを確認し、変更したGoファイルをフォーマットする。

```sh
gofmt -w <changed-go-files>
cd backend
go test ./...
go build ./...
```

次を確認する。

- DDLの外部キーと型がモデル・既存DDLに一致する
- 一覧取得、単件取得、存在しないID、不正なIDの挙動
- Application生成からRouter登録までコンパイルできる
- 既存APIの回帰がない

失敗した場合は原因を修正して再実行し、実行できなかった検証は成功として報告しない。
