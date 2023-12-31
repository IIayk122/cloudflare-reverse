package cleanhttp

import (
	"net/url"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

// create http client and return *CleanHttp. Take *Config as params.
func NewCleanHTTPClient(config *Config) (*CleanHTTP, error) {
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(20),
		tls_client.WithClientProfile(GetTLSProfile()),
		tls_client.WithInsecureSkipVerify(),
		tls_client.WithCookieJar(tls_client.NewCookieJar()),
	}

	if config.Proxy != "" {

		options = append(options, tls_client.WithProxyUrl(config.Proxy))
	}

	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		return nil, err
	}

	return &CleanHTTP{
		Config:    config,
		TLSClient: client,
		Cookies:   nil,
	}, nil
}

// Do request and return *http.Response, Take RequestOption in params.
func (c *CleanHTTP) Do(request RequestOption) (*http.Response, error) {
	if request.Header == nil {
		request.Header = c.GetDefaultHeader()
	}

	/* Give http 400*/
	/*
		if request.Body != nil {

			len, e := CalculateContentLength(request.Body)

			if e == nil {
				request.Header.Add("content-length", strconv.Itoa(int(len)))
			}
		}*/

	req, err := http.NewRequest(request.Method, request.URL, request.Body)
	if err != nil {
		return nil, err
	}

	req.Header = request.Header

	resp, err := c.TLSClient.Do(req)
	if err != nil {
		return nil, err
	}

	u, _ := url.Parse(request.URL) //nolint: errcheck
	cooks := c.TLSClient.GetCookies(u)
	c.Cookies = append(c.Cookies, cooks...)

	return resp, nil
}

// Take all cookies from client and return them as a header format string.
func (c *CleanHTTP) FormatCookies() string { // fuck cookiejar ok ?
	var builder strings.Builder
	lastIndex := len(c.Cookies) - 1

	for i, cookie := range c.Cookies {
		builder.WriteString(cookie.Name)
		builder.WriteString("=")
		builder.WriteString(cookie.Value)

		if i != lastIndex {
			builder.WriteString("; ")
		}
	}

	return builder.String()
}
