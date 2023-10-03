package global

import (
	go_grc "github.com/craftzbay/go_grc/v2"
)

type Config struct {
	// system
	Port        string `mapstructure:"PORT"`
	SystemCode  string `mapstructure:"SYSTEM_CODE"`
	SystemToken string `mapstructure:"SYSTEM_TOKEN"`
	// db
	DBHost            string `mapstructure:"DB_HOST"`
	DBUserName        string `mapstructure:"DB_USERNAME"`
	DBUserPassword    string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBSchema          string `mapstructure:"DB_SCHEMA"`
	DBTablePrefix     string `mapstructure:"DB_TABLE_PREFIX"`
	DBMaxIdleConn     int    `mapstructure:"DB_MAX_IDLE_CONN"`
	DBMaxOpenConn     int    `mapstructure:"DB_MAX_OPEN_CONN"`
	DBMaxConnLifetime int    `mapstructure:"DB_MAX_CONN_LIFETIME"`
	// file server (minio)
	MinioEndpoint  string `mapstructure:"MINIO_ENDPOINT"`
	MinioAccessKey string `mapstructure:"MINIO_ACCESSKEY"`
	MinioSecretKey string `mapstructure:"MINIO_SECRETKEY"`
	// core url
	UrlAccountBackend string `mapstructure:"URL_ACCOUNT_BACKEND"`
	UrlCoreBackend    string `mapstructure:"URL_CORE_BACKEND"`
	// jwt
	JwtSecret           string `mapstructure:"JWT_SECRET"`
	JwtSecretPubKeyPath string `mapstructure:"JWT_SECRET_PUBLIC_KEY_PATH"`
	JwtSecretPrvKeyPath string `mapstructure:"JWT_SECRET_PRIVATE_KEY_PATH"`
	// otp
	OtpWaitSecond   int    `mapstructure:"OTP_WAIT_SECOND"`
	UrlEmailService string `mapstructure:"URL_EMAIL_SERVICE"`
	// auth google
	GoogleClientId        string `mapstructure:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret    string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GoogleAuthCallbackUrl string `mapstructure:"GOOGLE_AUTH_CALLBACK_URL"`
	// auth fb
	FbClientId        string `mapstructure:"FB_CLIENT_ID"`
	FbClientSecret    string `mapstructure:"FB_CLIENT_SECRET"`
	FbAuthCallbackUrl string `mapstructure:"FB_AUTH_CALLBACK_URL"`
	// auth apple
	AppleAppClientId     string `mapstructure:"APPLE_APP_CLIENT_ID"`
	AppleWebClientId     string `mapstructure:"APPLE_WEB_CLIENT_ID"`
	AppleTeamId          string `mapstructure:"APPLE_TEAM_ID"`
	AppleClientSecret    string `mapstructure:"APPLE_CLIENT_SECRET"`
	AppleKeyId           string `mapstructure:"APPLE_KEY_ID"`
	AppleAuthCallbackUrl string `mapstructure:"APPLE_AUTH_CALLBACK_URL"`
}

var Conf *Config

func LoadConfig(path string) (err error) {

	if Conf, err = go_grc.LoadConfig("app", "env", path, Conf); err != nil {
		return err
	}

	return nil
}
