请继续编写platform/platform.go文件，实现功能：
平台接口封装文件，用于请求平台接口并解析返回结果。
以下是一个简单的platform.go文件示例，用于封装请求平台接口的逻辑：

go
Copy code
package platform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Platform 接口表示一个平台的信息，包括平台类型、接口地址和请求参数等
type Platform interface {
	Type() string
	Balance() (float64, error)
}

// platform 结构体表示一个平台的信息
type platform struct {
	typ     string                 // 平台类型
	address string                 // 接口地址
	params  map[string]interface{} // 请求参数
}

// Type 方法返回该平台的类型
func (p *platform) Type() string {
	return p.typ
}

// Balance 方法请求平台接口并解析余额信息
func (p *platform) Balance() (float64, error) {
	// 构造请求参数
	data := url.Values{}
	for k, v := range p.params {
		data.Set(k, fmt.Sprint(v))
	}
	reqBody := strings.NewReader(data.Encode())

	// 发起请求
	req, err := http.NewRequest("POST", p.address, reqBody)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// 解析返回结果
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	balance := result["balance"].(float64)
	return balance, nil
}

// NewPlatform 函数根据平台类型和配置信息创建一个平台实例
func NewPlatform(typ string, config map[string]interface{}) Platform {
	p := &platform{
		typ:     typ,
		address: config["address"].(string),
		params:  config["params"].(map[string]interface{}),
	}
	return p
}
该文件定义了一个Platform接口和一个platform结构体，用于表示一个平台的信息。platform结构体实现了Platform接口的方法，并提供了请求平台接口并解析余额信息的逻辑。具体来说，Balance方法构造了请求参数并发起了POST请求，然后解析返回结果并返回余额信息。NewPlatform函数根据平台类型和配置信息创建了一个平台实例，其中平台类型、接口地址和请求参数等信息都来自配置文件。