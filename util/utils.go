package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func PanicIfError(msg string, err error) {
	if err != nil {
		log.Panicf("%v, err=%v", msg, err)
	}
}

func LoadJSON(filename string, v interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	return err
}
