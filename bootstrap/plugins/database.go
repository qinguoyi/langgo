package plugins

import (
	"StorageProxy/app/models"
	"StorageProxy/bootstrap"
	"StorageProxy/config"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var lgDB = new(LangGoDB)

// LangGoDB 自定义DB结构
type LangGoDB struct {
	Once *sync.Once
	DB   *gorm.DB
}

func newLangGoDB() *LangGoDB {
	return &LangGoDB{
		DB:   &gorm.DB{},
		Once: &sync.Once{},
	}
}

// NewDB .
func (lg *LangGoDB) NewDB() *gorm.DB {
	if lgDB.DB != nil {
		return lgDB.DB
	} else {
		return lg.New().(*gorm.DB)
	}
}

// Name .
func (lg *LangGoDB) Name() string {
	return "DB"
}

// New 初始化DB
func (lg *LangGoDB) New() interface{} {
	lgDB = newLangGoDB()
	lgDB.initializeDB(bootstrap.NewConfig())
	return lg.DB
}

// Health .
func (lg *LangGoDB) Health() {
	tx := lgDB.DB.Exec("select now();")

	if tx.Error != nil {
		bootstrap.NewLogger().Logger.Error("db connect failed,", zap.Any("err", tx.Error))
	}
}

// Close .
func (lg *LangGoDB) Close() {}

func init() {
	p := &LangGoDB{}
	RegisteredPlugin(p)
}

func (lg *LangGoDB) initializeDB(conf *config.Configuration) {
	lg.Once.Do(func() {
		switch conf.Database.Driver {
		case "mysql":
			initMySqlGorm(conf)
		case "postgres":
			initPGGorm(conf)
		default:
			initMySqlGorm(conf)
		}
	})
}

func initPGGorm(conf *config.Configuration) {
	dbConfig := conf.Database

	if dbConfig.Database == "" {
		lgDB.DB = nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host,
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Database,
		strconv.Itoa(dbConfig.Port),
	)

	gormConfig := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if conf.Database.EnableLgLog {
		gormConfig = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,                // 禁用自动创建外键约束
			Logger:                                   getGormLogger(conf), // 使用自定义 Logger
		}
	}

	if db, err := gorm.Open(postgres.Open(dsn), gormConfig); err != nil {
		bootstrap.NewLogger().Logger.Error("mysql connect failed, err:", zap.Any("err", err))
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		// 执行数据库脚本建表
		//initMySqlTables(db)
		lgDB.DB = db
	}
}

// 初始化 mysql gorm.DB
func initMySqlGorm(conf *config.Configuration) {
	dbConfig := conf.Database

	if dbConfig.Database == "" {
		lgDB.DB = nil
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		strconv.Itoa(dbConfig.Port),
		dbConfig.Database,
		dbConfig.Charset,
	)

	gormConfig := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	if conf.Database.EnableLgLog {
		gormConfig = &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,                // 禁用自动创建外键约束
			Logger:                                   getGormLogger(conf), // 使用自定义 Logger
		}
	}

	if db, err := gorm.Open(mysql.Open(dsn), gormConfig); err != nil {
		bootstrap.NewLogger().Logger.Error("mysql connect failed, err:", zap.Any("err", err))
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		// 执行数据库脚本建表
		//initMySqlTables(db)
		lgDB.DB = db
	}
}

func getGormLogger(conf *config.Configuration) logger.Interface {
	var logMode logger.LogLevel

	switch conf.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(conf), logger.Config{
		SlowThreshold:             200 * time.Millisecond,             // 慢 SQL 阈值
		LogLevel:                  logMode,                            // 日志级别
		IgnoreRecordNotFoundError: false,                              // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  !conf.Database.EnableFileLogWriter, // 禁用彩色打印
	})
}

// 自定义 接管gorm日志，打印到文件 or 控制台
func getGormLogWriter(conf *config.Configuration) logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if conf.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   conf.Log.RootDir + "/" + conf.Database.LogFilename,
			MaxSize:    conf.Log.MaxSize,
			MaxBackups: conf.Log.MaxBackups,
			MaxAge:     conf.Log.MaxAge,
			Compress:   conf.Log.Compress,
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		bootstrap.NewLogger().Logger.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0) //这里为什么是0？
	}
}
