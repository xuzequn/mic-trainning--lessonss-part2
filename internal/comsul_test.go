package internal

import (
	"fmt"
	"testing"
)

func TestReg(t *testing.T) {
	err := Reg(AppConf.ProductWebConfig.Host, AppConf.ProductWebConfig.SrvName,
		AppConf.ProductWebConfig.SrvName, AppConf.ProductWebConfig.Port,
		AppConf.ProductWebConfig.Tags)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("注册成功")
	}
}

func TestGetServiceList(t *testing.T) {
	GetServiceList()
}

func TestFilterService(t *testing.T) {
	FilterService()
}
