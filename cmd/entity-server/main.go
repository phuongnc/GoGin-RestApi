package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"shapeservice/cmd/entity-server/repository"
	"shapeservice/cmd/entity-server/router"
	"shapeservice/configs"
	"time"

	log "shapeservice/internal/pkg/logger"
)

func main() {

	fmt.Println(reflect.TypeOf(10.5).Name())
	//init configuration
	wd, _ := os.Getwd()
	configPath := filepath.Join(wd, "configs/config.toml")
	cf := configs.InitFromFile(configPath)
	//init logger
	_ = log.InitLogger()

	//init database connection
	dbPath := filepath.Join(wd, cf.DBConnection)
	closeFunc, _ := repository.InitFromSQLLite(dbPath)

	//init router
	routersInit := router.InitRouter(cf.EnvPrefix)

	//init enpoint
	endPoint := fmt.Sprintf(":%d", cf.ServerPort)
	readTimeout := time.Minute
	writeTimeout := time.Minute
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.GetLogger().Info("[info] start http server listening %s", endPoint)
	_ = server.ListenAndServe()
	defer closeFunc()
}
