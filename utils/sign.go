package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

// encodeByteToHex 与hex标准库不同的是，字母大写。
func encodeByteToHex(src []byte) string {
	const HexTable = "0123456789ABCDEF"
	dst := make([]byte, hex.EncodedLen(len(src)))
	j := 0
	for _, v := range src {
		dst[j] = HexTable[v>>4]
		dst[j+1] = HexTable[v&0x0f]
		j += 2
	}
	return string(dst)
}

// TaobaoSign 按照淘宝要求进行签名，算法是hmac_md5
// 参考文档： https://open.taobao.com/doc.htm?docId=101617&docType=1
func TaobaoSign(data []byte, secret []byte) string {
	h := hmac.New(md5.New, secret)
	h.Write(data)
	return encodeByteToHex(h.Sum(nil))
}
