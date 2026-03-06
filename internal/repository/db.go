package repository

import (
	"campushelphub/internal/config"
	"campushelphub/model"
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) *gorm.DB {
	dsn := cfg.Database.User + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + strconv.Itoa(cfg.Database.Port) + ")/" + cfg.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Competition{})
	return db
}

// 测试用函数，用于删除数据库中所有表
func DropAllTables(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("gorm.DB 对象不能为空")
	}

	var tables []string
	var err error

	switch db.Dialector.Name() {
	case "mysql":
		// MySQL 获取所有表名
		err = db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = DATABASE()").
			Pluck("table_name", &tables).Error
	case "postgres":
		// PostgreSQL 获取所有表名（排除系统表）
		err = db.Raw(`
			SELECT table_name 
			FROM information_schema.tables 
			WHERE table_schema = current_schema() 
			AND table_type = 'BASE TABLE' 
			AND table_name NOT LIKE 'pg_%' 
			AND table_name != 'information_schema'
		`).Pluck("table_name", &tables).Error
	case "sqlite":
		// SQLite 获取所有表名（排除系统表）
		err = db.Raw(`
			SELECT name 
			FROM sqlite_master 
			WHERE type = 'table' 
			AND name NOT LIKE 'sqlite_%'
		`).Pluck("name", &tables).Error
	default:
		return fmt.Errorf("不支持的数据库类型: %s", db.Dialector.Name())
	}

	if err != nil {
		return fmt.Errorf("获取表名失败: %w", err)
	}

	if len(tables) == 0 {
		fmt.Println("数据库中无表可删除")
		return nil
	}

	if db.Dialector.Name() == "mysql" {
		if err := db.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
			return fmt.Errorf("禁用外键约束失败: %w", err)
		}
		defer db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	}

	for _, table := range tables {
		dropSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
		if err := db.Exec(dropSQL).Error; err != nil {
			return fmt.Errorf("删除表 %s 失败: %w", table, err)
		}
		fmt.Printf("成功删除表: %s\n", table)
	}

	return nil
}
