package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/product_web/handler"
	"mic-trainning-lessons-part2/util"
)

func init() {
	fmt.Println(internal.AppConf)
	err := internal.Reg(internal.AppConf.ProductWebConfig.Host,
		internal.AppConf.ProductWebConfig.SrvName,
		internal.AppConf.ProductWebConfig.SrvName,
		internal.AppConf.ProductWebConfig.Port,
		internal.AppConf.ProductWebConfig.Tags)
	if err != nil {
		panic(err)
	}

}

func main() {
	//addrs, err := net.InterfaceAddrs()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(addrs)
	//ip := flag.String("ip", "0.0.0.0", "输入Ip")
	//port := flag.Int("port", 8081, "输入端口")
	//flag.Parse()
	//addr := fmt.Sprintf("%s:%d", *ip, *port)
	ip := internal.AppConf.ProductWebConfig.Host
	port := util.GenRandomPort()
	if internal.AppConf.Debug {
		port = internal.AppConf.ProductWebConfig.Port
	}
	addr := fmt.Sprintf("%s:%d", ip, port)
	r := gin.Default()
	productGroup := r.Group("/v1/product")
	{
		productGroup.GET("/list", handler.ProductListHandler)
	}
	r.GET("/health", handler.HealthHandler)
	r.Run(addr)
}
