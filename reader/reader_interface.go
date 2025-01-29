package reader

type Reader interface {
	Read(location string) (any, error)
}
