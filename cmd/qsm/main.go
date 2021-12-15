package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Jarover/qsm/cmd/qsm/config"
	"github.com/Jarover/qsm/cmd/qsm/routes"
	"github.com/Jarover/qsm/pkg/utils"
	"gopkg.in/natefinch/lumberjack.v2"

	"database/sql"

	_ "github.com/godror/godror"
)

// Читаем флаги и окружение
func readFlag(configFlag *config.Flag) {
	flag.StringVar(&configFlag.ConfigFile, "f", config.GetEnv("CONFIGFILE", utils.GetBaseFile()+".json"), "config file")
	//flag.StringVar(&configFlag.Host, "h", readconfig.GetEnv("HOST", ""), "host")
	flag.UintVar(&configFlag.Port, "p", uint(config.GetEnvInt("PORT", 0)), "port")
	flag.Parse()

}

// Читаем конфиг

func getConfig(dir string) (*config.Config, error) {
	var configFlag config.Flag
	readFlag(&configFlag)

	fmt.Println(configFlag)
	fmt.Println(dir + "/" + configFlag.ConfigFile)

	config, err := config.ReadConfig(dir + "/" + configFlag.ConfigFile)
	if configFlag.Port != 0 {

		config.Port = configFlag.Port
	}

	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {

	start := time.Now()

	dir := utils.GetDir()

	logPath := dir + "/" + utils.GetBaseFile() + "_app.log"

	l := &lumberjack.Logger{ //nolint:typecheck
		Filename:   logPath,
		MaxSize:    500, // megabytes
		MaxBackups: 10,
		MaxAge:     1,     //days
		Compress:   false, // disabled by default
	}
	log.SetOutput(l)
	log.Println("Start program")
	log.Println(dir)

	err := config.Version.ReadVersionFile(dir + "/version.json")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(config.Version)

	Config, err := getConfig(dir)

	log.Println(Config)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(sql.Drivers())

	db, err := sql.Open("godror", `user="sysadm" password="sysadm" connectString="xeon2vm:1521/neva"`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	r := routes.SetupRouter()

	r.LoadHTMLGlob(dir + "/templates/*")
	r.Run(":" + strconv.FormatUint(uint64(Config.Port), 10)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Printf("%v: %v\n", "Время работы программы", time.Since(start))
}
