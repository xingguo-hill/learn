package x

import (
	"fmt"
	"testing"
	"time"

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
		TimeoutMs:           10000,
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

func getTestDataIdAndGroup() (DataId string, Group string) {
	DataId = "push_config_test"
	Group = "test"
	return DataId, Group
}

func TestNacosClientPublishConfig(t *testing.T) {
	client := newClientConfig()
	DataId, Group := getTestDataIdAndGroup()
	success, err := client.PublishConfig(vo.ConfigParam{
		DataId:  DataId,
		Group:   Group,
		Content: "hello world!",
		Type:    "text",
		SrcUser: "xg",
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err)
	}
	fmt.Println(success)
	time.Sleep(time.Second * 3)
}

func TestNacosClientDeleteConfig(t *testing.T) {
	client := newClientConfig()
	DataId, Group := getTestDataIdAndGroup()
	success, err := client.DeleteConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err != nil {
		fmt.Printf("TestNacosClientDeleteConfig err:%+v \n", err)
	}
	fmt.Println(success)
}

func TestNacosClientGetConfig(t *testing.T) {
	client := newClientConfig()
	DataId, Group := getTestDataIdAndGroup()
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err != nil {
		fmt.Printf("TestNacosClientGetConfig err:%+v \n", err)
	}
	fmt.Println(content)
}

/*
*
可结合管理页面更新测试，如果还是通过TestNacosClientPublishConfig测试，监听始终是保持的
*/
func TestNaocosListenConfig(t *testing.T) {
	client := newClientConfig()
	DataId, Group := getTestDataIdAndGroup()
	chMsg := make(chan string)
	err := client.ListenConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
		OnChange: func(namespace, group, dataId, data string) {
			chMsg <- fmt.Sprintf("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	if err != nil {
		fmt.Printf("TestNacosClientGetConfig err:%+v \n", err)
	}
	cnt := 0
	// listen message
	for {
		select {
		case msg := <-chMsg:
			fmt.Println(msg)
			cnt++
			if cnt > 2 {
				err := client.CancelListenConfig(vo.ConfigParam{
					DataId: DataId,
					Group:  Group,
				})
				if err != nil {
					fmt.Printf("CancelListenConfig err:%+v \n", err)
				}
				fmt.Println("CancelListenConfig finish!")
				return
			}
		}
	}
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
