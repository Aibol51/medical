package expert

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/expert"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /expert/create expert CreateExpert
//
// Create expert information | 创建Expert信息
//
// Create expert information | 创建Expert信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ExpertInfo
//
// Responses:
//  200: BaseMsgResp

func CreateExpertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpertInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := expert.NewCreateExpertLogic(r.Context(), svcCtx)
		resp, err := l.CreateExpert(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
