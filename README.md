# アーセナルベースDB
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
- [コントリビューション](#コントリビューション)

## 使い方
開発中...

## 技術スタック
* API: Gin - Go
* Front: Next.js - TypeScript
* Docker
* REST API

## 設計資料

## 開発環境構築
ローカルマシンでの開発環境のセットアップ方法を説明します。
【必要条件】
* Go（バージョン1.22以上）
* Node.js（バージョン20以上）
* Docker
* git

1. リポジトリのクローン
```
  git clone https://github.com/yoshin04/arsenal_database_db.git
  cd arsenal_database_db
```
2. APIのセットアップ
* api ディレクトリに移動。
* 必要なGoパッケージをインストール。
* docker build
* コンテナ起動
```
  cd api
  go mod download
  docker compose build
  docker compose up 
```
3. フロントエンドのセットアップ

## コントリビューション
