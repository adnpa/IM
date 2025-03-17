# go-base

* Etcd 用于rpc服务的服务发现和注册，etcd Schema是注册名前缀，建议改为公司名，支持集群部署
* MySQL 用于消息和用户关系数据的全量存储
* MongoDB 用于离线消息存储，默认存储7天
* Kafka 用于消息转发，支持集群部署

https://docs.openim.io/zh-Hans/restapi/introduction

gateway
包括rpc server和ws server，负责将请求转发到api服务和聊天服务

service列表
http
用户
好友
群聊

ws
聊天


注册中心 consul
配置中心 nacos

数据 
聊天数据 mongodb
缓存 redis


