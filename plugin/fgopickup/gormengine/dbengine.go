package gormengine

import (
	"github.com/jinzhu/gorm"
	"os"
)

var _dbEngine *Orm

type Orm struct {
	*gorm.DB
}

func GetOrmEngine() *Orm {
	return _dbEngine
}

func Initialize(dbpath string) *gorm.DB {
	var err error
	if _, err = os.Stat(dbpath); err != nil || os.IsNotExist(err) {
		f, err := os.Create(dbpath)
		if err != nil {
			return nil
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
	}
	gdb, err := gorm.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	//gdb.AutoMigrate(&model.Pickup{}, &model.PickupServant{}, &model.Servant{})
	orm := new(Orm)
	orm.DB = gdb
	_dbEngine = orm
	return gdb
}
