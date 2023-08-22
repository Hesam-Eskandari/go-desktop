package observer

type ISubscriber[T any] interface {
	Update(report T)
	GetId() string
}
