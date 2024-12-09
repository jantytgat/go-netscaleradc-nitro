package config

type PolicyStringmap struct {
	Name    string  `json:"name,omitempty" nitro:"permission=readwrite"`
	Comment string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count   float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r PolicyStringmap) GetTypeName() string {
	return "policystringmap"
}
