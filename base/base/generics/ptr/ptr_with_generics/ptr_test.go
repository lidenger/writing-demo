package ptr_with_generics

import (
	"testing"
)

func Test1(t *testing.T) {
	p := Ptr(2)
	t.Log(p, *p)

	p2 := Ptr("123abc")
	t.Log(p2, *p2)
}
