package data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type YamlFile struct {
	HomeFolder string `yaml:"homeFolder"`
	ServiceName string `yaml:"serviceName"`
	ServiceUrl string `yaml:"serviceUrl"`
	ServiceEnvironment string `yaml:"serviceEnvironment"`

}

func ImportYamlFile(fileName string) (d YamlFile) {
	b, err := ioutil.ReadFile(fileName)
	if err == nil {
		n := YamlFile{}
		err = yaml.Unmarshal(b, &n)
		if err == nil {

			d = n
		}

	}
	return
}

func GenerateExampleYaml(fileName string)  {
	y := YamlFile{
		HomeFolder:"/example",
		ServiceEnvironment:"Development",
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