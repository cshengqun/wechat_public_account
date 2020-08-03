package processor

import (
	"github.com/cshengqun/wechat_public_account/common"
	"github.com/gin-gonic/gin"
)

type EventProcessor interface {
	DealWithSubscribe(c *gin.Context, req common.Req)
}

type eventProcessor struct {
}

func NewEventProcessor() EventProcessor {
	return &eventProcessor{}
}

type EventReq struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Event string `xml:"Event"`
}

func (p *eventProcessor) DealWithSubscribe(c *gin.Context, req common.Req) {

}