package x

import (
	"fmt"
	"testing"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func getTestServiceInfo() (ServiceName string, GroupName string, ClusterName string) {
	ServiceName = "nacos.service.test"
	GroupName = "test"
	ClusterName = "nacos-cluster-test"
	return ServiceName, GroupName, ClusterName
}
func newNamingClient() naming_client.INamingClient {
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
	client, err := clients.NewNamingClient(
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

func TestBathRegisterNacosServiceInstance(t *testing.T) {
	client := newNamingClient()
	ServiceName, GroupName, ClusterName := getTestServiceInfo()
	//这里选nacos服务数据是用来注册便于测试数据，并非直接使用
	//批处理不支持持久实例
	success, err := client.BatchRegisterInstance(vo.BatchRegisterInstanceParam{
		ServiceName: ServiceName,
		GroupName:   GroupName,
		Instances: []vo.RegisterInstanceParam{
			{
				Ip:          "192.168.56.29",
				Port:        8848,
				Weight:      9,
				Enable:      true,
				Healthy:     true,
				Ephemeral:   true,        //临时实例，默认为临时服务
				ClusterName: ClusterName, //only 0-9a-zA-Z-
				Metadata:    map[string]string{"idc": "shanghai"},
			},
			{
				Ip:          "192.168.56.30",
				Port:        8848,
				Weight:      7,
				Enable:      true,
				Healthy:     true,
				Ephemeral:   true,
				ClusterName: ClusterName,
				Metadata:    map[string]string{"idc": "shanghai"},
			},
			{
				Ip:          "192.168.56.33",
				Port:        8848,
				Weight:      8,
				Enable:      true,
				Healthy:     true,
				Ephemeral:   true,
				ClusterName: ClusterName,
				Metadata:    map[string]string{"idc": "shanghai"},
			},
		},
	})
	if err != nil {
		fmt.Printf("BatchRegisterInstance err:%+v \n", err)
	}
	fmt.Printf("BatchRegisterInstance :%+v \n", success)
	time.Sleep(time.Second * 3600)
}

func TestRegisterNacosServiceInstance(t *testing.T) {
	ips := []string{"192.168.56.29", "192.168.56.30"}
	client := newNamingClient()
	ServiceName, GroupName, ClusterName := getTestServiceInfo()
	Ephemeral := false
	for _, ip := range ips {
		//这里选nacos服务数据是用来注册便于测试数据，并非直接使用
		success, err := client.RegisterInstance(
			vo.RegisterInstanceParam{
				ServiceName: ServiceName,
				GroupName:   GroupName,
				Ip:          ip,
				Port:        8848,
				Weight:      6,
				Enable:      true,
				Healthy:     true,
				Ephemeral:   Ephemeral,   //永久实例，默认为临时服务
				ClusterName: ClusterName, //only 0-9a-zA-Z-
				Metadata:    map[string]string{"idc": "bj1"},
			},
		)
		if err != nil {
			fmt.Printf("RegisterInstance err:%+v \n", err)
		}
		fmt.Printf("RegisterInstance :%+v \n", success)
	}
	if Ephemeral == true {
		time.Sleep(time.Second * 3600)
	}
}

func TestUpdateRegisterNacosServiceInstance(t *testing.T) {
	ips := []string{"192.168.56.29", "192.168.56.30"}
	client := newNamingClient()
	ServiceName, GroupName, ClusterName := getTestServiceInfo()
	Ephemeral := false
	for _, ip := range ips {
		//这里选nacos服务数据是用来注册便于测试数据，并非直接使用
		success, err := client.UpdateInstance(
			vo.UpdateInstanceParam{
				ServiceName: ServiceName,
				GroupName:   GroupName,
				Ip:          ip,
				Port:        8848,
				Weight:      3,
				Enable:      true,
				Healthy:     true,
				Ephemeral:   Ephemeral,   //永久实例，默认为临时服务
				ClusterName: ClusterName, //only 0-9a-zA-Z-
				Metadata:    map[string]string{"idc": "bj3"},
			},
		)
		if err != nil {
			fmt.Printf("UpdateInstance err:%+v \n", err)
		}
		fmt.Printf("UpdateInstance :%+v \n", success)
	}
	if Ephemeral == true {
		time.Sleep(time.Second * 3600)
	}
}

func TestDeRegisterNacosServiceInstance(t *testing.T) {
	client := newNamingClient()
	ServiceName, GroupName, ClusterName := getTestServiceInfo()
	//这里选nacos服务数据是用来注册便于测试数据，集群节点配置信息要完全一致，否则无法取消注册
	success, err := client.DeregisterInstance(
		vo.DeregisterInstanceParam{
			Ip:          "192.168.56.30",
			Port:        8848,
			Cluster:     ClusterName, //only 0-9a-zA-Z-
			ServiceName: ServiceName,
			GroupName:   GroupName,
			Ephemeral:   false, //永久实例，默认为临时服务
		},
	)
	if err != nil {
		fmt.Printf("DeregisterInstance err:%+v \n", err)
	}
	fmt.Printf("DeregisterInstance :%+v \n", success)
}

func TestSubscribeNacosinfo(t *testing.T) {
	client := newNamingClient()
	ServiceName, GroupName, ClusterName := getTestServiceInfo()
	chMsg := make(chan string)
	//Subscribe key=serviceName+groupName+cluster
	subscribeParam := &vo.SubscribeParam{
		ServiceName: ServiceName,
		GroupName:   GroupName,
		Clusters:    []string{ClusterName},
		SubscribeCallback: func(services []model.Instance, err error) {
			fmt.Printf("callback return services:%s \n\n", util.ToJsonString(services))
			chMsg <- fmt.Sprintf("callback return services:%s \n\n", util.ToJsonString(services))
		},
	}
	//带阻塞，调整为使用gorouting异步模式
	go func() {
		err := client.Subscribe(subscribeParam)
		if err != nil {
			fmt.Printf("Subscribe err:%+v \n", err)
		}
	}()

	cnt := 1
	// listen message
	for {
		select {
		case msg := <-chMsg:
			fmt.Println(msg)
			cnt++
			if cnt > 2 {
				err := client.Unsubscribe(subscribeParam)
				if err != nil {
					fmt.Printf("Unsubscribe err:%+v \n", err)
				}
				fmt.Println("Unsubscribe finish!")
				// return
			}
		}
	}
}

func TestNacosGetInfos(t *testing.T) {
	client := newNamingClient()
	ServiceName, GroupName, ClusterName := getTestServiceInfo()
	//GetService
	param := vo.GetServiceParam{
		ServiceName: ServiceName,
		GroupName:   GroupName,
		Clusters:    []string{ClusterName},
	}
	service, err := client.GetService(param)
	if err != nil {
		panic("GetService failed!" + err.Error())
	}
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
	time.Sleep(time.Second)

	//SelectAllInstances
	aparam := vo.SelectAllInstancesParam{
		ServiceName: ServiceName,
		GroupName:   GroupName,
		Clusters:    []string{ClusterName},
	}
	instances, err := client.SelectAllInstances(aparam)
	if err != nil {
		panic("SelectAllInstances failed!" + err.Error())
	}
	fmt.Printf("SelectAllInstances,param:%+v, result:%+v \n\n", aparam, instances)
	time.Sleep(time.Second)

	//selectHealthyInstances
	hparam := vo.SelectInstancesParam{
		ServiceName: ServiceName,
		GroupName:   GroupName,
		Clusters:    []string{ClusterName},
		HealthyOnly: true,
	}
	instances, err = client.SelectInstances(hparam)
	if err != nil {
		panic("SelectInstances failed!" + err.Error())
	}
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", hparam, instances)
	time.Sleep(time.Second)

	//selectOneHealthyInstance
	oparam := vo.SelectOneHealthInstanceParam{
		ServiceName: ServiceName,
		GroupName:   GroupName,
		Clusters:    []string{ClusterName},
	}
	//SelectOneHealthyInstance return one instance by WRR strategy for load balance And the instance should be health=true,enable=true and weight>0
	instanceOne, err := client.SelectOneHealthyInstance(oparam)
	if err != nil {
		panic("SelectOneHealthyInstance failed!" + err.Error())
	}
	fmt.Printf("SelectOneHealthyInstance,param:%+v, result:%+v \n\n", oparam, instanceOne)
}
