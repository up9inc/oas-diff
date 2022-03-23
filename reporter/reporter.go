package reporter

type Reporter interface {
	Build() ([]byte, error)
}
