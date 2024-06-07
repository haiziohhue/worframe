package router

import "github.com/gin-gonic/gin"

func AuthRouter(r *gin.Engine) {
	routerDept(r)
	routerRole(r)
	routerMenu(r)
}
