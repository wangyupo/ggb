package main

import (
	"fmt"
	"github.com/wangyupo/gin-cli/global"
	"github.com/wangyupo/gin-cli/initialize"
	"github.com/wangyupo/gin-cli/migration"
	"github.com/wangyupo/gin-cli/router"
	"os"
)

func init() {
	initialize.LoadEnvVariable()
	initialize.ConnectDB()
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	r := router.InitRouter()

	if global.DB != nil {
		migration.Migrate()
		fmt.Print(123)
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	r.Run(os.Getenv("SERVER_PORT"))
}
