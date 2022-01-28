package database

type Database interface {
	SetDBName(dbName string)
	DBName() string
	Insert(collection string, data interface{}) error
	Count(collection string, query interface{}) (int, error)
	FindOne(collection string, query interface{}, data interface{}) error
	FindOnePrimary(collection string, query interface{}, data interface{}) error
	Find(collection string, query interface{}, data interface{}) error
	FindAndSelect(collection string, query interface{}, selector interface{}, data interface{}) error
	FindLimit(collection string, query interface{}, limit int, data interface{}) error
	FindLimitAndOrderBy(collection string, query interface{}, limit int, orderBy string, data interface{}) error
	FindLimitAndOrderByM(collection string, query interface{}, limit int, orderBy []string, data interface{}) error
	FindOrderByOne(collection string, query interface{}, orderby string, data interface{}) error
	FindByOrder(collection string, query interface{}, orderBy []string, data interface{}) error
	FindOrderAndUpdate(collection string, query interface{}, orderBy []string, changeQuery interface{}, afterItem interface{}) error
	Upsert(collection string, query interface{}, data interface{}) error
	Update(collection string, query interface{}, data interface{}) error
	UpdateAll(collection string, query interface{}, data interface{}) error
	Aggregate(collection string, query interface{}, data interface{}) error
	RemoveAll(collection string, query interface{}) error
}

type Mongo struct {
	db string
}

func Insert(db Database, collection string, data interface{}) error {
	return db.Insert(collection, data)
}
