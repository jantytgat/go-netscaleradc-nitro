package config

var LbVserverServiceGroupBindingFieldNames = struct {
	Count            string
	Name             string
	Order            string
	ServiceGroupName string
	ServiceName      string
	StateFlag        string
	Weight           string
}{
	Count:            "__count",
	Name:             "name",
	Order:            "order",
	ServiceGroupName: "servicegroupname",
	ServiceName:      "servicename",
	StateFlag:        "stateflag",
	Weight:           "weight",
}

type LbVserverServiceGroupBinding struct {
	Order            float64 `json:"order,omitempty" nitro:"permission=readwrite"`
	Weight           float64 `json:"weight,omitempty" nitro:"permission=readwrite"`
	Name             string  `json:"name,omitempty" nitro:"permission=readwrite"`
	ServiceName      string  `json:"servicename,omitempty" nitro:"permission=readwrite"`
	ServiceGroupName string  `json:"servicegroupname,omitempty" nitro:"permission=readwrite"`
	StateFlag        string  `json:"stateflag,omitempty" nitro:"permission=readonly"`
	Count            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (t LbVserverServiceGroupBinding) GetTypeName() string {
	return "lbvserver_servicegroup_binding"
}
