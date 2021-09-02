package domain

import (
	"testing"

	"github.com/admpub/nging/v3/application/library/ddnsmanager/config"
	"github.com/stretchr/testify/assert"
	"github.com/webx-top/com"
)

func TestDomain(t *testing.T) {
	domains := parseDomainArr([]*config.DNSDomain{
		{Domain: `a.b.c.test.com.cn`},
		{Domain: `w.webx.top`},
		{Domain: `dl.eget.io`},
		{Domain: `webx.top`},
	})
	com.Dump(domains)
	expected := []*Domain{
		{
			DomainName:   "test.com.cn",
			SubDomain:    "a.b.c",
			UpdateStatus: "",
		},
		{
			DomainName:   "webx.top",
			SubDomain:    "w",
			UpdateStatus: "",
		},
		{
			DomainName:   "eget.io",
			SubDomain:    "dl",
			UpdateStatus: "",
		},
		{
			DomainName:   "webx.top",
			SubDomain:    "",
			UpdateStatus: "",
		},
	}
	assert.Equal(t, expected, domains)
}
