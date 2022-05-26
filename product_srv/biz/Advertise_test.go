package biz

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"mic-trainning-lessons-part2/model"
	"mic-trainning-lessons-part2/proto/pb"
	"testing"
)

//func init() {
//
//	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, 56531)
//	//grpc.Dial(addr, grpc.WithInsecure())
//	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		panic(err)
//	}
//	client = pb.NewProductServiceClient(conn)
//}

func TestProductServer_CreateAdvertise(t *testing.T) {
	advertiseList := [10]model.Advertise{{Index: 1, Image: "image-1", Url: "url-1"}, {Index: 2, Image: "image-2", Url: "url-2"}}
	for i := 3; i < 10; i++ {
		advertiseList[i].Index = int32(i)
		advertiseList[i].Image = fmt.Sprintf("image-%d", i)
		advertiseList[i].Url = fmt.Sprintf("url-%d", i)
	}
	for _, advertise := range advertiseList {
		res, err := client.CreateAdvertise(context.Background(), &pb.AdvertiseReq{
			Index:  advertise.Index,
			Images: advertise.Image,
			Url:    advertise.Url,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res)
	}

}

func TestProductServer_AdvertiseList(t *testing.T) {
	res, err := client.AdvertiseList(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_DeleteAdvertise(t *testing.T) {
	res, err := client.DeleteAdvertise(context.Background(), &pb.AdvertiseReq{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestProductServer_UpdateAdvertise(t *testing.T) {
	res, err := client.UpdateAdvertise(context.Background(), &pb.AdvertiseReq{
		Id:     3,
		Index:  10,
		Images: "image-10",
		Url:    "url-10",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
