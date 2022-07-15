package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}
type KafkaConf struct {
	Address string `ini:"address"`
}
type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
}

// --------unuser---------
type TaillogConf struct {
	FileName string `ini:"filename"`
}
