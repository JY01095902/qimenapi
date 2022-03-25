package request

import "errors"

var (
	ErrCallQiMenAPIFailed = errors.New("call Qi Men API failed")
	ErrQiMenAPIBizError   = errors.New("Qi Men API biz error")
)
