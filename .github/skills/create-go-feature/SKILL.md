---
name: create-go-feature
description: >
  対象モデルを指定して、Go/Gin機能の各工程Skillを順番に適用するための案内。
  個別工程だけを実行したい場合は、配下の各Skillを直接使用する。
---

# Go機能追加Skill群

このディレクトリでは、対象モデルを指定して各工程を個別に実行する。

```text
.github/skills/create-go-feature/
├── SKILL.md
├── references/common-rules.md
├── analyze-model/SKILL.md
├── create-schema/SKILL.md
├── create-infrastructure-query/SKILL.md
├── create-service/SKILL.md
├── create-handler/SKILL.md
├── register-route/SKILL.md
├── wire-application/SKILL.md
└── validate/SKILL.md
```

各Skillの依頼には、対象モデルを明示する。

```text
RaceDetailモデルを対象に、Create Schema Skillを実行してください。
```

## 実行前の必須確認

すべての工程で、最初に次のファイルを確認する。

1. `.github/skills/create-go-feature/references/common-rules.md`
2. 対象に関連するリポジトリルートの`docs/*.md`（サブディレクトリを含む）
3. 対象工程のSkillファイル

`docs/`の規約は、コードの配置、命名、コメント、テスト、アーキテクチャ、検証方法に
反映する。docsと既存コードまたは依頼内容が競合し、仕様を一意に決められない場合は、
推測で実装せず、確認が取れるまで後続工程へ進まない。

一連の処理は次の順序で実行する。

1. `analyze-model`
2. `create-schema`
3. `create-infrastructure-query`
4. `create-service`
5. `create-handler`
6. `register-route`
7. `wire-application`
8. `validate`

各工程は前工程の成果と既存コードを確認してから実行する。仕様が未確定なら後続工程へ進まない。
