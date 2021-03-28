package flatten

func ApplyFlatten(input []interface{}) (r []interface{}, d int) {
	r = flatten(input)
	d = depth(input)
	return
}

func flatten(input []interface{}) (r []interface{}) {
	for _, e := range input {
		switch i := e.(type) {
		case []interface{}:
			r = append(r, flatten(i)...)
		case interface{}:
			r = append(r, i)
		}
	}
	return
}

func depth(input []interface{}) int {
	maxdepth := -1
	for _, e := range input {
		switch i := e.(type) {
		case []interface{}:
			de := depth(i)
			if de > maxdepth {
				maxdepth = de
			}
		}
	}

	return maxdepth + 1
}
