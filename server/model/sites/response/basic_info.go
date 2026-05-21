package response

// BasicInfoBrief 对外返回的站点信息，去掉了 GVA_MODEL（id/时间戳）和 IP 字段。
type BasicInfoBrief struct {
	Domain      *string `json:"domain"`
	Url         *string `json:"url"`
	ApiToken    *string `json:"apiToken"`
	CmsType     *string `json:"cmsType"`
	SiteType    *string `json:"siteType"`
	Status      *string `json:"status"`
	CategoryIds *string `json:"categoryIds"`
	AuthorIds   *string `json:"authorIds"`
}
