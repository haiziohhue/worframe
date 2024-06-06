package router

import "github.com/gin-gonic/gin"

func RbacRouter(r *gin.Engine) {
	routerDept(r)
	routerRole(r)
	routerMenu(r)
}
