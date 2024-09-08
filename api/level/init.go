package lcli

import "github.com/spf13/viper"

func Init() {
	Start_port = viper.GetInt("level.start_port")
	End_port = viper.GetInt("level.end_port")
	now_port = Start_port
}
