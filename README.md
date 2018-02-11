# CAMPHOR- COIN

## インストール

```
$ go get github.com/CAMPHOR-COIN/camphor-coin
$ cd $GOROOT/src/github.com/CAMPHOR-COIN/camphor-coin
$ dep ensure
```

## サーバーの起動方法

```
$ go run server.go
```

上記コマンドで8080番でサーバーが起動。 `/blockchain`を叩くと最新のBlockchainが確認できる

* 現在はGenesisBlockのみが格納された配列です
