package nitro

import "log/slog"

type Credentials struct {
	Username string
	Password string
}

func (c Credentials) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("username", c.Username))
}
