package tools

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

/**
 * Get请求
 * @param apiURL string    // 请求URL
 * @return result []byte, err error
 */
func GetRequest(apiURL string) (result []byte, err error) {
	Url, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

/**
 * Post请求
 * @param apiURL string    // 请求URL
 * @return result []byte, err error
 */
func PostRequest(apiURL string, params url.Values) (result []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}
