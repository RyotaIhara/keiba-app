---
name: create-go-service
description: 対象モデルを指定してinfrastructureを呼び出すServiceユースケースを作成する。
---

# Service作成

実行前に`../references/common-rules.md`と、リポジトリルート配下の関連する`docs/*.md`を確認する。

対象モデル`<MODEL>`と既存または追加済みStoreを指定する。`backend/service/<domain>/`の命名と構成を確認する。

- 一覧・単件など要求されたユースケースだけを公開する
- ServiceからStoreを呼び出す
- エラーを握りつぶさない
- 既存モデルをそのまま返し、不要な変換を重複させない
- Applicationから生成できるコンストラクタを用意する

InfrastructureにSQLを追加したり、HandlerのHTTP処理を混ぜたりしない。
