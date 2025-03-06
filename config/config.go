package config

import (
	"it-cv/pkg/helper"
	"os"

	"gopkg.in/yaml.v2"
)

// Настройки
//
// # file - передавать без разширения
//
// сам файл с настройками должен быть с разширением .yaml
//
// приватный файл -local.yaml - переопределяет указанные настройки в файле без -local
//
//	type Setting struct {
//		Host string `yaml:"host"`
//		Port int    `yaml:"port"`
//	}
//
//	set := &Setting{}
//	// file: "/app/main.yaml" and "/app/main-local.yaml" (optional)
//	err := config.Get("/app/main", set);
//	if err != nil {
//		panic(err)
//	}
func Get(file string, out any) error {
	err := get(file+".yaml", out)
	if helper.FileExists(file + "-local.yaml") {
		get(file+"-local.yaml", out)
	}
	return err
}

func get(file string, out any) error {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, out)
	if err != nil {
		return err
	}
	return nil
}
