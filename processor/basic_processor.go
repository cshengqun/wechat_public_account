package processor

import (
	"crypto/sha1"
	"fmt"
	"github.com/cshengqun/wechat_public_account/global"
	"sort"
	"strings"
)

type BasicProcessor interface {
	ValidateRequest(req ValidateRequestReq) (bool, error)
}

type ValidateRequestReq struct {
	Signature string `form:"signature"`
	Timestamp string `form:"timestamp"`
	Nonce string     `form:"nonce"`
	Echostr string   `form:"echostr"`
}

type basicProcessor struct {
}

func NewBasicProcessor() BasicProcessor {
	return &basicProcessor{}
}

func (p *basicProcessor) ValidateRequest(req ValidateRequestReq) (bool, error) {
	tArray := []string{req.Nonce, req.Timestamp, global.GetConfig().AccountInfo.Token}
	sort.Strings(tArray)
	hash := sha1.Sum([]byte(strings.Join(tArray, "")))
	hashStr := fmt.Sprintf("%x", hash)
	return hashStr == req.Signature, nil
}


