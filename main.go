package main

import (
	"cleanArchitechureGo/config"
	"cleanArchitechureGo/module/item/storage"
	ginItem "cleanArchitechureGo/module/item/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
	"github.com/samber/do"
)

func init() {
	var pathFileEnv = "local.env"
	err := godotenv.Load(pathFileEnv)
	if err != nil {
		return
	}
}

func main() {
	//Load environment variables
	config.LoadCfgEnv()
	var injector = do.New()

	// Register mysql for sql store
	do.Provide(injector, storage.NewMySQL)

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		item := v1.Group("/item")
		{
			item.POST("", ginItem.CreateNewItem(injector))
		}
	}
}
