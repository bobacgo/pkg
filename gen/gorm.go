package gen

import (
	"fmt"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const (
	dsn = "root:xxx@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
)

var tableNames = []string{
	"test",
}

func generateModel() {
	base, _ := filepath.Abs("./")
	g := gen.NewGenerator(gen.Config{
		OutPath:           fmt.Sprintf("%s/dao", base),
		ModelPkgPath:      fmt.Sprintf("%s/model", base),
		WithUnitTest:      false,
		FieldNullable:     false,
		FieldCoverable:    true,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  false,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	gormDb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	g.UseDB(gormDb)
	for _, tbName := range tableNames {
		g.GenerateModel(tbName)
	}
	// g.ApplyBasic(model.xxx{}, model.xx{})
	g.Execute()
}
