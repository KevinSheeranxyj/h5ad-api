package model

// AppAccount 机器人配置
type AppAccount struct {
	AppID      string `json:"app_id"`
	AppSecret  string `json:"app_secret"`
	EncryptKey string `json:"encrypt_key"`
}
