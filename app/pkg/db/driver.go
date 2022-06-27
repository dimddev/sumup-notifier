package db

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"sumup-notifier/app/internal/config"
	"sumup-notifier/app/notifier/providers"
)

const (
	DSN = "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
)

func NewDriver(cfg providers.DBConfiger) (*gorm.DB, error) {
	conf, errCfg := cfg.Decoder(cfg.DBConfig())
	if errCfg != nil {
		return nil, errCfg
	}

	var dbConf config.DataBaseConfig

	errDecode := json.Unmarshal(conf, &dbConf)
	if errDecode != nil {
		return nil, errDecode
	}

	connectionDSN := fmt.Sprintf(
		DSN,
		dbConf.Host,
		dbConf.User,
		dbConf.Password,
		dbConf.DataBase,
		dbConf.Port,
		dbConf.SSL,
	)

	driver, err := gorm.Open(postgres.Open(connectionDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return driver, nil
}
