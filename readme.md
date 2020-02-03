## GRPC BEEGO

该项目基于BEEGO搭建，业务API基于标准的BEEGO，该项目采用了服务分离，通过GRPC调用miner_service_impl，miner_service_impl项目是GRPC的服务实现

## 实验的API地址  [查看](http://106.54.93.177:8081/v1/user/getall)

## 项目运行

go mod tidy && go mod vendor && go build -mod=vendor

## supersior

该项目采用supersior对进程进行管理，环境配置通过/conf/prod.app.conf这种方式，与beego保持一致，supervisor需要注意导入环境变量，environment=BEEGO_RUNMODE=prod

## 数据库ORM

miner_service_impl采用了gorm

## 封装了GO_SYSTEM包

包括了NET,CONFIG的封装