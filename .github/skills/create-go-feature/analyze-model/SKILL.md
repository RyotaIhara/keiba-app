---
name: analyze-go-model
description: 対象Goモデル、関連モデル、DDL、enum、既存実装を調査し、後続工程に必要な仕様を確定する。
---

# モデル調査

実行前に`references/common-rules.md`と、リポジトリルート配下の関連する`docs/*.md`を確認する。

対象モデルを`<MODEL>`として指定する。`references/common-rules.md`を読み、次を調査する。

- `backend/model/**/<model>.go`
- 関連する親モデル、値オブジェクト、enum
- `sql/schema/create_*.sql`
- 同じドメインのInfrastructure、Service、Handler

次を成果物として簡潔に整理する。

- Goパッケージ、テーブル名、リソース名
- カラムとGo型、NULL可否
- 主キー、外部キー、UNIQUE、INDEX
- 必要なJOIN、取得条件、並び順
- 想定するHTTPメソッドとパス

ファイル変更は原則行わない。テーブル名、API仕様、JOINが一意に決まらない場合は確認を求める。
