package taobaomodels

import (
	"errors"
	"github.com/saodd/taobaogo/constants"
	"time"
)

// TaobaoDatetime 负责解决淘宝时间格式的序列化和反序列化
type TaobaoDatetime struct {
	time.Time
}

func (t TaobaoDatetime) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	b := make([]byte, 0, len(constants.TaobaoDatetimeFormat)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, constants.TaobaoDatetimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *TaobaoDatetime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t.Time, err = time.ParseInLocation(`"`+constants.TaobaoDatetimeFormat+`"`, string(data), constants.CST)
	return err
}
