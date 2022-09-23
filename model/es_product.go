package model

type ESProduct struct {
	ID         int32   `json:"id"`
	BrandID    int32   `json:"brand_id"`
	CategoryID int32   `json:"category_id"`
	Selling    bool    `json:"selling"`
	ShipFree   bool    `json:"ship_free"`
	IsPop      bool    `json:"is_pop"`
	IsNew      bool    `json:"is_new"`
	Name       string  `json:"name"`
	FavNum     int32   `json:"fav_num"`
	SoldNum    int32   `json:"sold_num"`
	Price      float32 `json:"price"`
	RealPrice  float32 `json:"real_price"`
	ShortDesc  string  `json:"short_desc"`
}

func GetIndex() string {
	return "product"
}

func GetMapping() string {
	productMapping := `
    {
        "mappings" : {
                "id" : {
                    "type" : "integer"
                },
                "brand_id" : {
                    "type" : "integer"
                },
                "category_id" : {
                    "type" : "integer"
                },
                "selling" : {
                    "type" : "boolean"
                },
                "ship_free" : {
                    "type" : "boolean"
                },
                "is_new" : {
                    "type" : "boolean"
                },
                "name" : {
                    "type" : "text",
                    "analyzer" : "ik_max_word"
                },
                "fav_num" : {
                    "type" : "integer"
                },                
                "sold_num" : {
                    "type" : "integer"
                },
                "price" : {
                    "type" : "float"
                }, 
                "price" : {
                    "type" : "float"
                }, 
                "fav_name" : {
                    "type" : "integer"
                }, 
                "fav_name" : {
                    "type" : "integer"
                },
                "short_desc" : {
                    "type" : "text"
                    "analyzer" : "ik_max_word"
                }
        }
    }
`
	return productMapping
}
