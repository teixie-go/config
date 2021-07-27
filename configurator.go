package configurator

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func Load(config interface{}, files ...string) (err error) {
	for _, file := range files {
		if err = loadFile(config, file); err != nil {
			return
		}
	}
	return
}

// 加载目录下所有配置文件
func LoadDir(config interface{}, dir string) (err error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	files := make([]string, 0)
	for _, f := range fileInfos {
		if !f.IsDir() {
			files = append(files, strings.TrimRight(dir, "/")+"/"+f.Name())
		}
	}
	return Load(config, files...)
}

func loadFile(config interface{}, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
		return yaml.Unmarshal(data, config)
	case strings.HasSuffix(file, ".json"):
		return json.Unmarshal(data, config)
	default:
		return errors.New("supported file type:yaml,json")
	}
}
