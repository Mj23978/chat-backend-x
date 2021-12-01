module github.com/mj23978/chat-backend-x

go 1.15

replace (
	go.etcd.io/etcd/api/v3 v3.5.0-pre => go.etcd.io/etcd/api/v3 v3.0.0-20210107172604-c632042bb96c
	go.etcd.io/etcd/pkg/v3 v3.5.0-pre => go.etcd.io/etcd/pkg/v3 v3.0.0-20210107172604-c632042bb96c
)

require (
	github.com/DataDog/datadog-go v4.2.0+incompatible // indirect
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/chuckpreslar/emission v0.0.0-20170206194824-a7ddd980baf9
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/google/uuid v1.1.2
	github.com/julienschmidt/httprouter v1.3.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/nats-io/nats.go v1.10.0
	github.com/opentracing/opentracing-go v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/ory/analytics-go/v4 v4.0.0
	github.com/ory/herodot v0.9.1
	github.com/ory/jsonschema/v3 v3.0.1
	github.com/pborman/uuid v1.2.0
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.2.1
	github.com/prometheus/client_golang v1.9.0
	github.com/rs/zerolog v1.20.0
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shirou/gopsutil v2.20.8+incompatible
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/gjson v1.3.2
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/tidwall/sjson v1.0.4
	github.com/uber/jaeger-client-go v2.22.1+incompatible
	github.com/urfave/negroni v1.0.0
	go.etcd.io/etcd/client/v3 v3.0.0-20210107172604-c632042bb96c
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/tools v0.0.0-20210105210202-9ed45478a130 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.28.0
	gopkg.in/square/go-jose.v2 v2.2.2
	gopkg.in/yaml.v2 v2.4.0
)
