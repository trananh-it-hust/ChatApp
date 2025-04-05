package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"log"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Jwt    JwtSetting    `mapstructure:"jwt"`
}

type MySQLSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	MaxIdle  int    `mapstructure:"maxIdConns"`
	MaxOpen  int    `mapstructure:"maxOpenConns"`
	MaxLife  int    `mapstructure:"connectMaxLifetime"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type LoggerSetting struct {
	Level         string `mapstructure:"log_level"`
	LogFile       string `mapstructure:"log_file"`
	LogMaxSize    int    `mapstructure:"log_max_size"`
	LogMaxBackups int    `mapstructure:"log_max_backups"`
	LogMaxAge     int    `mapstructure:"log_max_age"`
	LogCompress   bool   `mapstructure:"log_compress"`
}

type JwtSetting struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
