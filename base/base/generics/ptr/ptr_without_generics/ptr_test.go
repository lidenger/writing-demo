package ptr_without_generics

import (
	"testing"
)

func Test1(t *testing.T) {
	p2 := Ptr(2)
	if p, ok := p2.(*int); ok {
		t.Log(p, *p)
	}

	p3 := PtrInt(3)
	t.Log(p3, *p3)
}
