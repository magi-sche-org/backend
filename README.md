# backend

## 前提条件

- `Docker`
- `Docker Compose V2`
- `make`
- `go 1.20`

## 初回実行

このディレクトリで

```console
cp example.env .env
```

（しなくてもよい）

## 起動

```console
make up
```

## DBマイグレーション

初回のみ

```console
go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
```

のちに，`make up`が実行された後に

```console
make migrate
```