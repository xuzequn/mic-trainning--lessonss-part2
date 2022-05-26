package biz

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/model"
	"mic-trainning-lessons-part2/proto/pb"
)

func (p ProductServer) AdvertiseList(ctx context.Context, empty *emptypb.Empty) (*pb.AdvertisesRes, error) {
	var adList []model.Advertise
	var adItemList []*pb.AdvertiseItemRes
	var advertiseRes pb.AdvertisesRes
	r := internal.DB.Find(&adList)
	for _, item := range adList {
		adItemList = append(adItemList, ConverAdModel2Pb(item))
	}
	advertiseRes.Total = int32(r.RowsAffected)
	advertiseRes.ItemList = adItemList
	return &advertiseRes, nil
}

func (p ProductServer) CreateAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*pb.AdvertiseItemRes, error) {
	var ad model.Advertise
	ad.Index = req.Index
	ad.Image = req.Images
	ad.Url = req.Url
	internal.DB.Save(&ad)
	return ConverAdModel2Pb(ad), nil
}

func (p ProductServer) DeleteAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateAdvertise(ctx context.Context, req *pb.AdvertiseReq) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func ConverAdModel2Pb(item model.Advertise) *pb.AdvertiseItemRes {
	ad := &pb.AdvertiseItemRes{
		Index:  item.Index,
		Images: item.Image,
		Url:    item.Url,
	}
	if item.ID > 0 {
		ad.Id = item.ID
	}
	return ad
}
