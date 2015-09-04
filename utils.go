package twitter

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
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

func (u TwitterUtils) ConvertTimestamp(timestamp string) string {
	var timeformat string = time.RubyDate
	tm_timestamp, _ := time.Parse(timeformat, timestamp)
	return tm_timestamp.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("2006/01/02 15:04:05")
}

func (u TwitterUtils) ConvertText(status Status) string {
	return status.Text
}
