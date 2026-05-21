package sites

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BasicInfoApi
	PublishStatsApi
}

var (
	basicInfoService    = service.ServiceGroupApp.SitesServiceGroup.BasicInfoService
	publishStatsService = service.ServiceGroupApp.SitesServiceGroup.PublishStatsService
)
