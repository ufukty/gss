package gss

func silent[T any](t T, _ error) T {
	return t
}
