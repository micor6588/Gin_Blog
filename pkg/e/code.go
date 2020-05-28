/***************************************************
 ** @Desc : 用于对错误码进行定义
 ** @Time : 2020/5/28 13:56
 ** @Author : JiangFeng
 ** @File : error_gateway
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2020/5/28 16:56
 ** @Software: VScode
****************************************************/

package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)
