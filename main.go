package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

func main() {
	client := &http.Client{}
	key, err := GetSignKey(client)
	if key != "" {
		SignIn(client, key)
	} else{
		fmt.Println("错误:", err)
        os.Exit(1)
    }
}

func extractSign(input string) (string, error) {
	pattern := `var sign = "(.*?)";`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)
    // 判断是否未登录

	if len(matches) > 1 {
		return matches[1], nil
	}else if strings.Contains(input, "请登录") {
		return "", fmt.Errorf("未登录！请更新Cookie后重新运行！")
    } else {
		return "", fmt.Errorf("sign not found")
	}
}

// 获取Key
func GetSignKey(client *http.Client) (string, error) {
	urlStr := "https://www.hifini.com/"
	cookie := os.Getenv("COOKIE")
	if cookie == "" {
		fmt.Println("COOKIE不存在，请检查是否添加")
		return "", nil
	}
	//提交请求
	formData := url.Values{}

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", cookie)
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//处理返回结果
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
	sign, err := extractSign(string(buf))
	if err != nil {
		fmt.Println("获取Sign失败:", err)
		return sign, err
	} else {
		fmt.Println("获取Sign成功:", sign)
		return sign, nil
	}
}

// SignIn 签到
func SignIn(client *http.Client, key string) bool {
	//生成要访问的url
	urlStr := "https://www.hifini.com/sg_sign.htm"
	cookie := os.Getenv("COOKIE")
	if cookie == "" {
		fmt.Println("COOKIE不存在，请检查是否添加")
		return false
	}
	if key == "" {
		fmt.Println("KEY获取失败，请检查Cookie是否过期")
		return false
	}

	//提交请求
	formData := url.Values{}
	formData.Set("sign", key)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(formData.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Cookie", cookie)
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//处理返回结果
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf, _ := ioutil.ReadAll(response.Body)
    fmt.Println("签到结果：")
	fmt.Println(string(buf))
	hasDing := os.Getenv("DINGDING_WEBHOOK")
	if hasDing != "" {
		dingding(string(buf))
	} else {
		fmt.Println("DINGDING_WEBHOOK 环境变量未定义，跳过通知步骤")
	}
	return strings.Contains(string(buf), "成功")
}

func dingding(result string) {
	// 构造要发送的消息
	message := struct {
		MsgType string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}{
		MsgType: "text",
		Text: struct {
			Content string `json:"content"`
		}{
			Content: "HiFiNi：\n" + result,
		},
	}

	// 将消息转换为JSON格式
	messageJson, _ := json.Marshal(message)
	DINGDING_WEBHOOK := os.Getenv("DINGDING_WEBHOOK")
	// 发送HTTP POST请求
	resp, err := http.Post(DINGDING_WEBHOOK,
		"application/json", bytes.NewBuffer(messageJson))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
