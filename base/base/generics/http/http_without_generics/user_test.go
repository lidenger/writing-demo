package http_without_generics

import "testing"

func TestGetAccount(t *testing.T) {
	resp, err := GetAccount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetOrder(t *testing.T) {
	resp, err := GetOrder()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestGetGoods(t *testing.T) {
	resp, err := GetGoods()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
