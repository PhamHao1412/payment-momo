package app

type Config struct {
	PartnerCode string `mapstructure:"PARTNER_CODE"`
	AccessKey   string `mapstructure:"ACCESS_KEY"`
	SecretKey   string `mapstructure:"SECRET_KEY"`
	AppPort     string `mapstructure:"APP_PORT"`
	AppBaseURL  string `mapstructure:"APP_BASE_URL"`
	Lang        string `mapstructure:"LANG"`
}
