package responsex

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
    Code int         `json:"code"`
	Type string		 `json:"type"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data,omitempty"`
}

func Responsex(w http.ResponseWriter, resp interface{}, err error, ty string) {
    var body Body
    if err != nil {
        body.Code = -1
		body.Type = ty

        if strings.Contains(err.Error(), "connection error") {
            body.Msg = "服务断开连接"
        } else if strings.Contains(err.Error(), "panic") {
            body.Msg = "服务异常"
        } else {
            before, after, found := strings.Cut(err.Error(), " desc = ")//rpc error: code = Unknown desc = xxx
            if found {
                body.Msg = after
            } else {
                body.Msg = before
            }
        }

        httpx.WriteJson(w, 500, body)
    } else {
		body.Code = 1
		body.Type = ty
        body.Msg = "SUCCESS"
        body.Data = resp
        httpx.OkJson(w, body)
    }
}