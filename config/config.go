/*
 * @Author: SpenserCai
 * @Date: 2023-08-16 11:05:40
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-09-29 19:17:11
 * @Description: file content
 */
package config

type ServerItem struct {
	Name          string `json:"name"`
	Host          string `json:"host"`
	MaxConcurrent int    `json:"max_concurrent"`
	MaxQueue      int    `json:"max_queue"`
	MaxVRAM       string `json:"max_vram"`
}

type DefaultSetting struct {
	CfgScale       float64 `json:"cfg_scale"`
	NegativePrompt string  `json:"negative_prompt"`
	Sampler        string  `json:"sampler"`
	Height         int64   `json:"height"`
	Width          int64   `json:"width"`
	Steps          int64   `json:"steps"`
	Model          string  `json:"sd_model_checkpoint"`
	Vae            string  `json:"sd_vae"`
	ClipSkip       int64   `json:"clip_skip"`
}

type DbConfig struct {
	Type string `json:"type"`
	DSN  string `json:"dsn"`
}

type UserCenter struct {
	Enable       bool     `json:"enable"`
	DbConfig     DbConfig `json:"db_config"`
	MustRegister bool     `json:"must_register"`
}

type WebSite struct {
	Api struct {
		Host      string `json:"host"`
		Port      int    `json:"port"`
		JwtSecret string `json:"jwt_secret"`
	} `json:"api"`
}

type Config struct {
	SDWebUi struct {
		Servers        []ServerItem   `json:"servers"`
		DefaultSetting DefaultSetting `json:"default_setting"`
	} `json:"sd_webui"`
	Discord struct {
		Token     string `json:"token"`
		ServerId  string `json:"server_id"`
		BotName   string `json:"bot_name"`
		BotAvatar string `json:"bot_avatar"`
		BotUrl    string `json:"bot_url"`
	} `json:"discord"`
	UserCenter           UserCenter `json:"user_center"`
	DisableReturnGenInfo bool       `json:"disable_return_gen_info"`
	WebSite              WebSite    `json:"website"`
}
