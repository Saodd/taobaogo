package taobaoapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/saodd/taobaogo/constants"
	"github.com/saodd/taobaogo/utils"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	AppKey            string
	AppSecret         []byte
	TaobaoApiEndpoint string
	HttpDo            func(req *http.Request) (*http.Response, error)

	// HandleError 是用来处理错误的函数。推荐使用 github.com/saodd/alog.CE
	HandleError func(context.Context, error, ...map[string]interface{})

	DebugFlag   bool
	DebugPrintf func(format string, v ...interface{})
}

// NewClient 提供常用的配置。如果需要更多特性，请自己实例化。
func NewClient(appKey string, appSecret string, debug bool) *Client {
	return &Client{
		AppKey:            appKey,
		AppSecret:         []byte(appSecret),
		TaobaoApiEndpoint: "https://eco.taobao.com/router/rest",
		HttpDo:            http.DefaultClient.Do,
		HandleError:       func(c context.Context, err error, v ...map[string]interface{}) { log.Println(err, v) },
		DebugFlag:         debug,
		DebugPrintf:       log.Printf,
	}
}

// Do 快捷方法，执行一个请求
func (client *Client) Do(ctx context.Context, body RequestParams, sp SystemParams, res interface{}) (err error) {
	// 0. Debug输出
	var u string
	var reqBody string
	var respBody []byte
	if client.DebugFlag {
		defer func() {
			client.DebugOutput(u, reqBody, respBody, err)
		}()
	}

	// 1. 构建请求
	u = client.BuildUrl(body, sp)
	reqBody = body.ToValues().Encode()
	req, _ := http.NewRequestWithContext(ctx, "POST", u, strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 2. 执行请求
	resp, err := client.HttpDo(req)
	if err != nil {
		client.HandleError(ctx, err, nil)
		return err
	}
	respBody, _ = ioutil.ReadAll(resp.Body)
	if err = resp.Body.Close(); err != nil {
		client.HandleError(ctx, err, nil)
	}

	// 3. 解析错误
	var errResp struct {
		ErrorResponse *SystemError `json:"error_response"`
	}
	if err = json.Unmarshal(respBody, &errResp); err != nil {
		return err
	}
	if errResp.ErrorResponse != nil {
		return errResp.ErrorResponse
	}

	// 4. 解析结果
	return json.Unmarshal(respBody, res)
}

// BuildUrl 对请求参数进行签名，并且仅仅将系统参数放在query中。
// 提醒：外部需要再将请求参数放入form中。
func (client *Client) BuildUrl(rp RequestParams, sp SystemParams) string {
	var now = time.Now().In(constants.CST).Format(constants.TaobaoDatetimeFormat)
	var spm = sp.ToSignMap()
	// 1. 准备公共参数
	var toSign = rp.ToSignMap()
	toSign["app_key"] = client.AppKey
	toSign["sign_method"] = "hmac"
	toSign["format"] = "json"
	toSign["v"] = "2.0"
	toSign["timestamp"] = now
	for k, v := range spm {
		toSign[k] = v
	}
	// 2. 排序
	var keys []string
	for key := range toSign {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var builder bytes.Buffer
	for _, key := range keys {
		builder.WriteString(key)
		builder.WriteString(toSign[key])
	}
	// 3. 签名（摘要算法）
	sign := utils.TaobaoSign(builder.Bytes(), client.AppSecret)

	// 4. 构建url
	u, _ := url.Parse(client.TaobaoApiEndpoint)
	q := u.Query()
	q.Set("app_key", client.AppKey)
	q.Set("sign_method", "hmac")
	q.Set("format", "json")
	q.Set("v", "2.0")
	q.Set("timestamp", now)
	for k, v := range spm {
		q.Set(k, v)
	}
	q.Set("sign", sign)
	u.RawQuery = q.Encode()
	return u.String()
}

func (client *Client) DebugOutput(u, reqBody string, respBody []byte, err error) {
	const template = `[TAOBAOAPI-DEBUG]
  ***Request:
    POST %s
    %s
  ***Response:
    %s
  ***Error:
    %s
  ***END***
`
	client.DebugPrintf(template, u, reqBody, string(respBody), err)
}

type RequestParams interface {
	// ToSignMap 仅提供需要签名的参数（[]byte不签名）
	ToSignMap() map[string]string

	// ToValues 提供所有私有请求参数（将放入 body form 里）
	ToValues() url.Values

	// Valid 检查参数是否合法
	Valid() error
}

// SystemParams 是淘宝接口的通用参数中的需要填写的部分。
type SystemParams struct {
	Method  string  `json:"method"`
	Session *string `json:"session"`
}

func (sp SystemParams) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["method"] = sp.Method
	if sp.Session != nil {
		m["session"] = *sp.Session
	}
	return m
}

// SystemParamsAll 是淘宝接口的通用参数，也是要求必须放在query中的内容，这里仅做文档，没有使用意义。
type SystemParamsAll struct {
	SystemParams
	AppKey     string `json:"app_key"`
	SignMethod string `json:"sign_method"`
	Format     string `json:"format"`
	V          string `json:"v"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}

// SystemError 封装了淘宝返回的错误信息。
type SystemError struct {
	RequestId string `json:"request_id"`
	Msg       string `json:"msg"`
	Code      int    `json:"code"`
	SubMsg    string `json:"sub_msg"`
	SubCode   string `json:"sub_code"`
}

func (e *SystemError) Error() string {
	return fmt.Sprintf("TaobaoAPIError: %s | %d | %s | %s | %s", e.RequestId, e.Code, e.Msg, e.SubCode, e.SubMsg)
}

// RequestParamInvalidError 用于 RequestParams.Valid() 的返回错误信息
type RequestParamInvalidError struct {
	FieldName string
	Detail    string
}

func (r *RequestParamInvalidError) Error() string {
	return fmt.Sprintf("[%s]: %s", r.FieldName, r.Detail)
}
