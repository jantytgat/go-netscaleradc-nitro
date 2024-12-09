package config

import "log/slog"

type Login struct {
	Username string `json:"username,omitempty" nitro:"permission=readwrite"`
	Password string `json:"password,omitempty" nitro:"permission=readwrite"`
}

func (r Login) GetTypeName() string {
	return "login"
}

func (r Login) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("username", r.Username))
}
