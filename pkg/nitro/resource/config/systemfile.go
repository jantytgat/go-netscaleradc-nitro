package config

import "encoding/base64"

var SystemFileFieldNames = struct {
	AccessTime   string
	Content      string
	Count        string
	Encoding     string
	Filename     string
	Location     string
	Mode         string
	ModifiedTime string
	Size         string
}{
	AccessTime:   "fileaccesstime",
	Content:      "filecontent",
	Count:        "__count",
	Encoding:     "fileencoding",
	Filename:     "filename",
	Location:     "filelocation",
	Mode:         "filemode",
	ModifiedTime: "filemodifiedtime",
	Size:         "filesize",
}

type SystemFile struct {
	AccessTime   string   `json:"fileaccesstime,omitempty" nitro:"permission=readonly"`
	Content      string   `json:"filecontent,omitempty" nitro:"permission=readwrite"`
	Count        float64  `json:"__count,omitempty" nitro:"permission=readonly"`
	Encoding     string   `json:"fileencoding,omitempty" nitro:"permission=readwrite"`
	Filename     string   `json:"filename,omitempty" nitro:"permission=readwrite"`
	Location     string   `json:"filelocation,omitempty" nitro:"permission=readwrite"`
	Mode         []string `json:"filemode,omitempty" nitro:"permission=readonly"`
	ModifiedTime string   `json:"filemodifiedtime,omitempty" nitro:"permission=readonly"`
	Size         string   `json:"filesize,omitempty" nitro:"permission=readonly"`
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
