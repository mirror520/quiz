package model

var Config = struct {
	BaseURL string

	DB struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	}
}{}
