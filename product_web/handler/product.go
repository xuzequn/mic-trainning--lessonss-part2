package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mic-trainning-lessons-part2/custom_error"
	"mic-trainning-lessons-part2/internal"
	"mic-trainning-lessons-part2/product_web/req"
	"mic-trainning-lessons-part2/proto/pb"
	"net/http"
	"strconv"
)

var productClient pb.ProductServiceClient

func init() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ProductSrvConfig.Host, internal.AppConf.ProductSrvConfig.Port)
	conn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{loadBalancingPolicy:"round_robbin"}`),
	)
	if err != nil {
		zap.S().Fatal(err)
		panic(err)
	}
	productClient = pb.NewProductServiceClient(conn)
}

func ProductListHandler(c *gin.Context) {

	var condition pb.ProductConditionReq
	//c.ShouldBind(&condition)
	//list?pageNo=1&pageSize=2

	minPriceStr := c.DefaultQuery("minPrice", "0")
	minPrice, err := strconv.Atoi(minPriceStr)
	if err != nil {
		zap.S().Error("minPrice error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ParamError,
		})
	}
	maxPriceStr := c.DefaultQuery("maxPrice", "0")
	maxPrice, err := strconv.Atoi(maxPriceStr)
	if err != nil {
		zap.S().Error("maxPrice error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ADNotExists,
		})
	}
	condition.MinPrice = int32(minPrice)
	condition.MaxPrice = int32(maxPrice)

	categoryIdStr := c.DefaultQuery("categoryId", "0")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		zap.S().Error("categoryId error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ADNotExists,
		})
	}
	condition.CategoryId = int32(categoryId)

	brandIdStr := c.DefaultQuery("brandId", "0")
	brandId, err := strconv.Atoi(brandIdStr)
	if err != nil {
		zap.S().Error("brandId error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ADNotExists,
		})
	}
	condition.BrandId = int32(brandId)

	isPop := c.DefaultQuery("isPop", "0")
	if isPop == "1" {
		condition.IsPop = true
	}

	isNew := c.DefaultQuery("isNew", "0")
	if isNew == "1" {
		condition.IsPop = true
	}

	pageNoStr := c.DefaultQuery("pageNo", "0")
	pageNo, err := strconv.Atoi(pageNoStr)
	if err != nil {
		zap.S().Error("pageNo error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ADNotExists,
		})
	}
	condition.PageNo = int32(pageNo)

	pageSizeStr := c.DefaultQuery("pageSize", "0")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		zap.S().Error("pageSize error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ADNotExists,
		})
	}
	condition.PageSize = int32(pageSize)

	keyWord := c.DefaultQuery("keyWord", "0")
	if err != nil {
		zap.S().Error("keyWord error")
		c.JSON(http.StatusOK, gin.H{
			"msg": custom_error.ADNotExists,
		})
	}
	condition.KeyWord = keyWord
	r, err := productClient.ProductList(context.Background(), &condition)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "产品列表查询失败",
			// 默认值
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "",
		"total": r.Total,
		"data":  r.ItemList,
	})
}

func AddHandler(c *gin.Context) {
	var productReq req.ProductReq
	err := c.ShouldBind(&productReq)
	if err != nil {
		zap.S().Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "参数解析错误",
		})
		return
	}
	r := ConvertProductReq2Pb(productReq)
	res, err := productClient.CreateProduct(context.Background(), r)
	if err != nil {
		zap.S().Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "添加产品失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "",
		"data": res,
	})
}

func DetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "参数错误",
		})
		return
	}
	res, err := productClient.GetProductDetail(context.Background(), &pb.ProductItemReq{
		Id: int32(id),
	})
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "获取详情失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "",
		"data": res,
	})
}

func DelHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "参数错误",
		})
		return
	}
	_, err = productClient.DeleteProduct(context.Background(), &pb.ProductDelItem{Id: int32(id)})
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "商品删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
	})
}

func UpdateHandler(c *gin.Context) {
	var productReq req.ProductReq
	err := c.ShouldBind(&productReq)
	if err != nil {
		zap.S().Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "参数解析错误",
		})
		return
	}
	r := ConvertProductReq2Pb(productReq)
	_, err = productClient.UpdateProduct(context.Background(), r)
	if err != nil {
		zap.S().Error(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "更新产品失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
	})

}

func ConvertProductReq2Pb(productReq req.ProductReq) *pb.CreateProductItem {
	item := pb.CreateProductItem{
		Name: productReq.Name,
		Sn:   productReq.SN,
		//Stocks:      productReq.Stocks,
		Price:       productReq.Price,
		RealPrice:   productReq.RealPrice,
		ShortDesc:   productReq.ShortDesc,
		ProductDesc: productReq.Desc,
		Images:      productReq.Images,
		DescImages:  productReq.DescImages,
		CoverImage:  productReq.CoverImage,
		IsNew:       productReq.IsNew,
		IsPop:       productReq.IsPop,
		Selling:     productReq.Selling,
		BrandId:     productReq.BrandId,
		FavNUm:      productReq.FavNum,
		SoldNum:     productReq.SoldNum,
		CategoryId:  productReq.CategoryId,
	}
	if productReq.Id > 0 {
		item.Id = productReq.Id
	}
	return &item
}
