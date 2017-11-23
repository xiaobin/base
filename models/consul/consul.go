package consul

import (
	"github.com/astaxie/beego"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
)


func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UP")
}

func Print(){
	beego.Debug("go,consul")
}

func init() {
	beego.Debug("初始化consul")
	config := consulapi.DefaultConfig()
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("consul client error : ", err)
	}
	checkPort, _ := beego.AppConfig.Int("consul.check-port")
	checkUri := beego.AppConfig.String("consul.check")
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = beego.AppConfig.String("consul.app-id")
	registration.Name = beego.AppConfig.String("consul.app-name")
	registration.Port, _ = beego.AppConfig.Int("consul.register.port")

	registration.Tags = []string{beego.AppConfig.String("")}
	registration.Address = beego.AppConfig.String("consul.register.ip")
	registration.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, checkUri),
		Timeout:                        "3s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s", //check失败后30秒删除本服务
	}

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		beego.Error("register server error : ", err)
	}

	http.HandleFunc(checkUri, consulCheck)
	go http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)
}

