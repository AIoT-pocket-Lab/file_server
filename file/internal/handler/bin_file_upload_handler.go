package handler

import (
	"net/http"

	"file_server/common/responsex"
	"file_server/file/internal/logic"
	"file_server/file/internal/svc"
	"file_server/file/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func binFileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BinFileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewBinFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.BinFileUpload(&req)
		responsex.Responsex(w, resp, err, "BinFileUpload")
	}
}
