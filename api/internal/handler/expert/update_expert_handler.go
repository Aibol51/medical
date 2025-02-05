package expert

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/expert"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /expert/update expert UpdateExpert
//
// Update expert information | 更新Expert信息
//
// Update expert information | 更新Expert信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: ExpertInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateExpertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpertInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := expert.NewUpdateExpertLogic(r.Context(), svcCtx)
		resp, err := l.UpdateExpert(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
