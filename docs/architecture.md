# アーキテクチャ

## 概要

本リポジトリは、競馬情報を扱うフロントエンドとバックエンドを分離した構成です。

- **フロントエンド**: Vue 3 + TypeScript + Vite
- **バックエンド**: Go + Gin
- **データベース**: MySQL 8.4
- **ローカル実行**: Docker Compose
- **Kubernetes 実行**: kind を想定した Deployment/Service マニフェスト

現時点ではフロントエンドはサンプル画面が中心で、バックエンドの API を呼び出す実装はありません。バックエンドにはユーザー、競馬場、レースの一覧取得 API が実装されています。

## システム構成

```text
ブラウザ
   |
   | :5173
   v
Frontend (Vue/Vite dev server)
   |
   | 現状、API 呼び出しコードなし
   |
   | 将来の HTTP/JSON 通信
   v
Backend (Go/Gin :3000)
   |
   v
MySQL 8.4 (:3306)
```

ローカルの Docker Compose では、`frontend`、`backend`、`mysql` の 3 サービスを起動します。バックエンドは MySQL の healthcheck が成功してから起動します。バックエンドは `DB_HOST=mysql` などの環境変数で Compose 内の MySQL に接続します。

## フロントエンド

### 技術スタックと起動

- `vue` 3.5 系、`vue-router` 5.3 系、TypeScript 6 系を使用します。
- Vite 8 を開発サーバーおよびビルドツールとして使用します。
- Tailwind CSS v4 は `@tailwindcss/vite` プラグインと `src/style.css` の `@import 'tailwindcss'` で読み込まれます。
- `vite-plugin-vue-devtools` が開発用プラグインとして有効です。
- `@` は `frontend/src` へのパスエイリアスです。
- Docker では Node.js 22 Alpine 上で `npm ci` を実行し、`npm run dev -- --host 0.0.0.0` でポート 5173 を公開します。

エントリーポイントは `src/main.ts` です。`App.vue` に Vue Router を登録し、`#app` にマウントします。

### レイヤー構成

```text
main.ts
  -> App.vue
  -> router/index.ts
  -> views/*                 # ルート単位の画面
  -> components/*            # 画面内で再利用する UI
  -> style.css               # グローバルスタイル
```

- **エントリーポイント層 (`main.ts`)**: Vue アプリケーションを生成し、Router とグローバル CSS を登録して DOM にマウントします。
- **アプリケーションルート層 (`App.vue`)**: アプリケーションの共通ルートです。現在は `RouterView` を配置し、ルート画面を表示します。
- **ルーティング層 (`router/index.ts`)**: `createWebHistory` を使って URL と画面コンポーネントを対応付けます。
- **画面層 (`views/`)**: URL 単位の画面を実装します。
- **UI コンポーネント層 (`components/`)**: 画面から利用する再利用可能な表示部品を配置します。
- **スタイル層 (`style.css`)**: Tailwind CSS の読み込みを行います。個別の UI スタイルは現状コンポーネントのテンプレート内に Tailwind のユーティリティクラスとして記述されています。

## バックエンド

### 起動と依存性注入

`backend/main.go` が以下の順序でアプリケーションを組み立てます。

1. `DB_HOST`、`DB_PORT`、`DB_NAME`、`DB_USER`、`DB_PASSWORD` を読み取って MySQL に接続する。
2. `application.New` で各 Store と Service を生成する。
3. Gin のデフォルトエンジンを生成し、ルーティングを登録する。
4. `:3000` で HTTP サーバーを起動する。

DB 接続は `database/sql` と `go-sql-driver/mysql` を使用します。接続時に `Ping` を行い、必要な環境変数がない場合や接続できない場合は起動に失敗します。開発用 Docker イメージでは Air による Go ファイルのホットリロードが有効です。

### レイヤー構成

```text
HTTP request
  -> config/router.go
  -> controller/*/handler.go
  -> service/*
  -> infrastructure/* (database/sql)
  -> MySQL
```

- **`controller`**: Gin の HTTP ハンドラー。Service を呼び出し、成功時は JSON、失敗時は 500 と固定のエラーメッセージを返します。
- **`service`**: 一覧取得などのユースケースを提供します。現在の実装では Store の呼び出しを委譲しています。
- **`infrastructure`**: SQL の発行と `database/sql.Rows` のモデルへの変換を担当します。
- **`model`**: `User`、`Race`、`Racecourse` と、レースの馬場・天候・方向などの型を定義します。
- **`application`**: DB 接続から各 Store/Service を組み立てる Composition Root です。

## データベース

MySQL 8.4 を使用し、Compose の named volume `mysql_data` にデータを永続化します。

## 実行環境

### Docker Compose

`docker compose` では次のポートを使用します。

| サービス | コンテナポート | ホスト公開ポート | 役割 |
| --- | ---: | ---: | --- |
| `mysql` | 3306 | 非公開 | データベース |
| `backend` | 3000 | 3000 | Gin API |
| `frontend` | 5173 | 5173 | Vite 開発サーバー |

Compose の MySQL は環境変数 `MYSQL_ROOT_PASSWORD`、`MYSQL_DATABASE`、`MYSQL_USER`、`MYSQL_PASSWORD` を必須とします。バックエンドは Compose のサービス名 `mysql` を DB ホストとして利用します。

### Kubernetes / kind

`k8s/` には frontend/backend の Deployment と ClusterIP Service があります。どちらもレプリカ数は 1 で、イメージは `tmp-app-frontend:latest` と `tmp-app-backend:latest` です。Makefile の `kind-load` でイメージを kind クラスターにロードし、`kind-deploy` でリソースを適用します。

現状の Kubernetes マニフェストには MySQL の Deployment/Service がなく、backend Deployment に DB 接続用環境変数も定義されていません。そのため、これらのマニフェストだけではバックエンドが DB に接続できず、Compose 構成と同等の一体型環境にはなりません。Service はどちらも `ClusterIP` のため、ローカルブラウザからアクセスする場合は `kubectl port-forward` が必要です。
