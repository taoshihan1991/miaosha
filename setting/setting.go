package setting

type GlobalConfig struct {
	Redis    RedisConfig    `json:"redis"`
	Database DatabaseConfig `json:"database"`
}
type RedisConfig struct {
	Ip   string
	Port string
}
type DatabaseConfig struct {
}
