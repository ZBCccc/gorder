package main

import (
	"github.com/ZBCccc/gorder-v2/common/config"
	"github.com/spf13/viper"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	log.Println(viper.Get("order.service-name"))
}
