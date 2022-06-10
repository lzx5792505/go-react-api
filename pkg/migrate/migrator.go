package migrate

import (
	"io/ioutil"
	"liu/pkg/console"
	"liu/pkg/database"
	"liu/pkg/file"

	"gorm.io/gorm"
)

type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

func NewMigrator() *Migrator {
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	migrator.createMigrationsTable()
	return migrator
}

func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) Up() {
	migrateFiles := migrator.readAllMigrationFiles()
	batch := migrator.getBatch()
	migrations := []Migration{}
	migrator.DB.Find(&migrations)
	runed := false
	for _, mfile := range migrateFiles {
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}
	if !runed {
		console.Success("database is up to date.")
	}
}

func (migrator *Migrator) Rollback() {
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?", lastMigration.Batch).Order("id DESC").Find(&migrations)
	// 回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 标记是否真的有执行了迁移回退的操作
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	runed := false
	for _, _migration := range migrations {
		console.Warning("rollback " + _migration.Migration)
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}
		runed = true
		migrator.DB.Delete(&_migration)
		console.Success("finsh " + mfile.FileName)
	}
	return runed
}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	batch := 1
	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	// 如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	files, err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile

	for _, f := range files {
		// 去掉后缀名称
		fileName := file.FileNameWithoutExtension(f.Name())
		mfile := getMigrationFile(fileName)
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}

	return migrateFiles
}

func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	if mfile.Up != nil {
		console.Warning("migrating " + mfile.FileName)
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		console.Success("migrated " + mfile.FileName)
	}

	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

// 回滚所有迁移
func (migrator *Migrator) Reset() {
	migrations := []Migration{}
	// 按照倒序读取所有迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)
	// 回滚所有迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

// 回滚所有迁移，并运行所有迁移
func (migrator *Migrator) Refresh() {
	// 回滚所有迁移
	migrator.Reset()
	// 再次执行所有迁移
	migrator.Up()
}

//  所有的表并重新运行所有迁移
func (migrator *Migrator) Fresh() {

	// 获取数据库名称，用以提示
	dbname := database.CurrentDatabase()

	// 删除所有表
	err := database.DeleteAllTables()
	console.ExitIf(err)
	console.Success("clearup database " + dbname)

	// 重新创建 migrates 表
	migrator.createMigrationsTable()
	console.Success("[migrations] table created.")

	// 重新调用 up 命令
	migrator.Up()
}
