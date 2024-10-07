// Copyright (c) 2022 R_Radzhabov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"crypto/tls"
	"fmt"
	"gitlab.com/AngelX/common/config"
	"gorm.io/gorm"
	"net/http"
	"os"
	"upserv/logger"
	"upserv/src"
	"upserv/src/controller"
	"upserv/src/http/middleware"
	"upserv/src/service"
	"upserv/src/service/cache"
	"upserv/src/storage"
	"upserv/src/storage/model"
)

// @title Basic API
// @version 1.0
// @description Basic api swagger description
// @contact.name Api service support
//
// @securitydefinitions.apikey UserToken
// @in header
// @name X-Token-Key
//
// @host your.endpoint.com
func main() {
	//Init config
	fmt.Print("Beginning to execute")
	err := config.Init("config.ini", "")
	if err != nil {
		fmt.Println("Fail to load config")
		fmt.Println(err.Error())
		return
	}
	// Init Logger
	loggerConfig := logger.Config{
		LogLevel:      config.Get("logger", "level"),
		Env:           config.Get("logger", "env_name"),
		AccessLogFile: config.Get("logger", "access_log_file"),
	}
	logger.Init(loggerConfig)
	logger.LaunchLog("Logger initialized...")
	logger.LaunchLog("Log level is :" + config.Get("logger", "level"))

	var db *gorm.DB
	if config.Get("db", "name") != "your_database" {
		// Get new DB
		db := model.NewDb(
			config.Get("db", "host"),
			config.Get("db", "port"),
			config.Get("db", "name"),
			config.Get("db", "user"),
			config.Get("db", "password"),
		)
		if db == nil {
			logger.Log.Panic("db is a nil pointer!")
		}
		sqlDB, err := db.DB()
		if err != nil {
			logger.Log.Panic(err)
		}
		defer func() {
			cErr := sqlDB.Close()
			if cErr != nil {
				logger.LaunchLog(cErr.Error())
			}
		}()
	}

	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "migrate" {
			Migrate()
			return
		} else {
			logger.LaunchLog("Unrecognized command: " + command)
			logger.LaunchLog("Exiting...")
		}
		os.Exit(2)
	}

	// init storages
	storages := storage.NewStorages(db)
	// init cache
	serviceCache := cache.NewInMemoryCache()
	// init services
	services := service.NewServices(storages, serviceCache)

	// init controllers
	controller.InitControllers(services)
	if !controller.IsValid() {
		logger.Log.Panic("Invalid controllers")
	}

	tgBot := src.NewTelegramBot(config.Get("telegram", "token"))
	go tgBot.ListenAndServ()

	//init middlewares
	middlewares := middleware.NewMiddlewares(services)

	// Init Router
	router := src.NewRouter(middlewares)

	//--> Run Server <--
	serv := &http.Server{
		Addr:    config.Get("server", "port"),
		Handler: router,
	}
	if config.Get("server", "schema") == "https" {
		serv.TLSConfig = tlsConfigCreate() // initialize certificates
		if len(serv.TLSConfig.Certificates) < 1 {
			logger.LaunchLog("Could not start HTTPS server without SSL certificate provided, Use HTTP or provide valid certificates")
			return
		}
		logger.LaunchLog("Server started on HTTPS. Listening port " + config.Get("server", "port"))
		logger.LaunchLog("Started") // String that detected by starting script to determinate if it was started successfully
		logger.Log.Fatal(serv.ListenAndServeTLS("", ""))
	} else {
		logger.LaunchLog("Server started on HTTP. Listening port " + config.Get("server", "port"))
		logger.LaunchLog("Started") // String that detected by starting script to determinate if it was started successfully
		logger.Log.Fatal(serv.ListenAndServe())
	}

}

func tlsConfigCreate() *tls.Config {
	comCrtPath := config.Get("certificate_com", "public")
	comKeyPath := config.Get("certificate_com", "private")

	tlsConfig := &tls.Config{}
	if comCrtPath != "" {
		comCert, err := tls.LoadX509KeyPair(comCrtPath, comKeyPath)
		if err != nil {
			logger.LaunchLog("Could not load .com cert. Path: " + comCrtPath)
			logger.LaunchLog(err.Error())
		} else {
			logger.LaunchLog("Load .com cert success Path: " + comCrtPath)
			tlsConfig.Certificates = append(tlsConfig.Certificates, comCert)
		}
	}
	cnCrtPath := config.Get("certificate_cn", "public")
	cnKeyPath := config.Get("certificate_cn", "private")
	//tlsConfig.Certificates = make([]tls.Certificate, 2)
	if cnCrtPath != "" {
		cnCert, err := tls.LoadX509KeyPair(cnCrtPath, cnKeyPath)
		if err != nil {
			logger.LaunchLog("Could not load .cn cert. Path: " + cnCrtPath)
			logger.LaunchLog(err.Error())
		} else {
			logger.LaunchLog("Load .cn cert success Path: " + cnCrtPath)
			tlsConfig.Certificates = append(tlsConfig.Certificates, cnCert)
		}
	}

	// Deprecated from go v1.18
	//	tlsConfig.BuildNameToCertificate()

	return tlsConfig
}

func Migrate() {
	logger.LaunchLog("Migration started")
	model.Migration()
}
