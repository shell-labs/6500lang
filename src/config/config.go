package config

type Config interface {
	Load()
	Save()
}
