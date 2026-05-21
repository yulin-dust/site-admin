
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PublishStatsSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      StartTodayCount  *int  `json:"startTodayCount" form:"startTodayCount"`
EndTodayCount  *int  `json:"endTodayCount" form:"endTodayCount"`
    request.PageInfo
}
