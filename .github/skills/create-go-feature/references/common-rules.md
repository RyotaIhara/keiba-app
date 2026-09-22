# 共通ルール

## 入力

依頼から次の値を確定する。`<MODEL>`は必須の対象モデル名で、例は`RaceDetail`です。

- `MODEL`: Goモデル名
- `PACKAGE`: モデルのGoパッケージ
- `TABLE`: DBテーブル名
- `RESOURCE`: APIリソース名
- 必要な場合の親ID、取得条件、JOIN対象、並び順

モデル名しか指定されていない場合は、モデル、関連DDL、関連モデル、既存の同種実装を調査する。テーブル名やAPI仕様が一意に決まらない場合は推測せず確認する。

## 調査対象

実装前に次を確認する。

1. `backend/model/**/<model>.go`
2. 関連モデルとenum
3. `sql/schema/create_*.sql`
4. 同じ責務の`backend/infrastructure/**`
5. 同じ責務の`backend/service/**`
6. 同じAPI形態の`backend/handler/**`
7. `backend/config/router.go`
8. `backend/application/application.go`
9. `backend/main.go`と既存テスト

## 既存規約

- モデルの型、enum、ポインタ関係とDB列の表現を一致させる
- SQLのパラメータはプレースホルダを使う
- `Query`後の`Close`と`rows.Err()`を処理する
- 一覧・単件で共有できるScan処理は専用関数へ切り出す
- `sql.ErrNoRows`はHandlerで404へ変換する
- その他のエラーは握りつぶさず、既存の`c.Error(err)`と固定メッセージの形式に合わせる
- 依頼されていないCRUDやルートを追加しない
- 既存の未コミット変更を上書きしない

## 検証

変更したGoファイルに対して`gofmt`を実行し、`backend`で次を実行する。

```sh
go test ./...
go build ./...
```

実行できなかった検証は成功として報告しない。
