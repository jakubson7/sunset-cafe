package models

type Model interface {
	Validate() error
	JSON() (string, error)
}
