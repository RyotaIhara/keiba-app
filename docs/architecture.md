# アーキテクチャ

## 概要

本リポジトリは、競馬情報を扱うフロントエンドとバックエンドを分離した構成です。

- **フロントエンド**: Vue 3 + TypeScript + Vite
- **バックエンド**: Go + Gin
- **データベース**: MySQL 8.4
- **ローカル実行**: Docker Compose
- **Kubernetes 実行**: kind を想定した Deployment/Service マニフェスト

フロントエンドとバックエンドは HTTP/JSON で通信します。フロントエンドからバックエンド API への接続先は `VITE_API_BASE_URL` で切り替えます。個別の画面や API エンドポイントは、それぞれの実装を参照してください。

## システム構成

```text
ブラウザ
   |
   | :5173
   v
Frontend (Vue/Vite dev server)
   |
   | API client
   | HTTP/JSON
   v
Backend (Go/Gin :3000)
   |
   v
MySQL 8.4 (:3306)
```

ローカルの Docker Compose では、`frontend`、`backend`、`mysql` の 3 サービスを起動します。バックエンドは MySQL の healthcheck が成功してから起動します。ブラウザ上のフロントエンドは `VITE_API_BASE_URL=http://localhost:3000` を使って、ホストに公開されたバックエンドへ直接 HTTP リクエストを送ります。バックエンドは `DB_HOST=mysql` などの環境変数で Compose 内の MySQL に接続します。

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
  -> api/*                   # バックエンド API クライアント
  -> mappers/*               # APIレスポンスから画面表示モデルへの変換
  -> style.css               # グローバルスタイル
```

- **エントリーポイント層 (`main.ts`)**: Vue アプリケーションを生成し、Router とグローバル CSS を登録して DOM にマウントします。
- **アプリケーションルート層 (`App.vue`)**: アプリケーションの共通ルートです。現在は `RouterView` を配置し、ルート画面を表示します。
- **ルーティング層 (`router/index.ts`)**: `createWebHistory` を使って URL と画面コンポーネントを対応付けます。
- **画面層 (`views/`)**: URL 単位の画面を実装します。
- **UI コンポーネント層 (`components/`)**: 画面から利用する再利用可能な表示部品を配置します。
- **API クライアント層 (`api/`)**: `import.meta.env.VITE_API_BASE_URL` をベースURLとして HTTP リクエストを実行し、HTTPエラーを画面側へ通知します。個別の API クライアントは機能単位で配置します。
- **マッピング層 (`mappers/`)**: バックエンドのレスポンスモデルを画面表示用のモデルへ変換します。APIのデータ形式と画面表示用モデルを分離します。
- **スタイル層 (`style.css`)**: Tailwind CSS の読み込みを行います。個別の UI スタイルは現状コンポーネントのテンプレート内に Tailwind のユーティリティクラスとして記述されています。

### フロントエンドとバックエンドの通信

画面層は API クライアントを介してバックエンドと通信し、レスポンスを必要に応じて表示用モデルへ変換します。通信中・通信失敗時の表示は、各画面の要件に応じて実装します。

`VITE_API_BASE_URL` が未設定の場合、API クライアントはエラーとして扱います。ベースURLの末尾のスラッシュを正規化するなど、具体的なリクエスト処理は共通の API クライアント側に集約します。

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

### CORS

バックエンドは `CORS_ALLOW_ORIGINS` をカンマ区切りで読み取り、許可したフロントエンドのオリジンだけからのクロスオリジンリクエストを受け付けます。設定がない場合、または空のオリジンしか含まない場合はバックエンドの起動に失敗します。ローカル開発では `http://localhost:5173` と `http://127.0.0.1:5173` を許可します。

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

### 環境変数

環境変数はリポジトリルートの `.env.local` でローカル開発用に管理し、Docker Compose の各サービスへ `env_file` で渡します。

| 変数 | 利用箇所 | 役割 |
| --- | --- | --- |
| `VITE_API_BASE_URL` | frontend | ブラウザから接続するバックエンドAPIのベースURL |
| `CORS_ALLOW_ORIGINS` | backend | APIへのアクセスを許可するフロントエンドのオリジン一覧 |
| `DB_HOST` / `DB_PORT` | backend | MySQLの接続先 |
| `DB_NAME` / `DB_USER` / `DB_PASSWORD` | backend | MySQLのデータベース名・認証情報 |
| `MYSQL_ROOT_PASSWORD` / `MYSQL_DATABASE` / `MYSQL_USER` / `MYSQL_PASSWORD` | mysql | MySQLコンテナの初期設定 |

`VITE_*` はViteによってフロントエンドへ埋め込まれるクライアント公開設定です。パスワードなどの秘密情報は設定しません。`VITE_API_BASE_URL` の変更は、開発サーバーの起動時またはフロントエンドのビルド時に反映されます。

### Kubernetes / kind

`k8s/` には frontend/backend の Deployment と ClusterIP Service があります。どちらもレプリカ数は 1 で、イメージは `keiba-app-frontend:latest` と `keiba-app-backend:latest` です。Makefile の `kind-load` でイメージを kind クラスターにロードし、`kind-deploy` でリソースを適用します。

現状の Kubernetes マニフェストには MySQL の Deployment/Service、backend Deployment の DB接続用環境変数、frontend Deployment の `VITE_API_BASE_URL` が定義されていません。そのため、これらのマニフェストだけではバックエンドが DB に接続できず、フロントエンドから接続するAPI URLも環境に合わせて注入できません。Kubernetesで利用する場合は、DBと環境変数を別途用意し、フロントエンドのビルド時にAPIのベースURLを設定する必要があります。Service はどちらも `ClusterIP` のため、ローカルブラウザからアクセスする場合は `kubectl port-forward` が必要です。
