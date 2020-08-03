package common


type CommonRsp struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
}

type Req struct {
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime int64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	Content string `xml:"Content"`
	MsgId uint64 `xml:"MsgId"`
	Event string `xml:"Event"`
}
