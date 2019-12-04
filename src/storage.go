package sdproject

type Storage struct {
	Data map[int64]string
}

func NewStorage() *Storage {
	return &Storage{
		Data: make(map[int64]string),
	}
}
