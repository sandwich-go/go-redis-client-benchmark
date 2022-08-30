package config

import (
	"bytes"
	"fmt"
	"github.com/sandwich-go/go-redis-client-benchmark/config/redis"
	"github.com/sandwich-go/xconf"
)

//go:generate optiongen --option_with_struct_name=false --new_func=NewConfig --xconf=true --empty_composite_nil=true --usage_tag_name=usage
func ConfigOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		// annotation@Redis(getter="redis.ConfInterface")
		"Redis": (*redis.Conf)(redis.NewConf()),
	}
}

func loadYamlFile(valPtr interface{}, files ...string) error {
	x := xconf.New(xconf.WithFiles(files...))
	if err := x.Parse(valPtr); err != nil {
		return err
	}
	x.Usage()
	return nil
}

func printYamlConfig(valPtr interface{}) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	xconf.MustSaveVarToWriter(valPtr, xconf.ConfigTypeYAML, bytesBuffer)
	fmt.Println(bytesBuffer)
	bytesBuffer.Reset()
}

func ReadConfigFiles(files ...string) (*Config, error) {
	var c = new(Config)
	err := loadYamlFile(c, files...)
	if err != nil {
		return nil, err
	}
	printYamlConfig(c)
	return c, nil
}
