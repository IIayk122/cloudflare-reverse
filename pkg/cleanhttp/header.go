package cleanhttp

import (
	"fmt"
	"strings"

	http "github.com/bogdanfinn/fhttp"
	"github.com/mssola/user_agent"
)

func GenerateAcceptLanguageHeader(languages []string) string {
	var headerContent strings.Builder

	for i, lang := range languages {
		if i > 0 {
			headerContent.WriteString(",")
		}
		headerContent.WriteString(lang)

		if i > 0 {
			quality := float64(1.0) - float64(i)*0.1
			headerContent.WriteString(";q=")
			headerContent.WriteString(fmt.Sprintf("%.1f", quality))
		}
	}

	return headerContent.String()
}

// Parse user-agent informations and return *UserAgentInfo.
func ParseUserAgent(userAgentString string) *UserAgentInfo {
	ua := user_agent.New(userAgentString)
	name, version := ua.Browser()

	info := &UserAgentInfo{
		BrowserName:    name,
		BrowserVersion: version,
		OSName:         ua.OSInfo().Name,
		OSVersion:      ua.OSInfo().Version,
		UaVersion:      strings.Split(version, ".")[0],
	}
	return info
}

func (c *CleanHTTP) GenerateBaseHeaders() *HeaderBuilder {
	ua := ParseUserAgent(c.Config.BrowserFp.Navigator.UserAgent)

	h := &HeaderBuilder{
		SecChUa:         fmt.Sprintf(`"Not.A/Brand";v="8", "Chromium";v="%s", "Google Chrome";v="%s"`, ua.UaVersion, ua.UaVersion),
		SecChUaPlatform: fmt.Sprintf(`"%s"`, ua.OSName), // need to fix with apple
		secChUaMobile:   "?0",                           // todo,
		AcceptLanguage:  GenerateAcceptLanguageHeader(c.Config.BrowserFp.Navigator.Languages),
		Cookies:         c.FormatCookies(),
	}

	return h
}

func (c *CleanHTTP) GetDefaultHeader() http.Header {
	baseHeader := c.GenerateBaseHeaders()

	return http.Header{
		"cache-control":             {`max-age=0`},
		"sec-ch-ua":                 {baseHeader.SecChUa},
		"sec-ch-ua-mobile":          {baseHeader.secChUaMobile},
		"sec-ch-ua-platform":        {baseHeader.SecChUaPlatform},
		"upgrade-insecure-requests": {`1`},
		"user-agent":                {c.Config.BrowserFp.Navigator.UserAgent},
		"accept":                    {`text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng;q=0.8,application/signed-exchange;v=b3;q=0.7`},
		"sec-fetch-site":            {`none`},
		"sec-fetch-mode":            {`navigate`},
		"sec-fetch-user":            {`?1`},
		"sec-fetch-dest":            {`document`},
		"accept-encoding":           {`gzip, deflate, br`},
		"accept-language":           {baseHeader.AcceptLanguage},
		"cookie":                    {baseHeader.Cookies},

		http.HeaderOrderKey: {
			"cache-control",
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"sec-ch-ua-platform",
			"upgrade-insecure-requests",
			"user-agent",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-user",
			"sec-fetch-dest",
			"accept-encoding",
			"accept-language",
			"cookie",
		},
	}
}
