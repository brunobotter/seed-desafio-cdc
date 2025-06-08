package datasql

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/brunobotter/casa-codigo/configs/mapping"
	"github.com/brunobotter/casa-codigo/internal/domain/contract"
	"github.com/go-sql-driver/mysql"
)

var (
	instance     *Conn
	dbInstance   *sql.DB
	onceOB       sync.Once
	onceInstance sync.Once
	connErr      error
)

type Conn struct {
	db       *sql.DB
	author   *authorRepository
	category *categoryRepository
	book     *bookRepository
	country  *countryRepository
	state    *stateRepository
	custumer *customerRepository
	payment  *paymentRepository
	coupon   *couponRepository
}

func Instance(cfg *mapping.Config) (contract.DataManager, error) {
	onceInstance.Do(func() {
		db, err := GetDB(cfg)
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}

		instance = &Conn{db: db}
		instance.author = &authorRepository{db, instance}
		instance.category = &categoryRepository{db, instance}
		instance.book = &bookRepository{db, instance}
		instance.country = &countryRepository{db, instance}
		instance.state = &stateRepository{db, instance}
		instance.custumer = &customerRepository{db, instance}
		instance.payment = &paymentRepository{db, instance}
		instance.coupon = &couponRepository{db, instance}
	})
	return instance, connErr
}

func GetDB(cfg *mapping.Config) (*sql.DB, error) {
	onceOB.Do(func() {
		mysqlCfg := getMySqlConfig(cfg)
		db, err := sql.Open("mysql", mysqlCfg.FormatDSN())
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}
		maxLifeTimeInMinutes, _ := time.ParseDuration(fmt.Sprintf("%vmin", cfg.DB.MaxLifeTimeInMinutes))
		db.SetConnMaxIdleTime(maxLifeTimeInMinutes)
		db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
		db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
		err = db.Ping()
		if err != nil {
			connErr = errors.New(err.Error())
			return
		}
		dbInstance = db
	})
	return dbInstance, connErr
}

func getMySqlConfig(cfg *mapping.Config) *mysql.Config {
	mysqlConfig := mysql.Config{
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		DBName:               cfg.DB.Name,
		User:                 cfg.DB.User,
		Passwd:               cfg.DB.Pass,
		ParseTime:            true,
		AllowNativePasswords: true,
		Params:               cfg.DB.Params,
	}
	return &mysqlConfig
}

func (c *Conn) AuthorRepo() contract.AuthorRepository {
	return c.author
}

func (c *Conn) CategoryRepo() contract.CategoryRepository {
	return c.category
}

func (c *Conn) BookRepo() contract.BookRepository {
	return c.book
}

func (c *Conn) CountryRepo() contract.CountryRepository {
	return c.country
}

func (c *Conn) StateRepo() contract.StateRepository {
	return c.state
}

func (c *Conn) CustomerRepo() contract.CustomerRepository {
	return c.custumer
}

func (c *Conn) PaymentRepo() contract.PaymentRepository {
	return c.payment
}

func (c *Conn) CouponRepo() contract.CouponRepository {
	return c.coupon
}
