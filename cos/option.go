package cos

type Option struct {
	SecretID  string `mapstructure:"secret_id" json:"secret_id"`
	SecretKey string `mapstructure:"secret_key" json:"secret_key"`
	Region    string `mapstructure:"region" json:"region"`
	Domain    string `mapstructure:"domain" json:"domain"`
	Bucket    string `mapstructure:"bucket" json:"bucket"`
}
