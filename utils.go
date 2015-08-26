package twitter

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

var Utils = TwitterUtils{}

type TwitterUtils struct{}

func (u TwitterUtils) Parse(r io.Reader, v interface{}) error {
	bits, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bits, &v)
	if err != nil {
		return err
	}
	return err
}
