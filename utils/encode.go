package utils

import (
	"net/url"
)

func QueryUnescape(s *string) error {
	newStr, err := url.QueryUnescape(*s)
	if err != nil {
		return err
	}
	*s = newStr
	return nil
}
