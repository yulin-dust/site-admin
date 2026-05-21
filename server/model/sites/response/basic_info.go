package response

// BasicInfoBrief 对外返回的站点信息，去掉了 GVA_MODEL 中的时间戳和 IP 字段；
// 保留 ID 用于外部调用 /publishStats/addPublishCount?siteId=... 时引用。
// TodayCount / PlanCount 来自 sites_publish_stats，供外部决定是否还要继续推送。
type BasicInfoBrief struct {
	ID          uint    `json:"id"`
	Domain      *string `json:"domain"`
	Url         *string `json:"url"`
	ApiToken    *string `json:"apiToken"`
	CmsType     *string `json:"cmsType"`
	SiteType    *string `json:"siteType"`
	Status      *string `json:"status"`
	CategoryIds *string `json:"categoryIds"`
	AuthorIds   *string `json:"authorIds"`
	TodayCount  *int64  `json:"todayCount" gorm:"column:today_count"`
	PlanCount   *int64  `json:"planCount" gorm:"column:plan_count"`
}
