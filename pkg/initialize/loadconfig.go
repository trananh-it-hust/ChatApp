package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/trananh-it-hust/ChatApp/global"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	// Log the loaded configuration
	fmt.Println("Port started at: ", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

}
