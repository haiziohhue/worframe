package main

import (
	authCore "worframe/pkg/auth/core"
	"worframe/share/core"
)

func main() {
	shareApp := core.
		NewApp("dev").InitPublicZap().InitDb().InitRedis()

	if shareApp.GetErr() != nil {
		panic(shareApp.GetErr())
	}

	App := authCore.
		NewAuthCore(shareApp).InitEngine()

	if App.GetErr() != nil {
		panic(App.GetErr())
	}
	App.Run()
}
