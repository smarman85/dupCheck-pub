package readYaml

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Vars struct {
	NonSecrets map[string]string `yaml:"envVars"`
	Secrets    map[string]string `yaml:"secretEnvs"`
}

func GetEnvValues(fileName string) (map[string]string, map[string]string) {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var contents Vars
	err = yaml.Unmarshal(yamlFile, &contents)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return contents.Secrets, contents.NonSecrets
}

//func Run() {
//	commonSecrets, commonNonSecrets := getCommonValues("../chime-cd/overrides/apps/prod.yaml")
//	fmt.Println(commonNonSecrets)
//	fmt.Println(commonSecrets)
//}
