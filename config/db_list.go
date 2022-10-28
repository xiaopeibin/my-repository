package config

type GeneralDB struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址
	Port         string `mapstructure:"port" json:"port" yaml:"port"` //:端口
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 密码
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 数据库连接配置
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
}
