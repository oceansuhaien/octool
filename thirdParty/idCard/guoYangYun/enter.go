package guoYangYun

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/oceansuhaien/octool/validRule"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type Client struct {
	path    string
	AppCode string
}

// 返回请求
type response struct {
	Msg     string `json:"msg"`
	Code    string `json:"code"`
	Balance string `json:"balance"`
	Data    data   `json:"data,omitempty"`
}

// 返回data
type data struct {
	ResponseDesc string `json:"responseDesc"`
	TraceNo      string `json:"traceNo"`
	ResponseCode string `json:"responseCode"`
}

func New(appCode string) *Client {
	return &Client{
		path:    "https://gyidcard.market.alicloudapi.com/verify/idcard2",
		AppCode: appCode,
	}
}

func (a *Client) baseHook() error {
	if a.AppCode == "" {
		return ErrAppCode
	}
	return nil
}

// 验证身份证
// name 真实姓名
// 身份证号
func (a *Client) Valid(name, idCardNo string) (ok bool, balance int, err error) {
	// 校验身份证号码
	if ok := regexp.MustCompile(validRule.ValidIdCard).MatchString(idCardNo); !ok {
		return false, 0, ErrIdCardFormat
	}
	// 校验姓名
	if ok := regexp.MustCompile(validRule.ValidRealName).MatchString(name); !ok {
		return false, 0, ErrName
	}
	parseUrl, err := url.Parse(a.path)
	if err != nil {
		return false, 0, ErrParseUrl
	}
	params := url.Values{}
	params.Add("name", name)
	params.Add("idCardNo", idCardNo)
	path := fmt.Sprintf("%s?%s", parseUrl.String(), params.Encode())
	request, err := http.NewRequest(http.MethodPost, path, bytes.NewReader([]byte{}))
	if err != nil {
		return false, 0, ErrInitRequest
	}
	request.Header.Set("Authorization", fmt.Sprintf("APPCODE %s", a.AppCode))
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return false, 0, ErrClientDo
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return false, 0, ErrAuth
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, 0, ErrReadBody
	}
	var res response
	err = json.Unmarshal(all, &res)
	if err != nil {
		return false, 0, ErrUnmarshal
	}
	// 获取剩余次数
	balance, err = strconv.Atoi(res.Balance)
	if res.Code != "0" {
		return false, balance, nil
	}
	if res.Data.ResponseCode != "0" {
		return false, balance, ErrFail
	}
	return true, 0, nil
}
