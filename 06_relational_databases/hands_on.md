# CRUDの一連の流れを実行する

PostgreSQLの公式ドキュメント：<https://www.postgresql.org/docs/12/tutorial.html>

## 1. Getting Started

## 1.1(インストール)

- ソースコード
- インストーラ
- homebrew
- zip

など方法は色々あるらしいですが、今回はhomebrewでインストールします。楽なので。

1. `brew install postgresql`を実行
1. `brew services start postgresql`でログイン時に自動起動するようにする。尚、やめたい場合は`start`→`stop`で。
1. `psql -U${USER} postgres`でログイン。`-U${USER}`は環境変数で、インストール時のmacOSログインユーザ名になっている。
1. 👆をそのまま使うのは良くないので、`postgres=# create user postgres with SUPERUSER;`で操作用のユーザ「postgres」を作成。
1. `psql -Upostgres` ログインできるかどうかを確認
1. `\l`でデータベースのリストが取得できたらOK!

## 1.2(構造的な基本事項)

- PostgreSQL（ぽすとぐれすきゅーえる）は、クライアント/サーバモデルを使用している
- セッションは以下の協調動作するプロセスから構成される
  - サーバプロセス：DBファイルの管理、クライアントアプリケーションからの接続を受け、処理を行う。
  - クライアント：データベース操作を行うフロントアプリケーション。テキスト指向、グラフィカルなもの、色々ある。
- PostgreSQLサーバは、クライアントから複数の同時接続を取り扱うことができる。
- このため、サーバは接続ごとに新しいプロセスを開始（fork）する。
- その時点から、クライアントと新しいサーバプロセスは、元のサーバプロセスによる干渉が無い状態で通信を行える。

大体こんな感じ

![クライアント/サーバ](クライアントサーバ.png)

## 1.3(データベースの作成)

## 1.4(データベースへのアクセス)

## 2.1(序文)

## 2.2(概念)


- 2. The SQL Language

※ MySQLの公式ドキュメント：<https://dev.mysql.com/doc/refman/5.6/ja/tutorial.html>

- 3.1 ~ 3.4, 3.6