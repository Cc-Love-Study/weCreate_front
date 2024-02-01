package conf

type AppConf struct {
	RunConf    `ini:"run"`
	ServerConf `ini:"server"`
}

func NewAppConf() *AppConf {
	return &AppConf{}
}

type RunConf struct {
	Address string `ini:"address"`
	Md5Key  string `ini:"md5key"`
}

type ServerConf struct {
	Address string `ini:"address"`
}
