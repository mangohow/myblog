package response

import (
	"blog_web/utils"
	"fmt"
)

func CheckError(err error, msg string, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			utils.Logger().Warning("%s:%v", fmt.Sprintf(msg, args...), err)
		} else {
			utils.Logger().Warning("%s:%v", msg, err)
		}
		return true
	}

	return false
}
