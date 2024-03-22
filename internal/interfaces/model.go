package interfaces

type Model interface {
	JSON() (string, error)
}
