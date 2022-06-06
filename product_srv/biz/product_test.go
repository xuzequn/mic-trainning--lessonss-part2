package biz

import (
	"context"
	"fmt"
	"mic-trainning-lessons-part2/proto/pb"
	"testing"
)

func TestProductServer_CreateProduct(t *testing.T) {
	for i := 15; i < 20; i++ {
		res, err := client.CreateProduct(context.Background(), &pb.CreateProductItem{
			Name:        fmt.Sprintf("黄金牛排%d", i),
			Sn:          "12345678",
			Stocks:      "",
			Price:       359.00,
			RealPrice:   199.00,
			ShortDesc:   "",
			ProductDesc: "",
			Images:      nil,
			DescImages:  nil,
			CoverImage:  "https://www.baidu.com",
			IsNew:       false,
			IsPop:       false,
			Selling:     false,
			BrandId:     6,
			FavNUm:      6666,
			SoldNum:     5432,
			CategoryId:  6,
			IsShipFree:  false,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	}
}
func TestProductServer_UpdateProduct(t *testing.T) {
	res, err := client.UpdateProduct(context.Background(), &pb.CreateProductItem{
		Id:         15,
		Name:       "战斧牛排666",
		BrandId:    6,
		CategoryId: 6,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_DeleteProduct(t *testing.T) {
	res, err := client.DeleteProduct(context.Background(), &pb.ProductDelItem{Id: 10})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_BatchGetProduct(t *testing.T) {
	ids := []int32{10, 11, 12}
	res, err := client.BatchGetProduct(context.Background(), &pb.BatchProductIdReq{Ids: ids})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_ProductList(t *testing.T) {
	res, err := client.ProductList(context.Background(), &pb.ProductConditionReq{
		PageNo:   3,
		PageSize: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

}

func TestProductServer_GetProductDetail(t *testing.T) {
	res, err := client.GetProductDetail(context.Background(), &pb.ProductItemReq{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
