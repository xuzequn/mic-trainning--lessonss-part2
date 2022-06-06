package biz

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/custom_error"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/model"
	"mic-trainning-lessons-part2/proto/pb"
)

func (p ProductServer) CategoryBrandList(ctx context.Context, req *pb.PagingReq) (*pb.CategoryBrandListRes, error) {
	// TODO 各种逻辑判断
	var items []model.ProductCategoryBrand
	var resList []*pb.CategoryBrandRes
	var count int64
	internal.DB.Model(model.ProductCategoryBrand{}).Count(&count)
	var res pb.CategoryBrandListRes

	internal.DB.Preload("Category").Preload("Brand").Scopes(internal.MyPaging(int(req.PageNo), int(req.PageSize))).Find(&items)
	for _, item := range items {
		pcb := ConventProductCategoryBrand2Pb(item)
		resList = append(resList, pcb)
	}
	res.Total = int32(count)
	res.ItemList = resList
	return &res, nil
}

func (p ProductServer) CreateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*pb.CategoryBrandRes, error) {
	var res pb.CategoryBrandRes
	var item model.ProductCategoryBrand
	var category model.Category
	var brand model.Brand
	// 分类判断
	r := internal.DB.First(&category, req.Category.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	// 品牌判断
	r = internal.DB.First(&brand, req.Brand.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}

	// 是否已经存在关系
	r = internal.DB.Where("category_id = ? and brand_id = ? ", req.Category.Id, req.Brand.Id).First(&item)
	if r.RowsAffected == 1 {
		fmt.Println("品牌分类关系已存在")
		res.Id = item.ID
		return &res, nil
	}
	item.CategoryID = req.Category.Id
	item.BrandID = req.Brand.Id
	internal.DB.Save(&item)
	res.Id = item.ID
	return &res, nil

}

func (p ProductServer) GetCategoryBrandList(ctx context.Context, req *pb.CategoryItemReq) (*pb.BrandRes, error) {
	var res pb.BrandRes
	var category model.Category
	var itemList []model.ProductCategoryBrand
	var itemListRes []*pb.BrandItemRes

	r := internal.DB.First(&category, req.Id)
	if r.RowsAffected == 0 {
		return nil, errors.New(custom_error.ProductCategoryBrandNotFind)
	}
	r = internal.DB.Preload("Category").Where(&model.ProductCategoryBrand{CategoryID: req.ParentCategoryId}).Find(&itemList)
	if r.RowsAffected > 0 {
		res.Total = int32(r.RowsAffected)
	}
	for _, item := range itemList {
		itemListRes = append(itemListRes, &pb.BrandItemRes{
			Id:   item.BrandID,
			Name: item.Brand.Name,
			Logo: item.Brand.Logo,
		})
	}
	res.ItemList = itemListRes

	return &res, nil
}

func (p ProductServer) DeleteCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	r := internal.DB.Delete(&model.ProductCategoryBrand{}, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.DelProductCategoryBrandFailed)
	}
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	var category model.Category
	var brand model.Brand
	var pcb model.ProductCategoryBrand
	//分类判断
	r := internal.DB.First(&category, req.Category.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	//品牌判断
	r = internal.DB.First(&brand, req.Brand.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}

	r = internal.DB.First(&pcb, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.ProductCategoryBrandNotFind)
	}
	pcb.CategoryID = req.Category.Id
	pcb.BrandID = req.Brand.Id
	internal.DB.Save(&pcb)

	return &emptypb.Empty{}, nil

}

func ConventProductCategoryBrand2Pb(pcb model.ProductCategoryBrand) *pb.CategoryBrandRes {
	return &pb.CategoryBrandRes{
		Id: pcb.ID,
		Brand: &pb.BrandItemRes{
			Id:   pcb.BrandID,
			Name: pcb.Brand.Name,
			Logo: pcb.Brand.Logo,
		},
		Category: &pb.CategoryItemRes{
			Id:               pcb.ID,
			Name:             pcb.Category.Name,
			ParentCategoryId: pcb.CategoryID,
			Level:            pcb.Category.Level,
		},
	}
}
