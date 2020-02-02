module cloud_miner_grpc

replace github.com/fogetu/cloud_miner_grpc => ./

go 1.13

require (
	github.com/astaxie/beego v1.12.0
	github.com/fogetu/cloud_miner_grpc v0.0.0-00010101000000-000000000000
	github.com/fogetu/go_system v0.0.0-20200117053147-e4826994dde0
	github.com/fogetu/miner_service_intf v0.0.0-20200113151624-96513cf11c88
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553
	google.golang.org/grpc v1.26.0
)
