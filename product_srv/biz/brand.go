package biz

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/custom_error"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/model"
	"mic-trainning-lessons-part2/proto/pb"
)

func (p ProductServer) CreateBrand(ctx context.Context, req *pb.BrandItemReq) (*pb.BrandItemRes, error) {
	var brand model.Brand
	r := internal.DB.Find("name=? and log=?", req.Name, req.Logo)
	if r.RowsAffected > 0 {
		return nil, errors.New(custom_error.BrandAlreadyExits)
	}
	brand.Name = req.Name
	brand.Logo = req.Logo
	internal.DB.Save(&brand)
	return ConverBrandModel2Pb(brand), nil

}

func (p ProductServer) BandList(ctx context.Context, req *pb.BrandPagingReq) (*pb.BrandRes, error) {
	var brandList []model.Brand
	var brands []*pb.BrandItemRes
	var brandRes pb.BrandRes

	//直接查询
	//r := internal.DB.Find(&brandList)
	//fmt.Println(r.RowsAffected)
	//for _, item := range brandList {
	//	brands = append(brands, ConverBrandModel2Pb(item))
	//}
	//brandRes.ItemList = brands
	//brandRes.Total = int32(r.RowsAffected)

	//第一种分页
	//var count int64
	//skip := (req.PageNo - 1) * req.PageSize
	//r := internal.DB.Model(&model.Brand{}).Count(&count).Offset(int(skip)).Limit(int(req.PageSize)).Find(&brandList)
	//if r.RowsAffected < 1 {
	//	// TODO: 根据业务需求进一步判断
	//}
	//brandRes.Total = int32(count)
	//for _, item := range brandList {
	//	brands = append(brands, ConverBrandModel2Pb(item))
	//}
	//brandRes.ItemList = brands
	//return &brandRes, nil

	//第二种分页
	_ = internal.DB.Scopes(internal.MyPaging(int(req.PageNo), int(req.PageSize))).Find(&brandList)
	for _, item := range brandList {
		brands = append(brands, ConverBrandModel2Pb(item))
	}
	brandRes.ItemList = brands
	var count int64
	internal.DB.Model(&model.Brand{}).Count(&count)
	brandRes.Total = int32(count)
	return &brandRes, nil
}

func (p ProductServer) DeleteBrand(ctx context.Context, req *pb.BrandItemReq) (*emptypb.Empty, error) {
	r := internal.DB.Delete(&model.Brand{}, req.Id)
	if r.Error != nil {
		return nil, errors.New(custom_error.DelBrandFail)
	}
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateBrand(ctx context.Context, req *pb.BrandItemReq) (*emptypb.Empty, error) {
	var brand model.Brand
	r := internal.DB.First(&brand, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}
	if req.Name != "" {
		brand.Name = req.Name
	}
	if req.Logo != "" {
		brand.Logo = req.Logo
	}
	internal.DB.Save(&brand)
	return &emptypb.Empty{}, nil
}

func ConverBrandModel2Pb(item model.Brand) *pb.BrandItemRes {
	brand := &pb.BrandItemRes{
		Name: item.Name,
		Logo: item.Logo,
	}

	if item.ID > 0 {
		brand.Id = item.ID
	}
	return brand
}
