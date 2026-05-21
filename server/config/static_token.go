package config

// StaticToken 通用静态 token 校验中间件的配置。
// 用于 middleware.StaticTokenAuth() 这种"共享密钥"式接口防护。
type StaticToken struct {
	HeaderName string `mapstructure:"header-name" json:"header-name" yaml:"header-name"` // 请求头字段名，留空时默认 "token"
	Value      string `mapstructure:"value" json:"value" yaml:"value"`                   // 期望的 token 值，留空时所有请求都会被拒绝
}
