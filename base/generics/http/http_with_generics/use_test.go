package http_with_generics

import "testing"

func TestGetAccount(t *testing.T) {
	result, err := GetAccount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetOrder(t *testing.T) {
	result, err := GetOrder()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGetGoods(t *testing.T) {
	result, err := GetGoods()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
