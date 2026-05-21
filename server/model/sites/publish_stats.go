// 自动生成模板PublishStats
package sites

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 网站更新数据 结构体  PublishStats
type PublishStats struct {
	global.GVA_MODEL
	SiteId          *int32     `json:"siteId" form:"siteId" gorm:"uniqueIndex;comment:关联网站基础表，唯一;column:site_id;" binding:"required"` //站点ID
	TodayCount      *int64     `json:"todayCount" form:"todayCount" gorm:"index;comment:当天发布数量，零点重置;column:today_count;"`             //今日发布数
	PlanCount       *int64     `json:"planCount" form:"planCount" gorm:"comment:待发布/排队中数量;column:plan_count;"`                        //计划发布数
	TotalCount      *int64     `json:"totalCount" form:"totalCount" gorm:"comment:历史总量;column:total_count;"`                          //总发布数
	LastPublishTime *time.Time `json:"lastPublishTime" form:"lastPublishTime" gorm:"comment:最近一次成功发布;column:last_publish_time;"`      //最后发布时间
	YesterdayCount  *int64     `json:"yesterdayCount" form:"yesterdayCount" gorm:"comment:昨天的发布数;column:yesterday_count;"`            //昨日发布数
}

// TableName 网站更新数据 PublishStats自定义表名 sites_publish_stats
func (PublishStats) TableName() string {
	return "sites_publish_stats"
}
