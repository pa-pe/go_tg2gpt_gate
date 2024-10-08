package model

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"os"
	"strings"
	"upserv/config"
	"upserv/logger"
)

var connUrl string
var dbDriver string

func ConnectDb() *gorm.DB {
	dbDriver = strings.ToLower(config.Get("db", "driver"))
	logger.LaunchLog("DB driver=" + dbDriver)

	if dbDriver == "none" {
		return nil
	}

	logLevel := gormLogger.Silent
	if strings.ToLower(config.Get("logger", "level")) == "debug" {
		logLevel = gormLogger.Info
	}

	gormCfg := &gorm.Config{
		Logger: gormLogger.Default.LogMode(logLevel),
	}

	var dialector gorm.Dialector

	if dbDriver == "postgres" {
		host := config.Get("db", "host")
		port := config.Get("db", "port")
		dbName := config.Get("db", "dbName")
		user := config.Get("db", "user")
		password := config.Get("db", "password")

		connUrl = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName + "?sslmode=disable"
		dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"

		//		db, err = gorm.Open(postgres.Open(dsn), gormCfg)
		dialector = postgres.Open(dsn)
	} else if dbDriver == "sqlite" {
		file := config.Get("db", "file")
		dialector = sqlite.Open(file)
	} else {
		logger.Log.Panic("unsupported config db.driver=", dbDriver)
	}

	db, err := gorm.Open(dialector, gormCfg)
	if err != nil {
		logger.Log.Panic("failed to connect database: ", err)
	}

	if db == nil {
		logger.Log.Panic("db is a nil pointer!")
	}

	masterDB, err := db.DB()
	if err != nil {
		logger.Log.Panic(err)
	}
	if masterDB == nil {
		logger.Log.Panic("MasterDb is a nil pointer!")
	}
	if err = masterDB.Ping(); err != nil {
		logger.Log.Panic(err)
	}
	logger.LaunchLog("DB connection initialized...")

	return db
}

func Migration() {
	if dbDriver == "postgres" {
		postgresMigration()
	} else if dbDriver == "sqlite" {
		sqliteMigration()
	} else {
		logger.Log.Panic("unsupported Migration() for db.driver=", dbDriver)
	}
}

func sqliteMigration() {
	logger.LaunchLog("sqliteMigration: implementation needed")
	os.Exit(2)
}

func postgresMigration() {
	m, err := migrate.New("file://migrations", connUrl)
	if err != nil {
		logger.Log.Panic(err)
	}
	startVer, dirty, err := m.Version()
	// no versioning detected
	if err != nil {
		logger.LaunchLog("Could not determinate current migration version")
		err = m.Up()    //init first version
		if err == nil { // migration applied successfully
			newVersion, b, _ := m.Version()
			logger.LaunchLog(fmt.Sprintf("Migration applied. Version %d, dirty: %t", newVersion, b))
		}
		return
	}
	// previous migration was dirty (Why??) Fixing it
	if dirty {
		err = m.Force(int(startVer))
		if err != nil {
			logger.LaunchLog(err.Error())
			logger.Log.Panic("Force set migration failed. Need to resolve issue manually")
		}
	}
	logger.LaunchLog(fmt.Sprintf("Current migration Version: %d", startVer))
	err = m.Up()
	if err == nil { // migration applied successfully
		newVersion, b, _ := m.Version()
		logger.LaunchLog(fmt.Sprintf("Migration applied. Version %d, dirty: %t", newVersion, b))
		return
	}
	switch err.Error() {
	case "no change":
		logger.LaunchLog("No new migration detected")
		break
	case "file does not exist":
		logger.LaunchLog("No migrations file detected")
		break
	default:
		logger.LaunchLog(err.Error())
		logger.LaunchLog(fmt.Sprintf("Starting to rollback migrations to version: %d", startVer))
		brokenVer, _, _ := m.Version()
		err = m.Force(int(brokenVer)) // make not dirty to rollback
		if err != nil {               // not gonna happens
			logger.LaunchLog(err.Error())
			logger.LaunchLog("Force set migration failed. Need to resolve issue manually")
			os.Exit(2)
		}
		err = m.Migrate(startVer)
		if err != nil { // migrate to specific version failed
			logger.LaunchLog(err.Error())
			logger.LaunchLog("Migration back failed. Need to resolve issue manually")
			os.Exit(2)
		}
		logger.LaunchLog(fmt.Sprintf("Rollback success. Version: %d", startVer))
		logger.LaunchLog("Migration is broken, exiting")
		os.Exit(2)
	}
}
