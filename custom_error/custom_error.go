package custom_error

const (
	BrandAlreadyExits = "品牌已存在"
	DelBrandFail      = "删除失败"
	BrandNotExits     = "品牌不存在"

	ADNotExists = "广告不存在"

	CategoryNotExits              = "分类没找到"
	MarshalCategoryFailed         = "序列化分类失败"
	DelProductCategoryBrandFailed = "删除分类品牌表失败"
	ProductCategoryBrandNotFind   = "分类品牌表找不到记录"

	DelProductFailed = "删除产品失败"
	ProductNotExits  = "产品不存在"

	ParamError = "参数错误"
)
