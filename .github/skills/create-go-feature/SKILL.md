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
