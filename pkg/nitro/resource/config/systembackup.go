package config

type SystemBackup struct {
	Filename         string  `json:"filename,omitempty" nitro:"permission=readwrite"`
	Level            string  `json:"level,omitempty" nitro:"permission=readwrite"`
	IncludeKernel    string  `json:"includekernel,omitempty" nitro:"permission=readwrite"`
	Comment          string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Size             float64 `json:"size,omitempty" nitro:"permission=readonly"`
	UseLocalTimezone bool    `json:"uselocaltimezone,omitempty" nitro:"permission=readonly"`
	CreationTime     string  `json:"creationtime,omitempty" nitro:"permission=readonly"`
	Version          string  `json:"version,omitempty" nitro:"permission=readonly"`
	CreatedBy        string  `json:"createdby,omitempty" nitro:"permission=readonly"`
	IpAddress        string  `json:"ipaddress,omitempty" nitro:"permission=readonly"`
	SkipBackup       bool    `json:"skipbackup,omitempty" nitro:"permission=readwrite"`
	Count            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemBackup) GetTypeName() string {
	return "systembackup"
}

func NewSystemBackupCreateRequest(name string, level string) SystemBackup {
	// TODO data validation
	return SystemBackup{
		Filename: name,
		Level:    level,
	}
}
