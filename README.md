# go-base

* Etcd 用于rpc服务的服务发现和注册，etcd Schema是注册名前缀，建议改为公司名，支持集群部署
* MySQL 用于消息和用户关系数据的全量存储
* MongoDB 用于离线消息存储，默认存储7天
* Kafka 用于消息转发，支持集群部署

https://docs.openim.io/zh-Hans/restapi/introduction