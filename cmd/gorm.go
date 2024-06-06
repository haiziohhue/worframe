package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "pkg/query",
		Mode:    gen.WithoutContext, // script mode
	})

	gormdb, err := gorm.Open(postgres.Open("postgres://postgres:123456@127.0.0.1:15432/sys_security"))
	if err != nil {
		panic(err)
	}
	g.UseDB(gormdb)
	generateTable(g)
}
func generateTable(g *gen.Generator) {
	g.GenerateAllTable()
	g.Execute()
}
func generateColumn(g *gen.Generator) {

	//Generate basic type-safe query API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(
	//	func(query Query) {},
	//	model.RolePerm{},
	//	model.UserPerm{},
	//	model.RbacPerm{},
	//	model.RbacRole{},
	//	model.DeptRole{},
	//	model.RbacAuditLog{},
	//	model.UserDept{},
	//	model.UserRole{},
	//	model.SysUser{},
	//	model.RbacDept{},
	//)

	//g.ApplyInterface(tableModles)
	//Generate the code
}
