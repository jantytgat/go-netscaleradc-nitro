package config

var SystemUserFieldNames = struct {
	AllowedManagementInterface     string
	AllowedManagementInterfaceKind string
	Count                          string
	Encrypted                      string
	ExternalAuth                   string
	HashMethod                     string
	LastPasswordChangeTimestamp    string
	Logging                        string
	MaxSession                     string
	Password                       string
	PromptInheritedFrom            string
	PromptString                   string
	Timeout                        string
	TimeoutKind                    string
	Username                       string
}{
	AllowedManagementInterface:     "allowedmanagementinterface",
	AllowedManagementInterfaceKind: "allowedmanagementinterfacekind",
	Count:                          "__count",
	Encrypted:                      "encrypted",
	ExternalAuth:                   "externalauth",
	HashMethod:                     "hashmethod",
	LastPasswordChangeTimestamp:    "lastpwdchangetimestamp",
	Logging:                        "logging",
	MaxSession:                     "maxsession",
	Password:                       "password",
	PromptInheritedFrom:            "promptinheritedfrom",
	PromptString:                   "promptstring",
	Timeout:                        "timeout",
	TimeoutKind:                    "timeoutkind",
	Username:                       "username",
}

type SystemUser struct {
	AllowedManagementInterface     []string `json:"allowedmanagementinterface,omitempty" nitro:"permission=readwrite"`
	AllowedManagementInterfaceKind string   `json:"allowedmanagementinterfacekind,omitempty" nitro:"permission=readonly"`
	Count                          float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Encrypted                      string   `json:"encrypted,omitempty" nitro:"permission=readonly"`
	ExternalAuth                   string   `json:"externalauth,omitempty" nitro:"permission=readwrite"`
	HashMethod                     string   `json:"hashmethod,omitempty" nitro:"permission=readonly"`
	LastPasswordChangeTimestamp    float64  `json:"lastpwdchangetimestamp,omitempty" nitro:"permission=readonly"`
	Logging                        string   `json:"logging,omitempty" nitro:"permission=readwrite"`
	MaxSession                     float64  `json:"maxsession,omitempty" nitro:"permission=readwrite"`
	Password                       string   `json:"password,omitempty" nitro:"permission=readwrite"`
	PromptInheritedFrom            string   `json:"promptinheritedfrom,omitempty" nitro:"permission=readonly"`
	PromptString                   string   `json:"promptstring,omitempty" nitro:"permission=readwrite"`
	Timeout                        float64  `json:"timeout,omitempty" nitro:"permission=readwrite"`
	TimeoutKind                    string   `json:"timeoutkind,omitempty" nitro:"permission=readonly"`
	Username                       string   `json:"username,omitempty" nitro:"permission=readwrite"`
}

func (r SystemUser) GetTypeName() string {
	return "systemuser"
}

func NewSystemUserAddRequest(username string, password string) SystemUser {
	return SystemUser{
		Username:     username,
		Password:     password,
		ExternalAuth: "DISABLED",
	}
}
