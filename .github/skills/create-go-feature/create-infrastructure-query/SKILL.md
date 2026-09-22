---
name: create-go-infrastructure-query
description: 対象モデルを指定してinfrastructureのStore、取得SQL、Scan処理を作成する。
---

# Infrastructure取得処理

実行前に`../references/common-rules.md`と、リポジトリルート配下の関連する`docs/*.md`を確認する。

対象モデル`<MODEL>`、テーブル`<TABLE>`、必要な取得条件を指定する。共通ルールと既存Storeを確認し、既存責務に合う`backend/infrastructure/<resource>/`へ追加する。

- `NewStore(db *sql.DB)`と既存Storeの構成を再利用する
- 一覧は`Query`、単件は`QueryRow`を使う
- SQL列順とScan順を一致させる
- 一覧と単件で共有できるScan処理を切り出す
- `rows.Close()`、`rows.Err()`、DBエラーを処理する
- JOINした関連モデルのポインタを初期化する
- パラメータはプレースホルダを使い、ORDER BYを明示する

HTTP型やステータスコードは追加しない。必要な取得メソッドだけを実装する。
