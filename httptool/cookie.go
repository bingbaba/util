package httptool

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	cookieDir = "/tmp"
)

func SetCookieDir(set_path string) {
	cookieDir = filepath.Dir(set_path)
}

func GetCookieFile(host string) string {
	return filepath.Join(cookieDir, host)
}

func AddCookies(req *http.Request, host string) error {
	fi, err := os.Open(GetCookieFile(host))
	if err != nil {
		return err
	}
	defer fi.Close()

	f_bytes, read_err := ioutil.ReadAll(fi)
	if read_err != nil {
		return read_err
	}

	cookies := make([]*http.Cookie, 0)
	json_err := json.Unmarshal(f_bytes, &cookies)
	if json_err != nil {
		return json_err
	}

	for _, cookie := range cookies {
		//fmt.Println(cookie)
		req.AddCookie(cookie)
	}

	cookie := &http.Cookie{Name: "noticeLoginFlag", Value: "1"}
	req.AddCookie(cookie)
	return nil
}

func SaveCookies(resp *http.Response, host string) error {
	cookie_bytes, cookie_err := json.Marshal(resp.Cookies())
	if cookie_err != nil {
		return cookie_err
	}
	cookie_w_err := ioutil.WriteFile(GetCookieFile(host), cookie_bytes, 0666)
	if cookie_w_err != nil {
		return cookie_w_err
	}

	return nil
}
