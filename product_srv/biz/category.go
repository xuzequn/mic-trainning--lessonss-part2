package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/proto/pb"
)

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
