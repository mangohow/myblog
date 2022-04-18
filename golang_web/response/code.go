package response

/*
* @Author: mgh
* @Date: 2022/2/24 18:45
* @Desc:
 */

const (
	QueryFailed uint32 = iota
	QuerySuccess

	OperateFailed = iota + 98
	OperateSuccess

	LoginFailed = iota + 196
	LoginSuccess

	Unauthorized = iota + 294

	DeleteFailed = iota + 393
	DeleteSuccess
)

var MessageForCode = map[uint32]string{
	QuerySuccess:   "查询成功",
	QueryFailed:    "查询失败",
	OperateFailed:  "操作失败",
	OperateSuccess: "操作成功",
	LoginFailed:    "登录失败",
	LoginSuccess:   "登录成功",
	Unauthorized:   "未获得授权",
	DeleteFailed:   "删除失败",
	DeleteSuccess:  "删除成功",
}

