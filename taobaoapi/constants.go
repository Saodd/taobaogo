package taobaoapi

import "time"

const (
	TaobaoDatetimeFormat = "2006-01-02 15:04:05"
)

var (
	CST, _ = time.LoadLocation("Asia/Shanghai")
)
