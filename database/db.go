package database

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhulida1234/school-rpc/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type DB struct {
	gorm *gorm.DB

	SchoolDB *SchoolDB
}

func NewDB(ctx context.Context, dbConf config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Name)
	if dbConf.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConf.Port)
	}
	if dbConf.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConf.User)
	}
	if dbConf.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConf.Password)
	}

	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
		CreateBatchSize:        3_000,
	}

	// 直接尝试打开数据库连接
	gorm, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &DB{
		gorm:     gorm,
		SchoolDB: NewSchoolDB(gorm),
	}, nil
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}
		return nil
	})
	return err
}
