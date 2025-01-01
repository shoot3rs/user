package database

import (
	"fmt"
	"github.com/shooters/user/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"strings"
	"time"
)

type gormDB struct {
	dbType string
	engine *gorm.DB
	config *gorm.Config
	models []interface{}
}

func (db *gormDB) GetEngine() interface{} {
	return db.engine
}

func (db *gormDB) GetConfig() *gorm.Config {
	return db.config
}

func (db *gormDB) Connect() {
	db.dbType = os.Getenv("DB.TYPE")
	log.Println("Loading db config for: ", db.dbType)
	var err error
	switch strings.ToLower(strings.TrimSpace(db.dbType)) {
	case "sqlite":
		db.engine, err = db.loadSqliteDB()
		if err != nil {
			log.Fatalln("[‼️] unable to load sqlite db", err)
		}
	case "postgres":
		log.Println("loading config for postgres")
		db.engine, err = db.loadPostgresDB()
		if err != nil {
			log.Fatalln("[‼️] unable to load postgres db", err)
		}

		log.Println("postgres connected successfully")
	}

	sqlDB, _ := db.engine.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 8765)

	if err := sqlDB.Ping(); err != nil {
		log.Println("[‼️] postgres connected successfully")
	}

	if db.dbType != "sqlite" {
		if err := db.createSequence(); err != nil {
			log.Println("[‼️] unable to create sequence", err)
		}
	}

	db.handleMigrations()
}

func (db *gormDB) loadSqliteDB() (*gorm.DB, error) {
	dbName := fmt.Sprintf("%s.db", os.Getenv("DB.NAME"))
	conn, err := gorm.Open(sqlite.Open(dbName), db.GetConfig())

	if err != nil {
		panic("failed to connect db")
		return nil, err
	}
	log.Println("connected to sqlite db")

	return conn, nil
}

func (db *gormDB) loadPostgresDB() (*gorm.DB, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB.HOST"),
		os.Getenv("DB.USER"),
		os.Getenv("DB.PASS"),
		os.Getenv("DB.NAME"),
		os.Getenv("DB.PORT"),
	)
	for {
		c, err := gorm.Open(postgres.Open(dsn), db.GetConfig())
		if err != nil {
			fmt.Println("Postgres not yet ready to connect!")
			counts++
		} else {
			fmt.Println("Connected to postgres!")
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
	}

	//db.createSequence()

	fmt.Println("Connection to postgres established successfully!")
	return connection, nil
}

func (db *gormDB) createSequence() error {
	// Step 1: Create the sequence if it does not exist
	if err := db.engine.Exec("CREATE SEQUENCE IF NOT EXISTS code_seq").Error; err != nil {
		return err
	}

	log.Println("Sequence created successfully!")
	return nil
}

func (db *gormDB) handleMigrations() {
	err := db.engine.AutoMigrate(db.models...)
	if err != nil {
		log.Fatalf("[‼️] migration failed: %s", err.Error())
	}
	log.Println("tables migrated successfully!")
}

func NewDatabase(cfg *gorm.Config) types.Connection {
	return &gormDB{
		config: cfg,
		models: []interface{}{},
	}
}
