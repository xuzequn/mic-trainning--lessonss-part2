package biz

import (
	"context"
	"fmt"
	"mic-trainning-lessons-part2/proto/pb"
	"testing"
)

var client pb.ProductServiceClient

//func init() {
//
//	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, 50773)
//	//grpc.Dial(addr, grpc.WithInsecure())
//	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		panic(err)
//	}
//	client = pb.NewProductServiceClient(conn)
//}

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
	brandRes, err := client.BandList(context.Background(), &pb.BrandPagingReq{
		PageNo:   1,
		PageSize: 3,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(brandRes)
}

func TestProductServer_DeleteBrand(t *testing.T) {
	res, err := client.DeleteBrand(context.Background(), &pb.BrandItemReq{
		Id: 3,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_UpdateBrand(t *testing.T) {
	res, err := client.UpdateBrand(context.Background(), &pb.BrandItemReq{
		Id:   4,
		Name: "测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
