package client

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"
)

func CreateHTTPClient(appserviceURL string, insecureSkipVerify bool) *http.Client {
	if strings.HasPrefix(appserviceURL, "/") {
		return &http.Client{
			Timeout: time.Second * 30,
			Transport: &http.Transport{
				DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
					return net.Dial("unix", appserviceURL)
				},
			},
		}
	} else {
		return &http.Client{
			Timeout: time.Second * 30,
			Transport: &http.Transport{
				DisableKeepAlives: true,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: insecureSkipVerify,
				},
				Proxy: http.ProxyFromEnvironment,
			},
		}
	}
}
