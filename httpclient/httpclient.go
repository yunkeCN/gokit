package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
)

var tracingClient = apmhttp.WrapClient(http.DefaultClient)

func Get(ctx context.Context, url string, header map[string]string) (respBody []byte, status int, err error) {
	return DoJson(ctx, "GET", url, nil, header)
}

func PostJson(ctx context.Context, url string, body interface{}, header map[string]string) (respBody []byte, status int, err error) {
	return DoJson(ctx, "POST", url, body, header)
}

func PutJson(ctx context.Context, url string, body interface{}, header map[string]string) (respBody []byte, status int, err error) {
	return DoJson(ctx, "PUT", url, body, header)
}

func PostForm(ctx context.Context, url string, data url.Values, header map[string]string) (respBody []byte, status int, err error) {
	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, 0, errors.Wrapf(err, "new request error. url=%s, data=%+v", url, data)
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := ctxhttp.Do(ctx, tracingClient, request)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "do request error. url=%s, data=%+v", url, data)
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "read resp body error. url=%s, data=%+v", url, data)
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, resp.StatusCode, errors.Errorf("do return error. code=%d, url=%s, body=%s", resp.StatusCode, url, respBody)
	}

	return respBody, resp.StatusCode, nil
}

func Delete(ctx context.Context, url string, header map[string]string) (respBody []byte, status int, err error) {
	return DoJson(ctx, "DELETE", url, nil, header)
}

func DoJson(ctx context.Context, method string, url string, body interface{}, header map[string]string) (respBody []byte, status int, err error) {
	var bodyBytes []byte
	if body != nil {
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			return respBody, status, errors.Wrapf(err, "json marshal body error. method=%s, url=%s, body=%v", method, url, body)
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return
	}

	for key, value := range header {
		request.Header.Set(key, value)
	}

	request.Header.Set("Content-Type", "application/json")
	resp, err := ctxhttp.Do(ctx, tracingClient, request)
	if err != nil {
		err = errors.Wrapf(err, "do request error. method=%s, url=%s", method, url)
		return
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "read resp body error. method=%s, url=%s", method, url)
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, resp.StatusCode, errors.Errorf("do return error. code:%d, method=%s, url=%s, resp=%s", resp.StatusCode, method, url, respBody)
	}

	return respBody, resp.StatusCode, nil
}
