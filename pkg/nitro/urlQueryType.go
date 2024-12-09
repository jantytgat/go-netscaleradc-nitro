package nitro

type UrlQueryType struct {
	string
}

func (t UrlQueryType) Prefix() string {
	return t.string + "="
}

var (
	ArgumentsQueryType  = UrlQueryType{"args"}
	FilterQueryType     = UrlQueryType{"filter"}
	AttributesQueryType = UrlQueryType{"attrs"}
)
