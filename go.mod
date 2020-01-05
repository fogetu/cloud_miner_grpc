module gitlab.33.cn/user_site/cloud_miner

replace gitlab.33.cn/site-res/go_system => github.com/fogetu/go_system v0.0.0-20200105054501-9d8c1acf1c3e

replace gitlab.33.cn/site-service/miner-service-intf => github.com/fogetu/miner_service_intf v0.0.0-20181011133534-7bec62e10828

replace github.com/fogetu/cloud_miner_grpc => ./

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/fogetu/cloud_miner_grpc v0.0.0-00010101000000-000000000000
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	gitlab.33.cn/site-res/go_system v0.0.0-20200105054501-9d8c1acf1c3e
	gitlab.33.cn/site-service/miner-service-intf v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
)
