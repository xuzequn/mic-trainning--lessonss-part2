package internal

import (
	"fmt"
	"testing"
)

func TestReg(t *testing.T) {
	err := Reg(AppConf.AccountWebConfig.Host, AppConf.AccountWebConfig.SrvName,
		AppConf.AccountWebConfig.SrvName, AppConf.AccountWebConfig.Port,
		AppConf.AccountWebConfig.Tags)
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
