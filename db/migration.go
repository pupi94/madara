package db

import (
	"errors"
	"fmt"
	"github.com/pupi94/madara/config"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"xorm.io/xorm"
)

func MigrateUp() {
	db := config.DB
	createSchemaTable(db)

	dbLastVersion := findLastVersion(db)

	files := listMigrationFiles("up")
	var unexecutedFiles []string
	if dbLastVersion != "" {
		for _, file := range files {
			version := getVersionByFileName(file)
			if version > dbLastVersion {
				unexecutedFiles = append(unexecutedFiles, file)
			}
		}
	} else {
		unexecutedFiles = files
	}

	for _, file := range unexecutedFiles {
		path := fmt.Sprintf("%s/%s", migrationsDirPath(), file)
		_, err := db.ImportFile(path)
		if err != nil {
			panic(err)
		}
		upMigrationVersion(db, getVersionByFileName(file))
	}
}

func MigrateDown(step int) {
	db := config.DB

	dbLastVersion := findLastVersion(db)
	if dbLastVersion == "" {
		return
	}

	files := listMigrationFiles("down")

	var executedFiles []string
	// 反向遍历 files
	for i, _ := range files {
		file := files[len(files)-i-1]
		version := getVersionByFileName(file)
		if dbLastVersion >= version {
			executedFiles = append(executedFiles, file)
			if len(executedFiles) >= step {
				break
			}
		}
	}

	for _, file := range executedFiles {
		path := fmt.Sprintf("%s/%s", migrationsDirPath(), file)
		_, err := db.ImportFile(path)
		if err != nil {
			panic(err)
		}
		downMigrationVersion(db, getVersionByFileName(file))
	}
}

func GenerateMigration(name string) {
	if name == "" {
		panic(errors.New("migration file name undefined"))
	}

	createMigrationsDir()

	t := time.Now().Format("20060102150405")
	up := fmt.Sprintf("%s/db/migrations/%s_%s.up.sql", basePath(), t, name)
	down := fmt.Sprintf("%s/db/migrations/%s_%s.down.sql", basePath(), t, name)

	upFile, err := os.Create(up)
	if err != nil {
		panic(err)
	}
	_ = upFile.Close()

	downFile, err := os.Create(down)
	if err != nil {
		panic(err)
	}
	_ = downFile.Close()
}

func createMigrationsDir() {
	path := fmt.Sprintf("%s/db/migrations", basePath())
	_, err := os.Stat(path)
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		panic(err)
	}
	if err = os.Mkdir(path, os.ModePerm); err != nil {
		panic(err)
	}
}

func getVersionByFileName(fileName string) string {
	return strings.Split(fileName, "_")[0]
}

// 当前执行命令的路径
func basePath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

func migrationsDirPath() string {
	return fmt.Sprintf("%s/db/migrations", basePath())
}

func listMigrationFiles(direction string) []string {
	dir, err := ioutil.ReadDir(migrationsDirPath())
	if err != nil {
		panic(err)
	}

	var files []string
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			continue
		}
		suffix := fmt.Sprintf(".%s.sql", direction)
		if ok := strings.HasSuffix(fi.Name(), suffix); !ok {
			continue
		}

		files = append(files, fi.Name())
	}
	return files
}

// schema_migrations
func createSchemaTable(db *xorm.Engine) {
	schemaExist, err := db.IsTableExist(&SchemaMigration{})
	if err != nil {
		panic(err)
	}
	if schemaExist {
		return
	}

	err = db.CreateTables(&SchemaMigration{})
	if err != nil {
		panic(err)
	}
}

// 记录已执行的 migration
type SchemaMigration struct {
	Version string `xorm:"varchar(14)"`
}

func (sm SchemaMigration) TableName() string {
	return "schema_migrations"
}

func findLastVersion(db *xorm.Engine) string {
	migration := new(SchemaMigration)
	ok, err := db.OrderBy("version DESC").Limit(1).Get(migration)
	if err != nil {
		panic(err)
	}
	if ok {
		return migration.Version
	}
	return ""
}

func upMigrationVersion(db *xorm.Engine, version string) {
	migration := &SchemaMigration{Version: version}
	_, err := db.Insert(migration)
	if err != nil {
		panic(err)
	}
}

func downMigrationVersion(db *xorm.Engine, version string) {
	_, err := db.Exec(`DELETE FROM schema_migrations WHERE version >= ?`, version)
	if err != nil {
		panic(err)
	}
}
