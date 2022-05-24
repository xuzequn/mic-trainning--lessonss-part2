package biz

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/proto/pb"
	"testing"
)

var client pb.ProductServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, internal.AppConf.ProductSrvConfig.Port)
	//grpc.Dial(addr, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = pb.NewProductServiceClient(conn)
}

func TestProductServer_CreateBrand(t *testing.T) {
	brands := []string{
		"大希地", "恒都", "小牛凯西",
	}
	for _, item := range brands {
		res, err := client.CreateBrand(context.Background(), &pb.BrandItemReq{
			Name: item,
			Logo: "www.baidu.com",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}

}

func TestProductServer_BandList(t *testing.T) {

}
