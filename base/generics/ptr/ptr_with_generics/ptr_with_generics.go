package ptr_with_generics

func Ptr[T any](v T) *T {
	return &v
}
