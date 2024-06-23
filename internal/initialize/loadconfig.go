package initialize

import (
	"fmt"
	"repodinh/go-backend-api/global"

	"github.com/spf13/viper"
)

func LoadConfig()  {
	//load configuration
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config
	viper.SetConfigName("local") //ten file
	viper.SetConfigType("yaml")

	//read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration file: %w \n", err))
	}

	//read server configuration
	fmt.Println("server Port::", viper.GetInt("server.port"))
	fmt.Println("server Port::", viper.GetString("security.jwt.key"))

	//configure structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration file: %w \n", err)
	}
}