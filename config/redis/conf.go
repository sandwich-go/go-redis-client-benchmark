package redis

type Mode = string

const (
	ModeSerial   = "serial"
	ModeParallel = "parallel"
)

//go:generate optiongen --new_func=NewConf --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"Addrs":      []string{"127.0.0.1:6379"}, // @MethodComment(Redis地址列表)
		"Cluster":    false,                      // @MethodComment(是否为Redis集群，默认为false，集群需要设置为true)
		"MasterName": "",                         // @MethodComment(哨兵模式master名)
		"KeySizes":   []int{16},                  // @MethodComment(key的大小)
		"ValueSizes": []int{64, 256, 1024},       // @MethodComment(value的大小)
		"PoolSizes":  []int{100, 1000},           // @MethodComment(连接池的大小)
		"Mode":       Mode(ModeSerial),           // @MethodComment(运行模式)
	}
}
