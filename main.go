package main

import (
	"fmt"
	"log"

	"github.com/hawkaii/blogator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Db_Url)
	fmt.Println(cfg.Current_User_Name)

	cfg.SetUser("hawkaii")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg.Current_User_Name)
	fmt.Println(cfg.Db_Url)

}
