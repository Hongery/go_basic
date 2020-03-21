package conf

//LogTransfer 全局配置
type LogTransfer struct {
	KafkaCfg `ini:"kafka"`
	ESCfg	`ini:"es"`
}

type KafkaCfg struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type ESCfg struct {
	Address string `ini:"address"`
	ChanSize int `ini:"chansize"`
	Nums int `ini:"nums"`
}
