package internal

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"mic-trainning-lessons-part2/model"
)

var ESClient *elastic.Client

type ESConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

func InitES() {
	host := fmt.Sprintf("%s:%d", AppConf.EsConfig.Host, AppConf.EsConfig.Port)
	var err error
	ESClient, err = elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	ok, err := ESClient.IndexExists(model.GetIndex()).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !ok {
		_, err := ESClient.CreateIndex(model.GetIndex()).
			BodyString(model.GetMapping()).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
