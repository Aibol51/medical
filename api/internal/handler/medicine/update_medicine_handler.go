package medicine

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/medicine"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /medicine/update medicine UpdateMedicine
//
// Update medicine information | 更新Medicine信息
//
// Update medicine information | 更新Medicine信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MedicineInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateMedicineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicineInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicine.NewUpdateMedicineLogic(r.Context(), svcCtx)
		resp, err := l.UpdateMedicine(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
