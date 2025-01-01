package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/morkid/paginate"
	"github.com/shooters/user/internal/types"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type appConfig struct {
}

func (cnf *appConfig) GetServerAddr() string {
	return fmt.Sprintf(":%s", os.Getenv("APP.PORT"))
}

func (cnf *appConfig) GetPaginate() paginate.Config {
	return paginate.Config{
		Operator:             "",
		FieldWrapper:         "",
		ValueWrapper:         "",
		DefaultSize:          50,
		PageStart:            0,
		LikeAsIlikeDisabled:  false,
		SmartSearchEnabled:   false,
		Statement:            nil,
		CustomParamEnabled:   false,
		SortParams:           nil,
		PageParams:           nil,
		OrderParams:          nil,
		SizeParams:           nil,
		FilterParams:         nil,
		FieldsParams:         nil,
		FieldSelectorEnabled: false,
		CacheAdapter:         nil,
		JSONMarshal:          nil,
		JSONUnmarshal:        nil,
		ErrorEnabled:         false,
	}
}

func (cnf *appConfig) Logger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := config.Build()
	if err != nil {
		return nil
	}

	return zapLogger
}

func (cnf *appConfig) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env")
	}
	log.Println(".env loaded successfully!")
}

func (cnf *appConfig) GetGorm() *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   logger.Default.LogMode(logger.Info),
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		TranslateError:                           true,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}
}

func New() types.GlobalConfig {
	return &appConfig{}
}
