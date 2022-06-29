package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type DDNSConfig struct {
	AccessKey       string `json:"AK"`
	AccessKeySecret string `json:"AKS"`
	RegionId        string `json:"RegionId"`
	DomainName      string `json:"DomainName"`
	Type            string `json:"Type"`
	RR              string `json:"RR"`
}

var DDNSConf *DDNSConfig

func init() {
	viper.SetConfigName("config") //指定配置文件的文件名称(不需要制定配置文件的扩展名)
	//viper.AddConfigPath("/etc/appname/")   //设置配置文件的搜索目录
	viper.SetConfigType("json")
	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(".")    // 设置配置文件和可执行二进制文件在用一个目录
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("读取配置文件出错: %v, 请参考 example.json 配置了 config.json 文件", err)
		}
	}

	viper.AutomaticEnv()

	viper.SetDefault("RegionId", "cn-hangzhou")
	DDNSConf = &DDNSConfig{
		AccessKey:       viper.GetString("AK"),
		AccessKeySecret: viper.GetString("AKS"),
		RegionId:        viper.GetString("RegionId"),
		DomainName:      viper.GetString("DomainName"),
		Type:            viper.GetString("Type"),
		RR:              viper.GetString("RR"),
	}

	// fmt.Println("获取配置文件的map[string]string", viper.GetStringMapString(`app`))
}

// 保存用户配置
func (d *DDNSConfig) Save(filename string) error {
	// 序列化
	ddnsEncodes, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return err
	}

	// 保存到文件中
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	wn, err := f.Write(ddnsEncodes)
	if err != nil {
		return err
	}

	if wn != len(ddnsEncodes) {
		return fmt.Errorf("保存文件出错，请检查文件")
	}
	return nil
}
