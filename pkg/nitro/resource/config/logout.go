package config

type Logout struct{}

func (r Logout) GetTypeName() string {
	return "logout"
}
