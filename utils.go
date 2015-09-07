package twitter

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

var Utils = TwitterUtils{}

type TwitterUtils struct{}

func (u TwitterUtils) Parse(r io.Reader, v interface{}) (string, error) {
	bits, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bits, &v)
	if err != nil {
		return string(bits), err
	}
	return string(bits), err
}

func (u TwitterUtils) ConvertTimestamp(in_timestamp string) (timestamp time.Time, timestampString string) {
	var timeformat string = time.RubyDate
	timestamp, _ = time.Parse(timeformat, in_timestamp)
	return timestamp, timestamp.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("2006/01/02 15:04:05")
}
