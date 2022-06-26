package config

import (
	"log"

	"github.com/spf13/viper"
)

type DDNSConfig struct {
	AccessKey       string
	AccessKeySecret string
	RegionId        string
	DomainName      string
	Type            string
	RR              string
}

var DDNSConf *DDNSConfig

func init() {
	viper.SetConfigName("config") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(".")    // 设置配置文件和可执行二进制文件在用一个目录
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("读取配置文件出错: %v, 请参考 example.json 配置了 config.json 文件", err)
		}
	}

	viper.AutomaticEnv()

	DDNSConf = &DDNSConfig{
		AccessKey:       viper.GetString("ak"),
		AccessKeySecret: viper.GetString("aks"),
		DomainName:      viper.GetString("DomainName"),
		Type:            viper.GetString("Type"),
		RR:              viper.GetString("RR"),
	}

	// fmt.Println("获取配置文件的map[string]string", viper.GetStringMapString(`app`))
}
