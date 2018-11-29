package tools

import(
    "net/url"
    "net/http"
    "io/ioutil"
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





