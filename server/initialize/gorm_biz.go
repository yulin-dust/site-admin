package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sites"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(sites.BasicInfo{}, sites.PublishStats{})
	if err != nil {
		return err
	}
	return nil
}
