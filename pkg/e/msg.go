package e

/*
* @Author:hanyajun
* @Date:2019/5/22 23:07
* @Name:e
* @Function: 错误信息定义
 */
var MsgFlags = map[int]string{
	SUCCESS:        "成功",
	ERROR:          "失败",
	INVALID_PARAMS: "参数错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
