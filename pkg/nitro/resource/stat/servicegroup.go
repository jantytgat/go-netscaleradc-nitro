package stat

var ServiceGroupFieldNames = struct {
	ClearStats  string
	Name        string
	ServiceType string
	State       string
}{
	ClearStats:  "clearstats",
	Name:        "servicegroupname",
	ServiceType: "servicetype",
	State:       "state",
}

type ServiceGroup struct {
	ClearStats  string `json:"clearstats,omitempty" nitro:"permission=readwrite"`
	Name        string `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	ServiceType string `json:"servicetype,omitempty" nitro:"permission=readonly"`
	State       string `json:"state,omitempty" nitro:"permission=readonly"`
}

func (r ServiceGroup) GetTypeName() string {
	return "servicegroup"
}
