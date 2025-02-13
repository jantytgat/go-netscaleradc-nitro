package config

import (
	"fmt"
	"regexp"
)

const (
	validVersion = `^NetScaler\sNS(\d+.[0|1]):\sBuild\s(\d+\.\d+)\.nc.*Date:\s([a-zA-Z0-9:\s]*,\s[0-9:]*).*$`
)

var reVersion = regexp.MustCompile(validVersion)

type NsVersion struct {
	InstalledVersion bool   `json:"installedversion,omitempty" nitro:"permission=readwrite"`
	Version          string `json:"version,omitempty" nitro:"permission=readonly"`
	Mode             string `json:"mode,omitempty" nitro:"permission=readonly"`
}

func (r NsVersion) GetTypeName() string {
	return "nsversion"
}

func (r NsVersion) Details() (NsVersionDetail, error) {
	if !reVersion.MatchString(r.Version) {
		return NsVersionDetail{}, fmt.Errorf("invalid version: %s", r.Version)
	}

	parts := reVersion.FindStringSubmatch(r.Version)
	return NsVersionDetail{
		Version: parts[1],
		Build:   parts[2],
		Date:    parts[3],
	}, nil
}

type NsVersionDetail struct {
	Version string
	Build   string
	Date    string
}
