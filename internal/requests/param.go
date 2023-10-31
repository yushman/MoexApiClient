package requests

const (
	Search = "q"
)

type IParam interface {
	AddParam(key, val string)
	Get() string
}

type Param struct {
	params map[string]string
}

func (p *Param) AddParam(key, val string) {
	p.params[key] = val
}

func (p Param) Get() string {
	result := "?"
	for key, val := range p.params {
		result += key + "=" + val
	}
	return result
}
