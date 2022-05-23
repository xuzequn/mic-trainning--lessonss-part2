package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/proto/pb"
)

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
