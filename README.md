# Eureka Summer Internship 2018 API

# 技術スタック

- go 1.11.0
- dep
- go-swagger
- goose
- direnv
- xorm

# swagger

localhost:8081 で swagger-editor (エディタ), localhost:8082 で swagger-ui (APIドキュメント) が開きます。

```
docker-compose up -d
```

# how to run the app

```
# 必要なライブラリの取得

go get -u bitbucket.org/liamstask/goose/cmd/goose
go get -u github.com/golang/dep/cmd/dep
go get -u github.com/go-swagger/go-swagger/cmd/swagger
go get -u github.com/direnv/direnv

# 依存関係のインストール (dep ensureとか)
make init

# 環境変数を.envrc (direnv) で管理している
cp .envrc.sample .envrc
direnv allow

# DBの初期化 & マイグレ
make setup-db

# サーバーを立ち上げる
make run
```

# dummy data

misc/dummy/ 下にダミーデータ生成のスクリプトを置いてます。以下makeコマンドは ## recreate db and insert dummy data

```
make setup-db
```

# migration with goose

マイグレーションツールのgooseを使用しています。

```
# ./db/migrations/20180809183923_createUser.sql が作成される
goose create createHoge sql

# up
goose up

# down
goose down

# redo
goose redo
```
