---
name: create-go-handler
description: 対象モデルを指定してGinの取得Handlerを作成する。
---

# Handler作成

対象モデル`<MODEL>`、Service、HTTPメソッド、パスを指定する。既存Handlerのエラー・JSON形式を確認する。

- `gin.HandlerFunc`を返す既存パターンを使う
- パスパラメータを数値・正数として検証する
- `sql.ErrNoRows`を404へ変換する
- その他のエラーは`c.Error(err)`後に既存形式の500を返す
- 成功時のステータスとJSON形状を明示する
- HandlerにSQLを書かない

指定されていない作成・更新・削除Handlerは追加しない。
