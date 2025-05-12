package configs

import (
	"os"

	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/brunobotter/casa-codigo/internal/domain/service"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

type Deps struct {
	Cfg *mapping.Config
	DB  *gorm.DB
	Svc contract.ServiceManager
}

func Init() *Deps {
	logger.Info("loading configs ")
	profile := os.Getenv("PROFILE")
	cfg, err := read(profile)
	if err != nil {
		logger.Errorf("Error to read config %v", err)
	}

	deps := &Deps{
		Cfg: cfg,
	}
	return deps
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetMySql() *gorm.DB {
	return db
}

func (deps *Deps) ConfigDB() *Deps {
	logger := GetLogger("mysql")
	dsn := "root:171191@tcp(127.0.0.1:3306)/casa-codigo?charset=utf8mb4&parseTime=true&loc=Local"
	//create db and connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("mysql carrinho error: %v", err)
		return nil
	}
	//migrate scheama
	/*err = db.AutoMigrate(&scheamas.Carrinho{})
	if err != nil {
		logger.Errorf("mysql automigration error: %v", err)
		return nil, err
	}*/
	deps.DB = db
	return deps
}

func (deps *Deps) ConfigService() *Deps {
	if deps.DB == nil {
		logger.Errorf("Db was not initialize")

	}

	svc, err := service.New(service.ServiceDeps{
		Cfg: deps.Cfg,
		DB:  deps.DB,
	})
	if err != nil {

	}
	deps.Svc = svc
	return deps
}

func (deps *Deps) ConfigAll() *Deps {
	return deps.ConfigDB().ConfigService()
}

func setupConfig(profile string) {
	viper.SetEnvPrefix("api")
	viper.AutomaticEnv()
	viper.SetConfigName(profile)
	viper.AddConfigPath(".")
}

func read(profile string) (*mapping.Config, error) {
	setupConfig(profile)
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Error to read config: %v", err)
		return nil, err
	}

	config := mapping.Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
