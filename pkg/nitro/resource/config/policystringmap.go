package config

var PolicyStringMapFieldNames = struct {
	Comment string
	Count   string
	Name    string
}{
	Comment: "comment",
	Count:   "__count",
	Name:    "name",
}

type PolicyStringmap struct {
	Comment string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count   float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	Name    string  `json:"name,omitempty" nitro:"permission=readwrite"`
}

func (r PolicyStringmap) GetTypeName() string {
	return "policystringmap"
}
