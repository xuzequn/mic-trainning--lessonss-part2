package model

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"mic-trainning-lessons-part2/internal"
	"strconv"
	"time"
)

type BaseModel struct {
	ID        int32 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(32);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID"`
}

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(32);not null"`
	Logo string `gorm:"type:varchar(256);not null;default:''"`
}

type Advertise struct {
	BaseModel
	Index int32  `gorm:"type:int;not null;default:1"`
	Image string `gorm:"type:varchar(256);not null"`
	Url   string `gorm:"type:varchar(256);not null"`
	Sort  int32  `gorm:"type:int;not null;default:1"`
}

type Product struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category

	BrandID int32 `gorm:"type:int;not null"`
	Brand   Brand

	Selling    bool   `gorm:"default:false"`
	IsShipFree bool   `gorm:"default:false"`
	IsPop      bool   `gorm:"default:false"`
	IsNew      bool   `gorm:"default:false"`
	KeyWord    string `gorm:"type:varchar(64);not null"`

	Name       string  `gorm:"type:varchar(64);not null"`
	SN         string  `gorm:"type:varchar(64);not null"`
	FavNum     int32   `gorm:"type:int;default:0"`
	SoldNum    int32   `gorm:"type:int;default:0"`
	Price      float32 `gorm:"not null"`
	RealPrice  float32 `gorm:"not null"`
	ShortDesc  string  `gorm:"type:varchar(256);not null"`
	Images     MyList  `gorm:"type:varchar(1024);not null"`
	DesImages  MyList  `gorm:"type:varchar(1024);not null"`
	CoverImage string  `gorm:"type:varchar(256);not null"`
}

type ProductCategoryBrand struct {
	BaseModel

	BrandID    int32 `gorm:"type:int;not null"`
	Brand      Brand
	CategoryID int32 `gorm:"type:int;not null"`
	Category   Category
}

type MyList []string

func (myList MyList) Value() (driver.Value, error) {
	return json.Marshal(myList)
}

func (myList MyList) Scan(v interface{}) error {
	return json.Unmarshal(v.([]byte), &myList)
}

func (p *Product) AfterCreate(tx *gorm.Tx) error {
	esProduct := ESProduct{
		ID:         0,
		BrandID:    0,
		CategoryID: 0,
		Selling:    false,
		ShipFree:   false,
		IsPop:      false,
		IsNew:      false,
		Name:       "",
		FavNum:     0,
		SoldNum:    0,
		Price:      0,
		RealPrice:  0,
		ShortDesc:  "",
	}
	_, err := internal.ESClient.Index().Index(GetIndex()).BodyJson(esProduct).Id(strconv.Itoa(int(p.ID))).Do(context.Background())
	if err != nil {
		panic(err)
	}
	return err
}

// 更新删除操作同步到es
