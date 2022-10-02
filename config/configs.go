package config

import (
	"errors"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type DDNSConfig struct {
	AccessKey       string `json:"access_key,omitempty"`
	AccessKeySecret string `json:"access_key_secret,omitempty"`
	RegionId        string `json:"region_id,omitempty"`
	DomainName      string `json:"domain_name,omitempty"`
	Type            string `json:"type,omitempty"`
	RR              string `json:"rr,omitempty"`
	NoticeUrl       string `json:"notice_url,omitempty"`
	ServerAddr      string `json:"server_addr,omitempty"`
}

var (
	DDNSConf *DDNSConfig
	one      sync.Once
)

func New() *DDNSConfig {
	if DDNSConf == nil {
		one.Do(func() {
			DDNSConf = &DDNSConfig{
				AccessKey:       viper.GetString("access_key"),
				AccessKeySecret: viper.GetString("access_key_secret"),
				RegionId:        viper.GetString("region_id"),
				DomainName:      viper.GetString("domain_name"),
				Type:            viper.GetString("type"),
				RR:              viper.GetString("rr"),
				NoticeUrl:       viper.GetString("notice_url"),
				ServerAddr:      viper.GetString("server_addr"),
			}
		})
	}
	return DDNSConf
}

// 导出配置
// path: 导出配置路径
func (d *DDNSConfig) Export(path string) error {
	err := viper.WriteConfigAs(path)
	if err != nil {
		log.Printf("viper.WriteConfigAs %s err: %v", path, err)
	}
	return nil
}

func Get() *DDNSConfig {
	return DDNSConf
}

func init() {
	viper.AutomaticEnv()
	loadDefaultConfig()
	// 加载配置文件
	if err := loadConfigWithFile(); err != nil && errors.Is(err, viper.ConfigFileNotFoundError{}) {
		log.Printf("load config from config file error, err msg: %v", err)
	}
}

func loadDefaultConfig() {
	viper.SetDefault("access_key", "YOUR ACCESS KEY ID")
	viper.SetDefault("access_key_secret", "YOUR ACCESS KEY SECRET")
	viper.SetDefault("region_id", "cn-hangzhou")
	viper.SetDefault("domain_name", "example.com")
	viper.SetDefault("type", "A")
	viper.SetDefault("rr", "@")
	viper.SetDefault("notice_url", "")
	viper.SetDefault("server_addr", "")
}

func loadConfigWithFile() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/fddns/")
	viper.AddConfigPath("$HOME/.fddns")
	viper.SetConfigType("json")
	viper.SetConfigType("yaml")
	viper.SetConfigType("toml")
	viper.SetConfigType("ini")
	// 读取配置
	return viper.ReadInConfig()
}
