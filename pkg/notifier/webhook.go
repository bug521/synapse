package notifier

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type WebhookConfig struct {
	URL     string
	Method  string // POST/PUT/GET
	Headers map[string]string
	Proxy   string // 可选
}

func SendWebhook(cfg WebhookConfig, body []byte) (string, error) {
	if cfg.URL == "" {
		return "", errors.New("Webhook URL 不能为空")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	if cfg.Proxy != "" {
		proxyURL, err := url.Parse(cfg.Proxy)
		if err != nil {
			return "", errors.New("代理地址格式错误")
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	method := cfg.Method
	if method == "" {
		method = "POST"
	}

	req, err := http.NewRequest(method, cfg.URL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	for k, v := range cfg.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return string(respBody), errors.New("Webhook 响应失败: " + resp.Status)
	}
	return string(respBody), nil
}
