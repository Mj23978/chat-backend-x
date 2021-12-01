package discovery

import (
	"encoding/json"
	"sync"
	"testing"
	"time"

	nprotoo "github.com/mj23978/chat-backend-x/broker/nats"
	log "github.com/mj23978/chat-backend-x/logger/zerolog"
)

const (
	EtcdAddr = "http://127.0.0.1:2386"
	NatsAddr = "http://127.0.0.1:4223"
)

var (
	wg *sync.WaitGroup
)

func init() {
	log.Init("info")
	wg = new(sync.WaitGroup)
}

func JsonEncode(str string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}
	return data
}

func ServiceNodeRegistry() {
	serviceNode := NewServiceNode([]string{EtcdAddr}, "dc1")
	serviceNode.RegisterNode("game-server", "game-node-test")
	protoo := nprotoo.NewNatsProtoo(NatsAddr)
	wg.Add(1)
	protoo.OnRequest(serviceNode.GetRPCChannel(), func(request nprotoo.Request, accept nprotoo.RespondFunc, reject nprotoo.RejectFunc) {
		log.Infof("method => %s, data => %v", request.Method, request.Data)
		if request.Method == "offer" {
			accept("We Accept Your Offer")
		}
		reject(404, "Not found")
		wg.Done()
	})

	protoo.OnBroadcast(serviceNode.GetEventChannel(), func(data nprotoo.Notification, subj string) {
		log.Infof("Got Broadcast subj => %s, data => %v", subj, data)
		wg.Done()
	})

	wg.Add(1)
	broadcaster := protoo.NewBroadcaster(GetEventChannel(serviceNode.NodeInfo()))
	broadcaster.Say("foo", JsonEncode(`{"hello": "world"}`))
}

func ServiceNodeWatch() {
	serviceWatcher := NewServiceWatcher([]string{EtcdAddr}, "dc1")
	protoo := nprotoo.NewNatsProtoo(NatsAddr)
	go serviceWatcher.WatchServiceNode("game-server", func(service string, state NodeStateType, node Node) {
		if state == UP {
			log.Infof("Service UP [%s] => %v", service, node)
			req := protoo.NewRequestor(GetRPCChannel(node))
			wg.Add(1)
			req.Request("offer", JsonEncode(`{ "sdp": "dummy-sdp"}`),
				func(result nprotoo.RawMessage) {
					log.Infof("offer success: =>  %s", result)
					wg.Done()
				},
				func(code int, err string) {
					log.Warnf("offer reject: %d => %s", code, err)
					wg.Done()
				})
		} else if state == DOWN {
			log.Infof("Service DOWN [%s] => %v", service, node)
		}
	})
}

func TestServiceNode(t *testing.T) {
	ServiceNodeWatch()
	ServiceNodeRegistry()
	wg.Wait()
}

func WrongTestServiceNode(t *testing.T) {
	ServiceNodeWatch()
	ServiceNodeRegistry()
	wg.Wait()
}

func ServiceUpTimeNode() {
	serviceNode := NewServiceNode([]string{EtcdAddr}, "dc1")
	serviceNode.RegisterNode("game-server", "game-node-0")
	serviceNode2 := NewServiceNode([]string{EtcdAddr}, "dc1")
	serviceNode2.RegisterNode("game-server", "game-node-1")
	wg.Add(1)

	time.Sleep(time.Second * 40)

	res, _ := serviceNode.reg.GetServiceNodes("game-server")
	log.Infof("Nodes : %v", res)

	serviceNode.UnregisterNode()
	log.Infof("Node : %v", serviceNode.NodeInfo())
	time.Sleep(time.Second * 10)
	wg.Done()
}

func ServiceUpTimeWatch() {
	serviceWatcher := NewServiceWatcher([]string{EtcdAddr}, "dc1")
	go serviceWatcher.WatchServiceNode("game-server", func(service string, state NodeStateType, node Node) {
		if state == UP {
			log.Infof("Service UP [%s] => %v", service, node)
		} else if state == DOWN {
			log.Infof("Service DOWN [%s] => %v", service, node)
		}
	})
}

func TestUpTime(t *testing.T) {
	ServiceUpTimeWatch()
	ServiceUpTimeNode()
	wg.Wait()
}
