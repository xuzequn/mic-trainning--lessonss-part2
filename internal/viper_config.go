package internal

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var AppConf AppConfig
var NacosConf NacosConfig

//var ViperConf ViperConfig
var fileName = "dev-config.yaml"

//var fileName = "../../dev-config.yaml"

func initNacos() {
	v := viper.New()
	v.SetConfigFile(fileName)
	v.ReadInConfig()
	err := v.Unmarshal(&NacosConf)
	fmt.Println(NacosConf)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func initFromNacos() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: NacosConf.Host,
			Port:   NacosConf.Port,
		},
	}
	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConf.NameSpace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConf.DataId,
		Group:  NacosConf.Group,
	})
	if err != nil {
		panic(err)
	}
	if content == "" {
		panic("配置文件为空")
	}
	fmt.Println(content)
	json.Unmarshal([]byte(content), &AppConf)
	fmt.Println(AppConf)
}

func init() {
	initNacos()
	initFromNacos()
	fmt.Println("config初始化完成。。。")
	InitRedis()
}
