# Backendテスト実装ルール

本ドキュメントでは、Backendのhandler、service、infrastructure、およびmodel・helperに対するテストコードの配置と実装ルールを定めます。

## テストファイルの配置

テストコードは実装パッケージの直下ではなく、`backend/test`配下に責務別で配置します。

```text
backend/test/
  handler/
    user/
    race/
    race_course/
  service/
    user/
    race/
    race_course/
    race_search/
  infrastructure/
    user/
    race/
    race_course/
```

handler、service、infrastructureのテストファイルは、それぞれ対応する実装パッケージをimportして利用します。テストパッケージは実装パッケージと分離し、公開APIを通じて動作を検証します。

## モックの利用ルール

### handler

handlerテストでは、handlerが依存するServiceをモックします。

- 成功時のレスポンスとHTTPステータスを検証する
- 入力値がServiceへ渡されることを検証する
- 入力不正、`sql.ErrNoRows`、Serviceエラーを検証する
- GinのContextにエラーが記録される処理は、記録内容も検証する

### service

serviceテストでは、infrastructureのStoreをモックします。

- 各Serviceメソッドが対応するStoreメソッドを呼び出すことを検証する
- 作成・更新後に対象モデルを再取得する処理を検証する
- StoreのエラーをServiceがそのまま返すことを検証する
- modelの値は実際のmodelオブジェクトを使用する

Serviceがモック可能になるよう、Storeは必要なメソッドを定義したインターフェース経由で受け取ります。本番コードでは既存のStore実装をそのまま渡します。

### infrastructure

infrastructureテストでは、実DBへ接続せず`github.com/DATA-DOG/go-sqlmock`を使用します。

- SQL文が実行されることを検証する
- SQLの引数が期待どおりであることを検証する
- DBの結果が正しいmodelへ変換されることを検証する
- Query、QueryRow、Execのエラーを検証する
- `sql.ErrNoRows`を返すケースを検証する
- 更新・削除対象が存在しない場合の存在確認処理を検証する
- `rows.Close`および`rows.Err`に関係する失敗経路を必要に応じて検証する

テストでは実際のmodel、`RaceInput`、`RaceSearchInput`などの入力オブジェクトを使用し、テスト専用の簡易構造体で代替しません。

## model・helperのテスト

modelやhelperの値をテストデータとして生成する必要がある場合は、実際の型・実際のオブジェクトを使用します。

- modelのフィールドには実装で定義された型を使用する
- enumは定義済みの定数を使用する
- 日付・時刻は`time.Time`を使用する
- helperの入力・出力は実際のGin Contextや実際のリクエスト形式を使用する

## テスト対象

既存のpublicメソッドは、原則として成功と失敗の両方をテストします。

- handler: 一覧、単件、作成、更新、削除、検索、入力検証
- service: 一覧、単件、作成、更新、削除、検索
- infrastructure: 一覧、単件、検索、詳細取得、作成、更新、削除

privateメソッドは、外部から直接呼び出すテストを追加するのではなく、publicメソッドを通じて結果を検証します。入力変換やscan処理も、対応するpublicメソッドの結果で検証します。

## 検証コマンド

Backendの変更後は、`backend`ディレクトリで次のコマンドを実行します。

```sh
gofmt -w test
go test ./...
go test -race ./...
go build ./...
```

テストは実DBや外部サービスに依存せず、ローカル環境で再現可能な状態にします。
