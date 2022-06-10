package migrate

import (
	"database/sql"

	"gorm.io/gorm"
)

type migrationFunc func(gorm.Migrator, *sql.DB)

var migrationFiles []MigrationFile

type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	FileName string
}

func Add(name string, up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles, MigrationFile{
		FileName: name,
		Up:       up,
		Down:     down,
	})
}


// 通过迁移文件的名称来获取到 MigrationFile 对象
func getMigrationFile(name string) MigrationFile {
    for _, mfile := range migrationFiles {
        if name == mfile.FileName {
            return mfile
        }
    }
    return MigrationFile{}
}

// 判断迁移是否已执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
    for _, migration := range migrations {
        if migration.Migration == mfile.FileName {
            return false
        }
    }
    return true
}