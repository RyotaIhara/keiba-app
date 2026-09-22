---
name: create-go-schema
description: 対象モデルを指定してMySQLのCreateTable SQLを作成する。
---

# CreateTable作成

実行前に`../references/common-rules.md`と、リポジトリルート配下の関連する`docs/*.md`を確認する。

対象モデル`<MODEL>`と、調査済みの`<TABLE>`を指定する。`../references/common-rules.md`とモデル・既存DDLを確認し、`sql/schema/create_<table>.sql`を作成する。

- `CREATE TABLE IF NOT EXISTS`を使う
- MySQL 8.4、InnoDB、既存のutf8mb4設定に合わせる
- モデルの型、NULL可否、enumの表現を反映する
- 外部キーの参照先、削除時の挙動、INDEXを確認する
- 既存DDLの作成順と矛盾させない

既存ファイルがある場合は内容を確認して統合し、無条件に上書きしない。DDLのみを変更し、Goコードは変更しない。
