package DATA

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func RequestHead(Main string, url string, bodys io.Reader, head map[string]string) (*http.Response, string, string, error) {
	/*考虑到后期实用性，将http请求方式、URL、body及HTTP请求头放入变量*/
	resq, err := http.NewRequest(Main, url, bodys)
	if err != nil {
		return nil, "", "", err
	}
	for key, val := range head {
		resq.Header.Add(key, val)
	}
	resqbody, err := httputil.DumpRequest(resq, true)
	if err != nil {
		return nil, "", "", err
	}
	client := http.Client{}
	resp, err := client.Do(resq)
	if err != nil {
		return nil, "", "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, "", "", err
	}
	return resp, string(body), string(resqbody), nil
}
