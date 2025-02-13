package nitro

var (
	ResourceTypeUnknown = ResourceType{""}
	ResourceTypeConfig  = ResourceType{"config"}
	ResourceTypeStat    = ResourceType{"stat"}
)

type ResourceType struct {
	string
}

func (r ResourceType) UrlPath() string {
	return "/nitro/v1/" + r.string + "/"
}
