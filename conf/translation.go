package conf

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

//Dictionary is a singleton for recoding the error we would get
//when try to validate the struct we want to validate by using
//valisate.v9 package.(we would use yaml to recoding the information.)
var Dictionary *map[interface{}]interface{}

// LoadLocales get the content we write in yaml
func LoadLocales(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}

	Dictionary = &m

	return nil
}

//Translate is use to translate error to languate peaple can read.
func Translate(key string) (result string) {
	dic := *Dictionary
	keys := strings.Split(key, ".")
	for index, path := range keys {
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); ok == false {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}

	return ""
}
