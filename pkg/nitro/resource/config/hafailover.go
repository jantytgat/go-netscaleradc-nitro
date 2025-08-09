package config

var HaFailoverFieldNames = struct {
	Force string
}{
	Force: "force",
}

type HaFailover struct {
	Force bool `json:"force,omitempty" nitro:"permission=readwrite"`
}

func (r HaFailover) GetTypeName() string {
	return "hafailover"
}
