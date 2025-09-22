package models

type Config struct {
	AppParam      AppParams      `json:"app_params"`
	PostgresParam PostgresParams `json:"postgres_params"`
}

type AppParams struct {
	GinMode    string `json:"gin_mode"`
	Port       string `json:"port"`
	ServerUrl  string `json:"server_url"`
	ServerName string `json:"server_name"`
}
type PostgresParams struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
}
