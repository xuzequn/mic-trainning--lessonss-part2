package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/proto/pb"
)

type ProductServer struct {
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductConditionReq) (*pb.ProductsRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) BatchGetProduct(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductsRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateProduct(ctx context.Context, item *pb.CreateProductItem) (*pb.ProductItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteProduct(ctx context.Context, item *pb.ProductDelItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateProduct(ctx context.Context, item *pb.CreateProductItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) GetProductDetail(ctx context.Context, req *pb.ProductItemReq) (*pb.ProductItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) GetAllCategoryList(ctx context.Context, empty *emptypb.Empty) (*pb.CategoriesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) GetSubCategory(ctx context.Context, req *pb.CategoriesReq) (*pb.SubCategoriesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateCategory(ctx context.Context, req *pb.CategoryItemReq) (*pb.CategoryItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteCategory(ctx context.Context, req *pb.CategoryDelReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateCategory(ctx context.Context, req *pb.CategoryItemReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) BandList(ctx context.Context, req *pb.BrandPagingReq) (*pb.BrandRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateBrand(ctx context.Context, req *pb.BrandItemReq) (*pb.BrandItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteBrand(ctx context.Context, req *pb.BrandItemReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateBrand(ctx context.Context, req *pb.BrandItemReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) AdvertiseList(ctx context.Context, empty *emptypb.Empty) (*pb.AdvertisesRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*pb.AdvertiseItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CategoryBrandList(ctx context.Context, req *pb.PagingReq) (*pb.CategoryBrandListRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) GetCategoryBrandList(ctx context.Context, req *pb.CategoryItemReq) (*pb.CategoryItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*pb.CategoryBrandRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateCategoryBrand(ctx context.Context, req *pb.CategoryBrandReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
