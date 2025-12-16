package ptr_without_generics

import (
	"generics/models"
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

func Test2(t *testing.T) {
	isEnable := true
	gender := 1
	cond := models.AccountCond{
		IsEnable: &isEnable,
		Gender:   &gender,
	}
	t.Log(cond)
	if cond.IsEnable != nil {
		t.Log(*cond.IsEnable)
	}
	if cond.Gender != nil {
		t.Log(*cond.Gender)
	}
}

func Test3(t *testing.T) {
	cond := models.AccountCond{
		IsEnable: PtrBool(true),
		Gender:   PtrInt(1),
	}
	t.Log(cond)
}
