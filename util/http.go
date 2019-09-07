package util

import (
	"bytes"
	"errors"
	"gitee.com/sky_big/mylog"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) ([]byte, error) {
	// 替换空格
	url = strings.Replace(url, " ", "%20", -1)
	mylog.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(url + ":" + resp.Status)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func HttpPost(url string, bs []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", " text/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		mylog.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
