package lor

type fnc[R any] func() R

func FunctionIf[R any](condition bool, trueFunc fnc[R], falseFunc fnc[R]) R {
	if condition {
		return trueFunc()
	}

	return falseFunc()
}

func FunctionDo[R any](condition bool, doFunc fnc[R]) R {
	var r R
	if !condition {
		return r
	}

	return doFunc()
}
