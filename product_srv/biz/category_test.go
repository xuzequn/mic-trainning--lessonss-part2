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

func init() {

	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, 61051)
	//grpc.Dial(addr, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = pb.NewProductServiceClient(conn)
}

func TestProductServer_CreateCategory(t *testing.T) {
	// 第一级
	res, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "鲜肉",
		ParentCategoryId: 1,
		Level:            2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
	// 第二级
	res2, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛肉",
		ParentCategoryId: 18,
		Level:            3,
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res2)
	// 第三级
	res3, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛排",
		ParentCategoryId: 25,
		Level:            4,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res3)
	res4, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛肋条",
		ParentCategoryId: 25,
		Level:            4,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res4)
}

func TestProductServer_DeleteCategory(t *testing.T) {
	res, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "ceshi",
		ParentCategoryId: 25,
		Level:            3,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(err)
	client.DeleteCategory(context.Background(), &pb.CategoryDelReq{
		Id: res.Id,
	})
}
