package yconf

import (
	"gopkg.in/yaml.v2"
	"os"
	"ytools/yerr"
)

/**
------------------------------------------------
Created on 2022-11-07 19:43
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

func LoadConf(conf interface{}, confPaths ...string) error {
	for _, confPath := range confPaths {
		cfgFile, err := os.ReadFile(confPath)
		if err != nil {
			return yerr.New(err.Error())
		}
		err = yaml.Unmarshal(cfgFile, conf)
		if err != nil {
			return yerr.New(err.Error())
		}
	}
	return nil
}
