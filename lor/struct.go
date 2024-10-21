package lor

import "encoding/json"

func StructToStruct[T, U any](t *T) (*U, error) {
	var u U
	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func MapStructs[T any, U []any](t T, arr U) (U, error) {
	bt, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	res := make(U, 0)

	for _, v := range arr {
		err := json.Unmarshal(bt, &v)
		if err != nil {
			return nil, err
		}

		res = append(res, v)
	}

	return res, nil
}
