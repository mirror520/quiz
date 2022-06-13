package model

var Config = struct {
	DB struct {
		Host     string
		Port     int
		Username string
		Password string
		DBName   string
	}
}{}
