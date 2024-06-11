package core

import (
	"fmt"
	"worframe/pkg/auth/server"
)

func (ac *AuthCore) Run() {
	port := fmt.Sprintf(":%d", ac.Conf.Server.Port)
	err := ac.Engine.Run(port)
	if err != nil {
		ac.Error = err
		panic(err)
	}
}
func (ac *AuthCore) InitEngine() *AuthCore {
	ac.Engine = server.InitEngine(ac.Logger, ac.DB)
	return ac
}
