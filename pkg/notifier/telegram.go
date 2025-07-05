package notifier

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"synapse/internal/model"
	"time"
)

func SendTelegramMessage(cfg model.TelegramConfig, message string) error {
	if cfg.BotToken == "" || cfg.ChatID == "" {
		return errors.New("Token 和 ChatID 不能为空")
	}

	apiURL := "https://api.telegram.org/bot" + cfg.BotToken + "/sendMessage"
	body := map[string]interface{}{
		"chat_id":    cfg.ChatID,
		"text":       message,
		"parse_mode": cfg.ParseMode,
	}
	jsonBody, _ := json.Marshal(body)

	client := &http.Client{Timeout: 10 * time.Second}
	if cfg.Proxy != "" {
		proxyURL, err := url.Parse(cfg.Proxy)
		if err != nil {
			return errors.New("代理地址格式错误")
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Telegram API 响应失败: " + resp.Status)
	}
	return nil
}
