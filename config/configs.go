package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type DDNSConfig struct {
	AccessKey       string `json:"AK,omitempty"`
	AccessKeySecret string `json:"AKS,omitempty"`
	RegionId        string `json:"RegionId,omitempty"`
	DomainName      string `json:"DomainName,omitempty"`
	Type            string `json:"Type,omitempty"`
	RR              string `json:"RR,omitempty"`
	NoticeUrl       string `json:"url,omitempty"`
	ServerAddr      string `json:"server_addr,omitempty"`
}

var (
	DDNSConf *DDNSConfig
	one      sync.Once
)

func New(ak, aks, region, domainName, t, rr, url string) {
	if DDNSConf == nil {
		one.Do(func() {
			DDNSConf = &DDNSConfig{
				AccessKey:       ak,
				AccessKeySecret: aks,
				RegionId:        region,
				DomainName:      domainName,
				Type:            t,
				RR:              rr,
				NoticeUrl:       url,
			}
		})
	}
}

func Get() *DDNSConfig {
	return DDNSConf
}

func Save(filename, fileType string) {
	viper.WriteConfigAs(fmt.Sprintf("%s.%s", filename, fileType))
}

func loadDefaultConfig() {
	viper.SetDefault("RR", "@")
	viper.SetDefault("Type", "A")
	viper.SetDefault("RegionId", "cn-hangzhou")
	viper.SetDefault("DomainName", "example.com")
	viper.SetDefault("AK", "YOUR ACCESS KEY ID")
	viper.SetDefault("AKS", "YOUR ACCESS KEY SECRET")
	viper.SetDefault("url", "")
}

func loadConfigWithFile() error {
	viper.SetConfigName("config") //
	//viper.AddConfigPath("/etc/appname/")   //
	viper.SetConfigType("json")
	//viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	// 读取配置
	return viper.ReadInConfig()
}

func init() {
	// 加载默认配置项
	loadDefaultConfig()
	// 加载配置文件
	_ = loadConfigWithFile()
	// 加载环境变量
	viper.AutomaticEnv()
	// 初始化配置
	New(viper.GetString("AK"),
		viper.GetString("AKS"),
		viper.GetString("RegionId"),
		viper.GetString("DomainName"),
		viper.GetString("Type"),
		viper.GetString("RR"),
		viper.GetString("url"))
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
