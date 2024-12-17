package request

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	applicationKey = "c5588654-b0eb-4241-8d3b-25b853f69bcb"
	hmacSecretKey  = "90cfa800-d968-4a04-b233-905d1d31db1f"
	apiURL         = "https://cloud.myscript.com/api/v4.0/iink/contextlessgesture"
)

// generateHMACSignature 生成 HMAC 签名
func generateHMACSignature(secretKey, payload string) string {
	h := hmac.New(sha512.New, []byte(secretKey))
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}

// processHandwriting 发送请求并接收响应
func processHandwriting(payload map[string]interface{}) (string, error) {
	// 将 payload 转换为 JSON 字符串
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	// 生成 HMAC 签名
	hmacSignature := generateHMACSignature(hmacSecretKey, string(payloadJSON))

	// 创建请求
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadJSON))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("applicationKey", applicationKey)
	req.Header.Set("hmac", hmacSignature)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应数据
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		return string(responseBody), nil
	} else {
		return "", fmt.Errorf("unexpected response status: %s, body: %s", resp.Status, string(responseBody))
	}
}
