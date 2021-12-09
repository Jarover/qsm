package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Jarover/qsm/cmd/qsm/config"
	"github.com/Jarover/qsm/pkg/utils"
)

func main() {

	start := time.Now()

	dir := utils.GetDir()
	err := config.Version.ReadVersionFile(dir + "/version.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Version)

	log.Printf("%v: %v\n", "Время работы программы", time.Since(start))
}
