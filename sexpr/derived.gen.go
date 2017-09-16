// Code generated by goderive DO NOT EDIT.

package sexpr

// deriveTuple returns a function, which returns the input values.
// Since tuples are not first class citizens in Go, this is a way to fake it, because functions that return tuples are first class citizens.
func deriveTuple(v0 string, v1 string) func() (string, string) {
	return func() (string, string) {
		return v0, v1
	}
}
