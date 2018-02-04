package main

import "app/infrastructure"

func main() {
	// サーバーの起動
	infrastructure.Router.Run()
}
