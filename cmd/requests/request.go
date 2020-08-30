package requests

import "net/http"

const (
	base            = "https://iss.moex.com/iss/"
	JSON ResultType = 0
	XML  ResultType = 1
	CSV  ResultType = 2
	HTML ResultType = 3
)

type request struct {
	IEndpoint
	IParam
}

type ResultType int

func New() *request {
	return &request{
		nil,
		&Param{make(map[string] string)},
	}
}

func (r *request) Execute() (*http.Response, error) {
	return http.Get(r.GetUrl())
}

func (r *request) ExexuteWithType(rt ResultType) (*http.Response, error) {
	return http.Get(r.GetUrlWithType(rt))
}

func (r *request) GetUrl() string {
	return base + r.IEndpoint.MakeUrl() + JSON.String() + r.IParam.Get()
}

func (r *request) GetUrlWithType(rt ResultType) string {
	return base + r.IEndpoint.MakeUrl() + rt.String() + r.IParam.Get()
}

func (r *request) AddQueryParam(q string) IParam {
	r.IParam.AddParam(Search, q)
	return r.IParam
}

func (r *request) AddEndpoint(url string) {
	r.IEndpoint = r.IEndpoint.addUrl(url)
}

func (r *request) NewSecurities() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Securities,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *request) NewHistory() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: History,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *request) NewMarkets() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Markets,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *request) NewSecurityGroups() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Securitygoups,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *request) NewSecurityTypes() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Secutitytypes,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *request) NewStatistics() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Statistics,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r *request) NewSitenews() IEndpoint {
	r.IEndpoint = &Endpoint{
		url: Sitenews,
		end: r.getReqSafe(),
	}
	return r.IEndpoint
}

func (r request) getReqSafe() IEndpoint {
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
