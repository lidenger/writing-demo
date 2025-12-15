package ptr_with_generics

import (
	"generics/models"
	"testing"
)

func Test1(t *testing.T) {
	p := Ptr(2)
	t.Log(p, *p)

	p2 := Ptr("123abc")
	t.Log(p2, *p2)
}

func Test2(t *testing.T) {
	cond := models.AccountCond{
		IsEnable: Ptr(true),
		Gender:   Ptr(1),
	}
	t.Log(cond)
}
