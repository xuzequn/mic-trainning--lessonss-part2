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

type ProductServer struct {
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductConditionReq) (*pb.ProductsRes, error) {
	iDb := internal.DB.Model(model.Product{})
	var productList []model.Product
	var itemList []*pb.ProductItemRes
	var res pb.ProductsRes

	if req.IsPop {
		iDb = iDb.Where("is_pop = ?", req.IsPop)
	}
	if req.IsNew {
		iDb = iDb.Where("is_new = ?", req.IsNew)
	}

	if req.BrandId > 0 {
		iDb = iDb.Where("brand = ?", req.BrandId)
	}
	if req.KeyWord != "" {
		iDb = iDb.Where("key_word like ?", "%"+req.KeyWord+"%")
	}
	if req.MinPrice > 0 {
		iDb = iDb.Where("min_price > ?", req.MinPrice)
	}
	if req.MaxPrice > 0 {
		iDb = iDb.Where("max_price > ?", req.MaxPrice)
	}
	if req.CategoryId > 0 {
		var category model.Category
		r := internal.DB.First(&category, req.CategoryId)
		if r.RowsAffected == 0 {
			return nil, errors.New(custom_error.CategoryNotExits)
		}
		var q string
		if category.Level == 1 {
			q = fmt.Sprintf("select id from category where parent_category_id in (select id from category Where parent_category_id=%d", req.CategoryId)
		} else if category.Level == 2 {
			q = fmt.Sprintf("select id from category Where parent_category_id=%d", req.CategoryId)
		} else if category.Level == 3 {
			q = fmt.Sprintf("select if from category where id = %d", req.CategoryId)
		}
		iDb = iDb.Where(fmt.Sprintf("category_id in %s", q))
	}
	var count int64
	iDb.Count(&count)
	fmt.Println(count)

	iDb.Joins("Category").Joins("Brand").Scopes(internal.MyPaging(int(req.PageNo), int(req.PageSize))).Find(&productList)
	for _, item := range productList {
		res := ConvertProductModel2Pb(item)
		itemList = append(itemList, res)
	}
	res.ItemList = itemList
	res.Total = int32(count)
	return &res, nil

}

func (p ProductServer) BatchGetProduct(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductsRes, error) {
	var productList []model.Product
	var res pb.ProductsRes
	r := internal.DB.Find(&productList, req.Ids)
	res.Total = int32(r.RowsAffected)
	for _, item := range productList {
		pro := ConvertProductModel2Pb(item)
		res.ItemList = append(res.ItemList, pro)
	}
	return &res, nil
}

func (p ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductItem) (*pb.ProductItemRes, error) {
	var category model.Category
	var brand model.Brand
	var res *pb.ProductItemRes
	// todo 业务逻辑判断 更复杂
	r := internal.DB.First(&category, req.CategoryId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	r = internal.DB.First(&brand, req.BrandId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}
	product := model.Product{}
	item := ConvertReq2Model(product, req, category, brand)
	internal.DB.Save(&item)
	res = ConvertProductModel2Pb(item)
	return res, nil
}

func (p ProductServer) DeleteProduct(ctx context.Context, req *pb.ProductDelItem) (*emptypb.Empty, error) {
	r := internal.DB.Delete(&model.Product{}, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.DelProductFailed)
	}
	return &emptypb.Empty{}, nil
}

func (p ProductServer) UpdateProduct(ctx context.Context, req *pb.CreateProductItem) (*emptypb.Empty, error) {
	//TODO 业务逻辑判断
	var product model.Product
	var c model.Category
	var b model.Brand
	r := internal.DB.First(&product, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.ProductNotExits)
	}
	r = internal.DB.First(&c, req.CategoryId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.CategoryNotExits)
	}
	r = internal.DB.First(&b, req.BrandId)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.BrandNotExits)
	}
	pro := ConvertReq2Model(product, req, c, b)
	internal.DB.Updates(&pro)
	return &emptypb.Empty{}, nil
}

func (p ProductServer) GetProductDetail(ctx context.Context, req *pb.ProductItemReq) (*pb.ProductItemRes, error) {
	var pro model.Product
	var res *pb.ProductItemRes
	r := internal.DB.First(&pro, req.Id)
	if r.RowsAffected < 1 {
		return nil, errors.New(custom_error.ProductNotExits)
	}
	res = ConvertProductModel2Pb(pro)
	return res, nil
}

func ConvertReq2Model(p model.Product, req *pb.CreateProductItem, category model.Category, brand model.Brand) model.Product {
	if req.CategoryId > 0 {
		p.CategoryID = req.CategoryId
		p.Category = category
	}
	if req.Id > 0 {
		p.ID = req.Id
	}

	if req.BrandId > 0 {
		p.BrandID = req.BrandId
		p.Brand = brand
	}

	if req.Selling {
		p.Selling = req.Selling
	}
	if req.IsShipFree {
		p.IsShipFree = req.IsShipFree
	}
	if req.IsPop {
		p.IsPop = req.IsPop
	}
	if req.IsNew {
		p.IsNew = req.IsNew
	}
	if req.Name != "" {
		p.Name = req.Name
	}

	if req.Sn != "" {
		p.SN = req.Sn
	}

	if req.FavNUm > 0 {
		p.FavNum = req.FavNUm
	}
	return p
}

func ConvertProductModel2Pb(pro model.Product) *pb.ProductItemRes {
	p := &pb.ProductItemRes{
		Id:          pro.ID,
		CategoryId:  pro.CategoryID,
		Name:        pro.Name,
		Sn:          pro.SN,
		SoldNum:     pro.SoldNum,
		FavNum:      pro.FavNum,
		Price:       pro.Price,
		RealPrice:   pro.RealPrice,
		ShortDesc:   pro.ShortDesc,
		Images:      pro.Images,
		DescImages:  pro.DesImages,
		CoverImages: pro.CoverImage,
		IsNew:       pro.IsNew,
		IsPop:       pro.IsPop,
		Selling:     pro.Selling,
		Category: &pb.CategoryItemRes{
			Id:   pro.Category.ID,
			Name: pro.Category.Name,
		},
		Brand: &pb.BrandItemRes{Id: pro.Brand.ID,
			Name: pro.Brand.Name, Logo: pro.Brand.Logo},
	}
	return p
}
