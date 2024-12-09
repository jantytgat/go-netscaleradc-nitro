package config

import "encoding/base64"

type SystemFile struct {
	Filename     string   `json:"filename,omitempty" nitro:"permission=readwrite"`
	Content      string   `json:"filecontent,omitempty" nitro:"permission=readwrite"`
	Location     string   `json:"filelocation,omitempty" nitro:"permission=readwrite"`
	Encoding     string   `json:"fileencoding,omitempty" nitro:"permission=readwrite"`
	AccessTime   string   `json:"fileaccesstime,omitempty" nitro:"permission=readonly"`
	ModifiedTime string   `json:"filemodifiedtime,omitempty" nitro:"permission=readonly"`
	Mode         []string `json:"filemode,omitempty" nitro:"permission=readonly"`
	Size         string   `json:"filesize,omitempty" nitro:"permission=readonly"`
	Count        float64  `json:"__count,omitempty" nitro:"permission=readonly"`
}

func (r SystemFile) GetTypeName() string {
	return "systemfile"
}

func NewSystemFileAddRequest(fileName string, location string, fileContent []byte) SystemFile {
	return SystemFile{
		Filename: fileName,
		Location: location,
		Content:  base64.StdEncoding.EncodeToString(fileContent),
		Encoding: "BASE64",
	}
}
