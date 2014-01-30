package util

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

/* Get the public AWS hostname for this machine.
 * TODO: Make this a generic utility for getting public info from
 * the aws meta server?
 */
func GetAWSPublicHostname() (hostname string, err error) {
	req := &http.Request{Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Host:   "169.254.169.254",
			Path:   "/latest/meta-data/public-hostname"}}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var hostBytes []byte
		hostBytes, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			hostname = string(hostBytes)
		}
	}
	return
}
