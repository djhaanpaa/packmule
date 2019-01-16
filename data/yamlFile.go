package data

import (
	"gopkg.in/yaml.v2"
	"os"
)

type YamlFile struct {
	HomeFolder string `yaml:"homeFolder"`
	ServiceName string `yaml:"serviceName"`
	ServiceUrl string `yaml:"serviceUrl"`
	ServiceEnv string `yaml:"serviceEnvironment"`

}

func GenerateExampleYaml(fileName string)  {
	y := YamlFile{
		HomeFolder:"/example",
		ServiceEnv:"Development",
		ServiceName:"serviceName",
		ServiceUrl:"http://localhost"}
	d, err := yaml.Marshal(&y)
	println(string(d))
	if fileName != "" {
		if err == nil {

			w, err := os.Create(fileName)
			if err != nil {
				panic(err)
			} else {
				_, _ = w.WriteString(string(d))
			}
			defer w.Close()
		} else {
			println(err)
		}
	}
}