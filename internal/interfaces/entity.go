package interfaces

import "context"

type EntityRepository[T any] interface {
	Create(c context.Context, object *T) error
	Get(c context.Context, object *[]T) error
	GetById(c context.Context, id int, object *T) error
	Update(c context.Context, object *T) error
	Delete(c context.Context, id int) error
}

type EntityService[T any] interface {
	Create(c context.Context, name string) (*T, error)
	Get(c context.Context) *[]T
	GetById(c context.Context, id int) (*T, error)
	Update(c context.Context, id int, name string) (*T, error)
	Delete(c context.Context, id int) error
}

// знаю,ты будешь недоволен моей редакцией,но ты оч зря поменял сущность репозиториев и сервисов на шаблоны выше. мне потребовался свой. Надеюсь ты внимательно изучишь весь маршрут
type EntityServiceUser[T any] interface {
	Create(c context.Context, object *T, password string) (*T, error)
	Get(c context.Context) *[]T
	GetById(c context.Context, id int) (*T, error)
	Update(c context.Context, id int, name string) (*T, error)
	Delete(c context.Context, id int) error
}
