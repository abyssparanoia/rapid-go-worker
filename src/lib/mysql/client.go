package mysql

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL Driverの読み込み
	"github.com/jinzhu/gorm"
)

// Client ... クライアント
type Client struct {
	db *gorm.DB
}

// GetDB ... GormのDBを取得する
func (c *Client) GetDB(ctx context.Context) *gorm.DB {
	db := c.db.New()
	db.SetLogger(gorm.Logger{
		LogWriter: NewLogger(ctx),
	})
	return db
}

// NewClient ... クライアントを作成する
func NewClient(cfg *Config) *Client {
	dbs := fmt.Sprintf("%s:%s@%s/%s?parseTime=true&loc=Asia%%2FTokyo",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.DB)
	db, err := gorm.Open("mysql", dbs)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return &Client{
		db: db,
	}
}
