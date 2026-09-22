---
name: register-go-route
description: 対象モデルのHandlerをGinルーターへ登録する。
---

# Router登録

対象リソース`<RESOURCE>`、Handler、HTTPメソッド、パスを指定する。`backend/config/router.go`の引数と既存ルートを確認する。

- 既存の複数形・スネークケース規約に合わせる
- 固定パスと`/:id`の競合を確認する
- Handlerを正しいServiceへ接続する
- 要求されていないルートを追加しない
- `Routing`のシグネチャ変更が必要なら呼び出し元も確認する
