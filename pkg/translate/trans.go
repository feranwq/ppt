package translate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// 百度翻译开放平台信息
type BaiduInfo struct {
	AppID     string
	Salt      string
	SecretKey string
	From      string
	To        string
	Text      string
}

// 返回结果
type TransResult struct {
	From      string   `json:"from"`
	To        string   `json:"to"`
	Result    []Result `json:"trans_result"`
	ErrorCode string   `json:"error_code"`
	ErrorMsg  string   `json:"error_msg"`
}
type Result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

// 自动生盐
// 入口参数为盐的长度
func Salt(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成32位MD5
func Sign(bi *BaiduInfo) string {
	text := bi.AppID + bi.Text + bi.Salt + bi.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// 翻译 传入需要翻译的语句
func (bi *BaiduInfo) Translate() string {
	apiUrl := "http://api.fanyi.baidu.com"
	resource := "/api/trans/vip/translate"
	params := url.Values{}
	params.Add("from", bi.From)
	params.Add("to", bi.To)
	params.Add("appid", bi.AppID)
	params.Add("salt", bi.Salt)
	params.Add("sign", Sign(bi))
	params.Add("q", bi.Text)
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(params.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		log.Println("网络异常")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var ts TransResult
	err = json.Unmarshal(body, &ts)
	if err != nil {
		panic(err)
	}
	if ts.ErrorCode != "" {
		return ts.ErrorMsg
	} else {
		res := ""
		for _, v := range ts.Result {
			res = strings.Join([]string{res, v.Dst}, "\n")
		}
		return res
	}
}
