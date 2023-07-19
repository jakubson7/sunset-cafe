package storage

type LocalStorage struct {
	dir  string
	addr string
}

func NewLocalStorage(dir string, addr string) *LocalStorage {
	return &LocalStorage{
		dir:  dir,
		addr: addr,
	}
}

func (s *LocalStorage) GetID() string {
	return "LOCAL"
}
func (s *LocalStorage) NewImage(src []byte) (*ImageURL, error) {
	return &ImageURL{}, nil
}

func (s *LocalStorage) Setup() error {
	return nil
}
