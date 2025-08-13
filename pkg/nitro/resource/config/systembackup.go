package config

var SystemBackupFieldNames = struct {
	Comment          string
	Count            string
	CreatedBy        string
	CreationTime     string
	Filename         string
	IncludeKernel    string
	IpAddress        string
	Level            string
	Size             string
	SkipBackup       string
	UseLocalTimezone string
	Version          string
}{
	Comment:          "comment",
	Count:            "__count",
	CreatedBy:        "createdby",
	CreationTime:     "creationtime",
	Filename:         "filename",
	IncludeKernel:    "includekernel",
	IpAddress:        "ipaddress",
	Level:            "level",
	Size:             "size",
	SkipBackup:       "skipbackup",
	UseLocalTimezone: "uselocaltimezone",
	Version:          "version",
}

type SystemBackup struct {
	Comment          string  `json:"comment,omitempty" nitro:"permission=readwrite"`
	Count            float64 `json:"__count,omitempty" nitro:"permission=readonly"`
	CreatedBy        string  `json:"createdby,omitempty" nitro:"permission=readonly"`
	CreationTime     string  `json:"creationtime,omitempty" nitro:"permission=readonly"`
	Filename         string  `json:"filename,omitempty" nitro:"permission=readwrite"`
	IncludeKernel    string  `json:"includekernel,omitempty" nitro:"permission=readwrite"`
	IpAddress        string  `json:"ipaddress,omitempty" nitro:"permission=readonly"`
	Level            string  `json:"level,omitempty" nitro:"permission=readwrite"`
	Size             float64 `json:"size,omitempty" nitro:"permission=readonly"`
	SkipBackup       bool    `json:"skipbackup,omitempty" nitro:"permission=readwrite"`
	UseLocalTimezone bool    `json:"uselocaltimezone,omitempty" nitro:"permission=readonly"`
	Version          string  `json:"version,omitempty" nitro:"permission=readonly"`
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
