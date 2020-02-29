package sdk

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"

	dao "url-shortener-go/dao"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// LinkEvent 链接事件
type LinkEvent struct {
	ActingUser string                 `json:"actingUser"`
	TargetUser string                 `json:"targetUser"`
	Value      map[string]interface{} `json:"value"`
	Event      MQEvent                `json:"event"`
}

// LinkDTO LinkDTO

// MQEvent MQ事件类型
type MQEvent struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func toLinkMesssageBody(link dao.DBLink) []byte {
	linkDTO := make(map[string]interface{})
	linkDTO["id"] = link.ID
	linkDTO["url"] = link.URL
	linkDTO["openId"] = link.OpenID
	linkDTO["shortCode"] = link.ShortCode
	linkDTO["createTime"] = link.CreateTime
	linkDTO["updateTime"] = link.UpdateTime
	linkEvent := LinkEvent{
		ActingUser: "system",
		TargetUser: link.OpenID,
		Value:      linkDTO,
		Event: MQEvent{
			Name: "share_link_record",
			Type: "share_link_record",
		},
	}
	bodyText, _ := json.Marshal(linkEvent)
	return bodyText
}

// SendShareLinkMessage 发送短链消息
func SendShareLinkMessage(link dao.DBLink) {
	mqConfig := os.Getenv("MQ_CONF")
	conn, err := amqp.Dial(mqConfig)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	failOnError(err, "Failed to declare a queue")
	bodyText := toLinkMesssageBody(link)
	err = ch.Publish(
		"motivation_message", // exchange
		"user.motivation",    // routing key
		false,                // mandatory
		false,                // immediate
		amqp.Publishing{
			ContentType: "text/json",
			Body:        []byte(bodyText),
		})
	log.Printf(" [x] Sent %s", bodyText)
	failOnError(err, "Failed to publish a message")
}
