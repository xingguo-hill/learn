package x

import (
	"fmt"
	"testing"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func newClientConfig() config_client.IConfigClient {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		//命名空间id需要提前创建,置空为public
		NamespaceId:         "2619afdf-e57a-4469-96fb-6740e392b1ec",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "192.168.56.29",
			Port:   8848,
		},
		{
			IpAddr: "192.168.56.30",
			Port:   8848,
		},
		{
			IpAddr: "192.168.56.31",
			Port:   8848,
		},
	}

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)

	if err != nil {
		panic(err)
	}
	return client
}

func TestNacosClientConfig(t *testing.T) {
	client := newClientConfig()
	success, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "push_config_test",
		Group:   "test",
		Content: "hello world!",
		Type:    "text",
		AppName: "appName",
		SrcUser: "xg",
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err)
	}
	println(success)
}
func TestNacosClientDiscovery(t *testing.T) {

	// // 配置Nacos服务地址等信息
	// serverConfigs := []constant.ServerConfig{
	// 	{
	// 		IpAddr: "192.168.56.30",
	// 		Port:   8848,
	// 	},
	// 	{
	// 		IpAddr: "192.168.56.31",
	// 		Port:   8848,
	// 	},
	// }

}
