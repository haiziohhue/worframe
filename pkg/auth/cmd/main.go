package main

import (
	authCore "worframe/pkg/auth/core"
	"worframe/share/core"
)

func main() {
	shareApp := core.
		NewApp("dev").InitPublicZap().InitDb().InitRedis()

	if shareApp.Error != nil {
		panic(shareApp.Error)
	}

	App := authCore.
		NewAuthCore(shareApp).InitEngine()

	if App.Error != nil {
		panic(App.Error)
	}

	App.Run()

}
