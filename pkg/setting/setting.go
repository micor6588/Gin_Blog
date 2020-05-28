/***************************************************
 ** @Desc : 用于存储配置文件
 ** @Time : 2020/5/28 13:56
 ** @Author : JiangFeng
 ** @File : error_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2020/5/28 16:56
 ** @Software: VScode
****************************************************/
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg          *ini.File
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSecret    string
)

// 初始化配置
func init() {
	var err error
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {

}
