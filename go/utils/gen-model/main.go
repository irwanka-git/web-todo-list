package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:112w4nka@tcp(host.docker.internal:8036)/db_todolist?charset=utf8mb4&parseTime=True&loc=Local"
	g := gen.NewGenerator(gen.Config{
		OutPath:           "utils/gen-model/export",
		ModelPkgPath:      "utils/gen-model/entity",
		FieldWithTypeTag:  false,
		FieldWithIndexTag: false,
		FieldSignable:     false,
	})
	gormdb, _ := gorm.Open(mysql.Open(dsn))
	g.UseDB(gormdb)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
