package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shoot3rs/user/internal/types"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type appConfig struct {
}

func (cnf *appConfig) GetServerConfig() *http2.Server {
	return &http2.Server{
		MaxHandlers:                  0,
		MaxConcurrentStreams:         0,
		MaxDecoderHeaderTableSize:    0,
		MaxEncoderHeaderTableSize:    0,
		MaxReadFrameSize:             0,
		PermitProhibitedCipherSuites: false,
		IdleTimeout:                  0,
		ReadIdleTimeout:              0,
		PingTimeout:                  0,
		WriteByteTimeout:             0,
		MaxUploadBufferPerConnection: 0,
		MaxUploadBufferPerStream:     0,
		NewWriteScheduler:            nil,
		CountError:                   nil,
	}
}

func (cnf *appConfig) GetServerAddr() string {
	return fmt.Sprintf(":%s", os.Getenv("APP.PORT"))
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
