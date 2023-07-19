package storage

type StorageProvider interface {
	GetID() string
	NewImage(raw []byte) (*ImageURL, error)
	Setup() error
}

type ImageURL struct {
	small  string
	medium string
	big    string
}
