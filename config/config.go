package config

// Config structure so we can access all attributes
type Config struct {
    DatabaseString string `json:"db_connection_string"`
    DatabaseEngine string `json:"db_engine"`
    SecretKey      string `json:"secret_key"`
}
