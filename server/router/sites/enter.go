package sites

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BasicInfoRouter
	PublishStatsRouter
}

var (
	basicInfoApi    = api.ApiGroupApp.SitesApiGroup.BasicInfoApi
	publishStatsApi = api.ApiGroupApp.SitesApiGroup.PublishStatsApi
)
