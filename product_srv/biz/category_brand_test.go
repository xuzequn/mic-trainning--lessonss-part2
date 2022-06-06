package biz

import (
	"context"
	"fmt"
	"mic-trainning-lessons-part2/proto/pb"
	"testing"
)

func TestProductServer_CreateCategoryBrand(t *testing.T) {
	res, err := client.CreateCategoryBrand(context.Background(), &pb.CategoryBrandReq{
		Brand:    &pb.BrandItemRes{Id: 6},
		Category: &pb.CategoryItemRes{Id: 6},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_UpdateCategoryBrand(t *testing.T) {
	res, err := client.UpdateCategoryBrand(context.Background(), &pb.CategoryBrandReq{
		Id:       1,
		Brand:    &pb.BrandItemRes{Id: 1},
		Category: &pb.CategoryItemRes{Id: 6},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_DeleteCategoryBrand(t *testing.T) {
	res, err := client.DeleteCategoryBrand(context.Background(), &pb.CategoryBrandReq{
		Id: 3,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

}

func TestProductServer_CategoryBrandList(t *testing.T) {
	res, err := client.CategoryBrandList(context.Background(), &pb.PagingReq{
		PageNo:   1,
		PageSize: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_GetCategoryBrandList(t *testing.T) {
	res, err := client.GetCategoryBrandList(context.Background(), &pb.CategoryItemReq{
		Id:               6,
		ParentCategoryId: 5,
		Level:            2,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
