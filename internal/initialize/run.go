package initialize

import (
	"fmt"
	"repodinh/go-backend-api/global"
)

func Run()  {
	//load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql...", m.UserName)
	r := InitRouter()

	r.Run(":8002")
}
