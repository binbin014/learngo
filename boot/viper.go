package boot

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"learngin/library/constant"
	"learngin/library/global"
	"os"
)

var Viper = new(_viper)

type _viper struct {
	err  error
	path string
}

func (v *_viper) Initialize(path ...string) {
	if len(path) == 0 {
		flag.StringVar(&v.path, "c", "", "choose config file")
		flag.Parse()
		if v.path == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
				v.path = constant.ConfigFile
			} else {
				v.path = constant.ConfigEnv
			}
		}
	} else {
		v.path = path[0]
	}
	_v := viper.New()
	_v.SetConfigFile(v.path)
	if v.err = _v.ReadInConfig(); v.err != nil {
		panic(fmt.Sprintf(`读取config.yaml, err: %v`, v.err))
	}

	_v.WatchConfig()
	_v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(`配置文件已修改并更新,文件为: `, e.Name)
		if v.err = _v.Unmarshal(&global.Config); v.err != nil {
			fmt.Println(v.err)
		}
	})
	if v.err = _v.Unmarshal(&global.Config); v.err != nil {
		fmt.Println(`Json 序列化数据失败, err :`, v.err)
	}
	global.Viper = _v
}
