package config

type Config struct {
    DB_String string `json:"db_connection_string"`
    DB_Engine string `json:"db_engine"`
    SESS_Secret string `json:"secret_key"`
}
