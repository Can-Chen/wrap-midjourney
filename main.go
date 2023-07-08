package main

import (
	"fmt"
	"wrap-midjourney/handlers"
	"wrap-midjourney/initialization"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	cfg := pflag.StringP("config", "c", "./config.yaml", "api server config file path.")

	pflag.Parse()

	_, err := initialization.LoadConfig(*cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	initialization.LoadDiscordClient(handlers.DiscordMsgCreate, handlers.DiscordMsgUpdate)

	r := gin.Default()

	r.POST("/v1/trigger/midjourney-bot", handlers.MidjourneyBot)
	r.POST("/v1/trigger/upload", handlers.UploadFile)

	err = r.Run(":16007")
	if err != nil {
		fmt.Println(err)
	}
}
