package discovery

import (
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	log "github.com/mj23978/chat-backend-x/logger/zerolog"
	"github.com/mj23978/chat-backend-x/utils"
)

//ServiceRegistry lib
type ServiceRegistry struct {
	Scheme string
	etcd   *Etcd
	quit   chan bool
}

//Node service node info
type Node struct {
	ID   string
	Info map[string]string
}

//NewServiceRegistry ServiceRegistry factory method
func NewServiceRegistry(endpoints []string, scheme string) *ServiceRegistry {
	r := &ServiceRegistry{
		Scheme: scheme,
		quit:   make(chan bool, 1),
	}
	if etc, err := newEtcd(endpoints); err != nil {
		log.Errorf("NewRegistry.Etcd Error err=%v ", err)
	} else {
		etcd = etc
	}
	r.etcd = etcd
	return r
}

//RegisterServiceNode .
func (r *ServiceRegistry) RegisterServiceNode(serviceName string, node Node) error {
	if serviceName == "" {
		return fmt.Errorf("Service name must be non empty ")
	}
	if node.ID == "" {
		return fmt.Errorf("Node name must be non empty")
	}
	node.Info["ip"] = utils.GetInterfaceIP()
	go r.keepRegistered(serviceName, node)
	return nil
}

func (r *ServiceRegistry) keepRegistered(serviceName string, node Node) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			nodePath := r.Scheme + serviceName + "-" + node.ID
			err := r.etcd.keepWhile(nodePath, encode(node.Info), r.quit)
			if err != nil {
				log.Warnf("Registration got errors. Restarting. err=%s", err)
			} else {
				log.Infof("Node [%s] registration success!", nodePath)
				return
			}
		case <-r.quit:
			return
		}
	}
}

//GetServiceNodes returns a list of active service nodes
func (r *ServiceRegistry) GetServiceNodes(serviceName string) ([]Node, error) {
	rsp, err := r.etcd.GetResponseByPrefix(r.servicePath(serviceName))
	if err != nil {
		return nil, err
	}
	nodes := make([]Node, 0)
	if len(rsp.Kvs) == 0 {
		log.Debugf("No services nodes were found under %s", r.servicePath(serviceName))
		return nodes, nil
	}

	for _, n := range rsp.Kvs {
		node := Node{}
		node.ID = string(n.Key)
		node.Info = decode(n.Value)
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func encode(m map[string]string) string {
	if m != nil {
		b, _ := json.Marshal(m)
		return string(b)
	}
	return ""
}

func decode(ds []byte) map[string]string {
	if len(ds) > 0 {
		var s map[string]string
		err := json.Unmarshal(ds, &s)
		if err != nil {
			log.Errorf("service.decode err => %+v", err)
			return nil
		}
		return s
	}
	return nil
}

func (r *ServiceRegistry) servicePath(serviceName string) string {
	service := strings.Replace(serviceName, "/", "-", -1)
	return path.Join(r.Scheme, service)
}
