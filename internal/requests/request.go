package requests

import "net/http"

const (
	base            = "https://iss.moex.com/iss/"
	JSON ResultType = 0
	XML  ResultType = 1
	CSV  ResultType = 2
	HTML ResultType = 3
)

type Request struct {
	IEndpoint
	IParam
}

type ResultType int

func New() *Request {
	return &Request{
		nil,
		&Param{make(map[string]string)},
	}
}

func (r *Request) Execute() (*http.Response, error) {
	return http.Get(r.GetUrl())
}

func (r *Request) ExexuteWithType(rt ResultType) (*http.Response, error) {
	return http.Get(r.GetUrlWithType(rt))
}

func (r *Request) GetUrl() string {
	return base + r.IEndpoint.MakeUrl() + JSON.String() + r.IParam.Get()
}

func (r *Request) GetUrlWithType(rt ResultType) string {
	return base + r.IEndpoint.MakeUrl() + rt.String() + r.IParam.Get()
}

func (r *Request) AddQueryParam(q string) IParam {
	r.IParam.AddParam(Search, q)
	return r.IParam
}

func (r *Request) AddEndpoint(url string) {
	r.IEndpoint = r.IEndpoint.addUrl(url)
}

func (r *Request) NewSecurities() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Securities,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewSecurity(s string) IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Securities,
		end: r.getReqSafe(),
	}
	r.AddEndpoint(s)
	return r.IEndpoint
}

func (r *Request) NewEngines() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Engines,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewEngine(s string) IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Engines,
		end: r.getReqSafe(),
	}
	r.AddEndpoint(s)
	return r.IEndpoint
}

func (r *Request) NewHistory() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: History,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewMarkets() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Markets,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewMarket(s string) IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Markets,
		end: r.getReqSafe(),
	}
	r.AddEndpoint(s)
	return r.IEndpoint
}

func (r *Request) NewSecurityGroups() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Securitygoups,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewSecurityTypes() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Secutitytypes,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewStatistics() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Statistics,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) NewSitenews() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Sitenews,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *Request) getReqSafe() IEndpoint {
	if r.IEndpoint != nil {
		return r.IEndpoint
	} else {
		return nil
	}
}

func (t ResultType) String() string {
	types := [...]string{
		".json",
		".xml",
		".csv",
		".html",
	}

	if t < JSON || t > HTML {
		return types[JSON]
	}
	return types[t]
}
