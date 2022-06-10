package seeders

import (
	"fmt"
	"liu/database/factories"
	"liu/pkg/console"
	"liu/pkg/logger"
	"liu/pkg/seed"

	"gorm.io/gorm"
)

func init() {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		users := factories.MakeUsers(10)
		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
