package hdwallet_api

type configService interface {
	IsDev() bool
	IsDebug() bool
	IsLocal() bool

	GetBindPort() string
}