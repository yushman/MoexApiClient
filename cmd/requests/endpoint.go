package requests

const (
	Securities    = "securities/"
	History       = "history/"
	Markets       = "markets/"
	Securitygoups = "securitygoups/"
	Secutitytypes = "securitytypes/"
	Statistics    = "statistics/"
	Sitenews      = "sitenews/"
)

type IEndpoint interface {
	MakeUrl() string
	addUrl(url string) *Endpoint
}

type Endpoint struct {
	url string
	end IEndpoint
}

func (r *Endpoint) addUrl(url string) *Endpoint {
	return &Endpoint{
		url: url + "/",
		end: r,
	}
}

func (r *Endpoint) MakeUrl() string {
	if r.end != nil {
		return r.end.MakeUrl() + r.url
	} else {
		return r.url
	}
}

func (r Endpoint) String() string {
	return r.MakeUrl()
}
