package processor

import (
	"github.com/cshengqun/wechat_public_account/constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MsgProcessor interface {
	DealWithTextMsg(c *gin.Context, req MsgReq)
}

type msgProcessor struct {
}

func NewMsgProcessor() MsgProcessor {
	return &msgProcessor{}
}

type MsgReq struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
	MsgId uint64 `xml:"MsgId"`
}

type TextMsgRsp struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
}

func (p *msgProcessor) DealWithTextMsg(c *gin.Context, req MsgReq) {
	c.XML(http.StatusOK, &TextMsgRsp{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      constant.MsgTypeText,
		Content:      req.Content,
	})
}
