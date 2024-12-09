package config

type PolicyStringmapPatternBinding struct {
	Name  string `json:"name,omitempty" nitro:"permission=readwrite"`
	Key   string `json:"key,omitempty" nitro:"permission=readwrite"`
	Value string `json:"value,omitempty" nitro:"permission=readwrite"`
}

func (r PolicyStringmapPatternBinding) GetTypeName() string {
	return "policystringmap_pattern_binding"
}

func NewPolicyStringmapPatternBindingAddRequest(name string, key string, value string) PolicyStringmapPatternBinding {
	return PolicyStringmapPatternBinding{
		Name:  name,
		Key:   key,
		Value: value,
	}
}
