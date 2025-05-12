package mapping

type Config struct {
	DB DBConfig `mapstructure:"db"`
}

type DBConfig struct {
}
