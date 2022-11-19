package cmd

import (
	"context"
	user "eshop/api/user/v1"
	"eshop/app/user/config"
	"eshop/app/user/service"
	"flag"
	"github.com/go-kratos/kratos/v2/log"

	"google.golang.org/grpc"
	"net"
	"reflect"
	"strings"
)

func AddNewRequestTraceLog(request interface{}) {
	fields := make(map[string]interface{})
	e := reflect.ValueOf(request).Elem()
	for i := 0; i < e.NumField(); i++ {
		if !e.Field(i).CanInterface() {
			continue
		}
		varName := e.Type().Field(i).Name
		varValue := e.Field(i).Interface()

		if strings.HasPrefix(varName, "XXX_") == true {
			continue
		}
		fields[varName] = varValue
	}
	requestType := reflect.TypeOf(request).String()
	traceLogStr := "Receive one new " + requestType
	log.Debugf("fields:%+v,%s", fields, traceLogStr)
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	AddNewRequestTraceLog(req)
	return handler(ctx, req)
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config-path", configPath, "service.yaml")
	flag.Parse()
	if len(configPath) == 0 {
		log.Fatal("config path is empty")
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("load config failed,err:%s", err.Error())
	}
	srv := service.New(cfg)
	if err = srv.Initialize(); err != nil {
		log.Fatalf("init server failed,err:%s", err.Error())
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor))
	grpcSrv := grpc.NewServer(opts...)
	lis, err := net.Listen("tcp", cfg.Service.Host)
	//todo: 注册etcd
	if err != nil {
		log.Fatalf("listen server failed,err:%s", err.Error())
	}

	user.RegisterUserServer(grpcSrv, srv)
}
