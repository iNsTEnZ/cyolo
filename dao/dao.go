package dao

type DAO interface {
	Put(key interface{}, value interface{}) error
	Get(key interface{}, defaultValue interface{}) (interface{}, error)
}
