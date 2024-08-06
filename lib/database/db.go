package database

import (
	"context"
	"go-server/lib/domain/user"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	db *gorm.DB
}

func NewDB(lc fx.Lifecycle) (*Database, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		return nil, err
	}

	database := &Database{db: db}

	// Lifecycle을 사용하여 시작과 종료를 관리
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// 데이터베이스가 성공적으로 열렸음을 로깅
			log.Println("Database connection established")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// 데이터베이스 연결을 안전하게 닫기
			sqlDB, err := db.DB()
			if err != nil {
				return err
			}
			log.Println("Database connection closed")
			return sqlDB.Close()
		},
	})

	return database, nil
}
