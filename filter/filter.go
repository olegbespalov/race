package filter

func New() func() bool {
	return func() bool {
		return false
	}
}
