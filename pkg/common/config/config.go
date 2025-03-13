package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	ServerIP      string `yaml:"serverip"`
	ServerVersion string `yaml:"serverversion"`

	Mysql struct {
		DBAddress      []string `yaml:"dbMysqlAddress"`
		DBUserName     string   `yaml:"dbMysqlUserName"`
		DBPassword     string   `yaml:"dbMysqlPassword"`
		DBDatabaseName string   `yaml:"dbMysqlDatabaseName"`
		DBTableName    string   `yaml:"DBTableName"`
		DBMsgTableNum  int      `yaml:"dbMsgTableNum"`
		DBMaxOpenConns int      `yaml:"dbMaxOpenConns"`
		DBMaxIdleConns int      `yaml:"dbMaxIdleConns"`
		DBMaxLifeTime  int      `yaml:"dbMaxLifeTime"`
	}
	Mongo struct {
		DBAddress           []string `yaml:"dbAddress"`
		DBDirect            bool     `yaml:"dbDirect"`
		DBTimeout           int      `yaml:"dbTimeout"`
		DBDatabase          string   `yaml:"dbDatabase"`
		DBSource            string   `yaml:"dbSource"`
		DBUserName          string   `yaml:"dbUserName"`
		DBPassword          string   `yaml:"dbPassword"`
		DBMaxPoolSize       int      `yaml:"dbMaxPoolSize"`
		DBRetainChatRecords int      `yaml:"dbRetainChatRecords"`
	}
	Etcd struct {
		EtcdSchema string   `yaml:"etcdSchema"`
		EtcdAddr   []string `yaml:"etcdAddr"`
	}

	RpcPort struct {
		UserPort              []int `yaml:"UserPort"`
		FriendPort            []int `yaml:"FriendPort"`
		RpcMessagePort        []int `yaml:"rpcMessagePort"`
		RpcPushMessagePort    []int `yaml:"rpcPushMessagePort"`
		GroupPort             []int `yaml:"GroupPort"`
		RpcModifyUserInfoPort []int `yaml:"rpcModifyUserInfoPort"`
		RpcGetTokenPort       []int `yaml:"rpcGetTokenPort"`
	}
	RpcRegisterName struct {
		UserName               string `yaml:"UserName"`
		FriendName             string `yaml:"FriendName"`
		OfflineMessageName     string `yaml:"OfflineMessageName"`
		PushName               string `yaml:"PushName"`
		OnlineMessageRelayName string `yaml:"OnlineMessageRelayName"`
		GroupName              string `yaml:"GroupName"`
		AuthName               string `yaml:"AuthName"`
	}

	ModuleName struct {
		LongConnSvrName string `yaml:"longConnSvrName"`
		MsgTransferName string `yaml:"msgTransferName"`
		PushName        string `yaml:"pushName"`
	}

	Jwt struct {
		Secret string `yaml:"secret"`
		Expire int64  `yaml:"expire"`
	}

	Kafka struct {
		Ws2mschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		Ms2pschat struct {
			Addr  []string `yaml:"addr"`
			Topic string   `yaml:"topic"`
		}
		ConsumerGroupID struct {
			MsgToMongo string `yaml:"msgToMongo"`
			MsgToMySql string `yaml:"msgToMySql"`
			MsgToPush  string `yaml:"msgToPush"`
		}
	}

	LongConnSvr struct {
		WebsocketPort       []int `yaml:"WsPort"`
		WebsocketMaxConnNum int   `yaml:"websocketMaxConnNum"`
		WebsocketMaxMsgLen  int   `yaml:"websocketMaxMsgLen"`
		WebsocketTimeOut    int   `yaml:"websocketTimeOut"`
	}
}

func init() {
	//path, _ := os.Getwd()
	//log.Println(path)
	cfgFile, err := os.Open("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(cfgFile)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err)
	}
}
