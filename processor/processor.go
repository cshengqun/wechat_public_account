package processor

import (
	"fmt"
	"github.com/cshengqun/wechat_public_account/common"
	"github.com/cshengqun/wechat_public_account/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateWechatReq(c *gin.Context) {
	var req ValidateRequestReq
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Printf("invalid parameter, err:%v\n", err)
		c.JSON(http.StatusOK, &common.CommonRsp{
			Code : constant.CodeInvalidReq,
			Msg: "invalid req",
		})
		return
	}
	fmt.Printf("rcv req:%+v", req)
	valid, err := NewBasicProcessor().ValidateRequest(req)
	if err != nil {
		fmt.Printf("validate request error:%v\n", err)
		c.JSON(http.StatusOK, &common.CommonRsp{
			Code : constant.CodeInvalidReq,
			Msg: "invalid req",
		})
		return
	}
	if valid {
		c.String(http.StatusOK, req.Echostr)
		return
	}
	c.JSON(http.StatusForbidden, &common.CommonRsp{
		Code: constant.CodeInvalidReq,
		Msg:  "invalid req",
	})
}

func DealWithMsg(c *gin.Context) {
	var req common.Req
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Printf("invalid parameter, err:%v\n", err)
		c.JSON(http.StatusOK, &common.CommonRsp{
			Code : constant.CodeInvalidReq,
			Msg: "invalid req",
		})
		return
	}
	fmt.Printf("rcv req:%+v\n", req)
	switch req.MsgType {
	case constant.MsgTypeText:
		NewMsgProcessor().DealWithTextMsg(c, req)
	case constant.MsgTypeEvent:
		switch req.Event {
		case constant.EventSubscribe:

		}
	default:
		fmt.Printf("unsupport msg type ignore")
		c.String(http.StatusOK, "success")
	}

}