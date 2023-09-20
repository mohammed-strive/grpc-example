package interfaces

type StoreClient[T any] interface {
	GetItem(string) (T, error)
	CreateItem(T) (T, error)
	UpdateItem(string, T) (T, error)
	DeleteItem(string) error
}
