package expert

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/expert"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /expert/delete expert DeleteExpert
//
// Delete expert information | 删除Expert信息
//
// Delete expert information | 删除Expert信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteExpertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := expert.NewDeleteExpertLogic(r.Context(), svcCtx)
		resp, err := l.DeleteExpert(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
