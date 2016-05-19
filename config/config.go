package config

// Config structure so we can access all attributes
type Config struct {
	DbString  string `json:"db_connection_string"`
	DbEngine  string `json:"db_engine"`
	SecretKey string `json:"secret_key"`
}
