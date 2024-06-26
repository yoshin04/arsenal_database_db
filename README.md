# アーセナルベース DB

アーケードカードゲーム「ガンダムアーセナルベース」のデッキシミュレーター。

### 公式サイト

https://gundam-ab.com/

### 参考資料

https://docs.google.com/spreadsheets/d/1JZhjK050sCrYD0AITa6aBPbTbRLUFpXrbH5ItDif0as/edit?usp=drivesdk

## 目次

- [使い方](#使い方)
- [技術スタック](#技術スタック)
- [設計資料](#設計資料)
- [開発環境構築](#開発環境構築)

## 使い方

開発中...

## 技術スタック

- API: Gin - Go
- Front: Next.js - TypeScript
- Docker
- REST API

## 設計資料
### 要件定義書
https://docs.google.com/spreadsheets/d/1SCDyXZiIJsBeGlSgbpCxFYA66yRFzzEkoxgR75UvUVo/edit?pli=1#gid=0

### ワイヤーフレーム
【figma】
https://www.figma.com/file/BmWFMKjVwnRXcToXs7DvTY/%E7%84%A1%E9%A1%8C?type=design&node-id=0%3A1&mode=design&t=OzSmpu1W6zZRssBT-1

### テーブル設計図
https://docs.google.com/spreadsheets/d/13ZCmZQymBbb-SQ5BRk4H2OxBvnJjfP74emdfa3CtdDo/edit?usp=sharing

### ER図
未作成。

### APIドキュメント
Apidogで管理。
https://app.apidog.com/

## 開発環境構築

ローカルマシンでの開発環境のセットアップ方法を説明します。

【必要条件】

- Go（バージョン 1.22 以上）
- Node.js（バージョン 20 以上）
- Docker
- git

1. リポジトリのクローン

```
  git clone https://github.com/yoshin04/arsenal_database_db.git
  cd arsenal_database_db
```

2. API のセットアップ

- api ディレクトリに移動。
- 必要な Go パッケージをインストール。
- docker build
- コンテナ起動

```
  cd api
  go mod download
  docker compose build
  docker compose up
```

3. フロントエンドのセットアップ

```
  cd arsenalbase_decksimulator
  yarn dev
  http://localhost:3000
```
