package configuration

import (
	"strings"
	"time"
	"github.com/QubitProducts/bamboo/Godeps/_workspace/src/github.com/samuel/go-zookeeper/zk"
)

/*
	Mesos Marathon configuration
*/
type Marathon struct {
	// comma separated marathon http endpoints including port number
	Endpoint string
	UseZookeeper bool
	Zookeeper Zookeeper
	User string
	Password string
	UseEventStream bool
}

func (m Marathon) Endpoints() []string {
	if (m.UseZookeeper) {
		endpoints, err := _zkEndpoints(m.Zookeeper)
		if err != nil {
			return []string{}
		}
		return endpoints
	} else {
		return strings.Split(m.Endpoint, ",")
	}

}

func _zkEndpoints(zkConf Zookeeper) ([]string, error) {
	// Only tested with marathon 0.11.1

	// TODO might want to reuse ZK Connection?
	var scheme = "http://"

	conn, _, err := zk.Connect(zkConf.ConnectionString(), time.Second*10)

	if err != nil {
		return nil, err
	}

	keys, _, err2 := conn.Children(zkConf.Path + "/leader")

	if err2 != nil {
		return nil, err2
	}

	endpoints := make([]string, 0, len(keys))

	for _, childPath := range keys {
		data, _, e := conn.Get(zkConf.Path + "/leader" + "/" + childPath)
		if e != nil {
			return nil, e
		}
		// TODO configurable http://??
		endpoints = append(endpoints, scheme + string(data))
	}
	return endpoints, nil
}
