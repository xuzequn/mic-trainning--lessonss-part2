package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/custom_error"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/model"
	"mic-trainning-lessons-part2/proto/pb"
)

func (p ProductServer) CreateCategory(ctx context.Context, req *pb.CategoryItemReq) (*pb.CategoryItemRes, error) {
	//var res *pb.CategoryItemRes
	category := model.Category{}
	// TODO 业务逻辑判断
	category.Name = req.Name
	category.Level = req.Level
	if category.Level > 1 {
		category.ParentCategoryID = req.ParentCategoryId
	}
	r := internal.DB.Save(&category)
	fmt.Println(r.Error.Error())
	res := ConventCategoryModel2Pb(category)
	return res, nil
}

func (p ProductServer) GetAllCategoryList(ctx context.Context, empty *emptypb.Empty) (*pb.CategoriesRes, error) {
	var categoryList []model.Category
	internal.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categoryList)
	var res pb.CategoriesRes
	var items []*pb.CategoryItemRes
	for _, c := range categoryList {
		items = append(items, ConventCategoryModel2Pb(c))
	}
	b, err := json.Marshal(items)
	if err != nil {
		return nil, errors.New(custom_error.MarshalCategoryFailed)
	}
	res.InfoResList = items
	res.CategoryJsonFormat = string(b)
	return &res, nil
}

func (p ProductServer) GetSubCategory(ctx context.Context, req *pb.CategoriesReq) (*pb.SubCategoriesRes, error) {
	var category model.Category
	var subItemList []*pb.CategoryItemRes
	var res pb.SubCategoriesRes
	r := internal.DB.First(&category, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	pre := "SubCategory"
	if category.Level == 1 {
		pre = "SubCategory.SubCategory"
	}
	var subCategoryList []model.Category
	internal.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(pre).Find(&subCategoryList)
	for _, c := range subCategoryList {
		subItemList = append(subItemList, ConventCategoryModel2Pb(c))
	}
	b, err := json.Marshal(subItemList)
	if err != nil {
		return nil, errors.New(custom_error.MarshalCategoryFailed)
	}
	res.SubCategoryList = subItemList
	res.CategoryJsonFormat = string(b)
	return &res, nil
}

func (p ProductServer) DeleteCategory(ctx context.Context, req *pb.CategoryDelReq) (*emptypb.Empty, error) {
	internal.DB.Delete(&model.Category{}, req.Id)
	//ToDo 逻辑判断 1、 如果你删除的是1 级分类，下面的2级3级也要删掉, 级联删除
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateCategory(ctx context.Context, req *pb.CategoryItemReq) (*emptypb.Empty, error) {
	var category model.Category
	r := internal.DB.Find(&category, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategoryId > 0 {
		category.ParentCategoryID = req.ParentCategoryId
	}
	if req.Level > 0 {
		category.Level = req.Level
	}
	internal.DB.Save(&category)
	return &emptypb.Empty{}, nil
}

func ConventCategoryModel2Pb(c model.Category) *pb.CategoryItemRes {
	item := &pb.CategoryItemRes{
		Id:    c.ID,
		Name:  c.Name,
		Level: c.Level,
	}
	if c.Level > 1 {
		item.ParentCategoryId = c.ParentCategoryID
	}
	return item
}
