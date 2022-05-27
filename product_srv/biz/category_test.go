package biz

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/proto/pb"
	"testing"
)

func init() {

	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, 60833)
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
		ParentCategoryId: 5,
		Level:            3,
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res2)
	// 第三级
	res3, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛排",
		ParentCategoryId: 6,
		Level:            4,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res3)
	res4, err := client.CreateCategory(context.Background(), &pb.CategoryItemReq{
		Name:             "牛肋条",
		ParentCategoryId: 6,
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
		ParentCategoryId: 6,
		Level:            3,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(err)
	res2, err := client.DeleteCategory(context.Background(), &pb.CategoryDelReq{
		Id: res.Id,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res2)
}

func TestProductServer_UpdateCategory(t *testing.T) {
	res, err := client.UpdateCategory(context.Background(), &pb.CategoryItemReq{
		Id:   7,
		Name: "牛筋皮子",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_GetAllCategoryList(t *testing.T) {
	res, err := client.GetAllCategoryList(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_GetSubCategory(t *testing.T) {
	res, err := client.GetSubCategory(context.Background(), &pb.CategoriesReq{
		Id:    6,
		Level: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
