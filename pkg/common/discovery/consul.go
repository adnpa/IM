package discovery

import (
	"github.com/hashicorp/consul/api"
)

type Discovery interface {
}

func Register(addr string, port int, name string, tags []string, id string) {
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	check := api.AgentServiceCheck{
		HTTP:                           "http://local:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registeration := api.AgentServiceRegistration{
		ID:      id,
		Address: addr,
		Port:    port,
		Name:    name,
		Tags:    tags,
		Check:   &check,
	}

	err = cli.Agent().ServiceRegister(&registeration)
	if err != nil {
		panic(err)
	}
}

func AllService() map[string]*api.AgentService {
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := cli.Agent().Services()
	if err != nil {
		panic(err)
	}
	return data
}

func FilterService(name string) {

	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	cli, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := cli.Agent().ServicesWithFilter(`Service == ""`)
	if err != nil {
		panic(err)
	}
	return data
}

func UnRegister() {
	Register("loaclhost", 11000, "user-web", []string{"maxsho", "bobby"}, "user-web")
}
