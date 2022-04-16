package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/akromibn37/hospitalityCollaboration/pkg/setting"
)

var db *sql.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// Setup initializes the database instance
func Setup() {
	var err error
	conn := setting.DatabaseSetting.User + ":" + setting.DatabaseSetting.Password + "@tcp(" + setting.DatabaseSetting.Host + ")/" + setting.DatabaseSetting.Name
	db, err = sql.Open(setting.DatabaseSetting.Type, conn)
	// db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	setting.DatabaseSetting.User,
	// 	setting.DatabaseSetting.Password,
	// 	setting.DatabaseSetting.Host,
	// 	setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("database.Setup err: %v", err)
	}

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return setting.DatabaseSetting.TablePrefix + defaultTableName
	// }

	// db.SingularTable(true)
	// db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatalf("No DB Connected")
	}
	return db
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
