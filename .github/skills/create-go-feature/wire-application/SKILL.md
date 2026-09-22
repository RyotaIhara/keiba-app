---
name: wire-go-application
description: 対象モデルのStoreとServiceをapplication.goのComposition Rootへ登録する。
---

# Application依存性注入

対象モデル`<MODEL>`、Store、Serviceを指定する。`backend/application/application.go`と`backend/main.go`を確認する。

- DBからStoreを生成する
- StoreからServiceを生成する
- `Application`構造体へServiceを追加する
- `application.New`の戻り値へ設定する
- Routerへ渡す必要がある場合は呼び出し元も更新する
- 既存の依存性を壊さない

Application層へSQLやHTTP処理を追加しない。
