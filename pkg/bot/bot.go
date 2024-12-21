package bot

import (
	"github.com/LagrangeDev/LagrangeGo/client"
)

type Account struct {
	Uin      uint32 `json:"uin"`
	Password string `json:"password"`
	SigPath  string `json:"sig_path"`
}

type Config struct {
	LogLevel           string  `json:"log_level"`
	AppInfo            string  `json:"app_info"`
	SignServerUrl      string  `json:"sign_server_url"`
	MusicSignServerUrl string  `json:"music_sign_server_url"`
	CommandPrefix      string  `json:"command_prefix"`
	Account            Account `json:"account"`
	IgnoreSelf         bool    `json:"ignore_self"`
}

type Bot struct {
	*client.QQClient

	Config Config
	Logger Logger
}
