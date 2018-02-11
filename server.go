package main

import "github.com/CAMPHOR-COIN/camphor-coin/infrastructure"

func main() {
	// サーバーの起動
	infrastructure.Router.Run()
}
