package config

var NsSavedConfigFieldNames = struct {
	TextBlob string
}{
	TextBlob: "textblob",
}

type NsSavedConfig struct {
	TextBlob string `json:"textblob,omitempty" nitro:"permission=readonly"`
}

func (r NsSavedConfig) GetTypeName() string {
	return "nssavedconfig"
}
