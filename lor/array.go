package lor

type FindOption string

const (
	FindFirst FindOption = "first"
	FindLast  FindOption = "last"
)

func Find[T any](arr []T, option FindOption) *T {
	if len(arr) == 0 {
		return nil
	}

	if option == FindFirst {
		return &arr[0]
	}

	return &arr[len(arr)-1]
}

func FindWithCondition[T any](arr []T, f func(T) bool) *T {
	if len(arr) == 0 {
		return nil
	}

	for _, v := range arr {
		if f(v) {
			return &v
		}
	}

	return nil
}
