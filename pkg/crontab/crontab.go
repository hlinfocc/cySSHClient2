package crontab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	hostextent "github.com/hlinfocc/cySSHClient2/pkg/dao/hostExtent"
)

// 钉钉机器人消息结构体
type DingTalkMessage struct {
	MsgType  string   `json:"msgtype"`
	Link     Link     `json:"link"`
	Markdown Markdown `json:"markdown"`
}

type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicURL     string `json:"picUrl"`
	MessageURL string `json:"messageUrl"`
}
type Markdown struct {
	Text  string `json:"text"`
	Title string `json:"title"`
}

func StartCrond(webhookUrl string) {
	s := gocron.NewScheduler(time.UTC)

	// 每天3:30执行
	_, _ = s.Every(1).Day().At("3:30").Do(func() {
		list, _, err := hostextent.Querylist(0, 0, 1)
		if err == nil {
			for i := 0; i < len(list); i++ {
				item := list[i]
				days, derr := calculateDaysUntilExpiration(item.EndTime)
				if derr != nil {
					continue
				}
				if days > 0 && days < 30 {
					sendMain(webhookUrl, item.Host, item.CloudType, item.EndTime)
				}
			}
		}
	})

	s.StartAsync() // 非阻塞启动
}

func sendMain(webhookUrl string, host string, cloudType string, endTime string) {
	// 构建消息内容
	message := DingTalkMessage{
		MsgType: "markdown",
		Markdown: Markdown{
			Text:  fmt.Sprintf("## 主机监控提醒 \n >您的%s主机[%s]将于〔%s〕过期，请注意续费。", host, cloudType, endTime),
			Title: "主机监控提醒",
		},
	}

	// 发送消息
	err := sendDingTalkMessage(webhookUrl, message)
	if err != nil {
		log.Printf("发送钉钉消息失败: %v\n", err)
	} else {
		log.Println("钉钉消息发送成功!")
	}
}

// 发送钉钉消息函数
func sendDingTalkMessage(webhookURL string, message DingTalkMessage) error {
	// 序列化消息为JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("JSON序列化失败: %w", err)
	}

	// 发送HTTP POST请求
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("钉钉API返回错误状态码: %d", resp.StatusCode)
	}

	// 解析响应内容（可选）
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("解析响应失败: %w", err)
	}

	// 检查钉钉返回的错误码
	if result["errcode"] != float64(0) {
		return fmt.Errorf("钉钉返回错误: %v", result["errmsg"])
	}

	return nil
}

func calculateDaysUntilExpiration(endTime string) (int, error) {
	// 定义时间格式（Go使用特定的参考时间格式）
	layout := "2006-01-02 15:04:05"

	// 解析过期时间
	expirationTime, err := time.Parse(layout, endTime)
	if err != nil {
		return 0, fmt.Errorf("解析过期时间失败: %v", err)
	}

	// 获取当前时间
	currentTime := time.Now()

	// 计算时间差
	duration := expirationTime.Sub(currentTime)

	// 转换为天数（24小时制）
	days := int(duration.Hours() / 24)

	return days, nil
}
