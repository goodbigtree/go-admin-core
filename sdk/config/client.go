package config

type Client struct {
	TboAppKey       string
	TboAppSecret    string
	TboSessionKey   string
	vipAppKey       string
	vipAppSecret    string
	vipSessionKey   string
	SnoAppKey       string
	SnoAppSecret    string
	SnoSessionKey   string
	PddAppKey       string
	PddAppSecret    string
	PddSessionKey   string
	JdoAppKey       string
	JdoAppSecret    string
	JdoSessionKey   string
	AlscoAppKey     string
	AlscoAppSecret  string
	AlscoSessionKey string
}

var ClientConfig = new(Client)
