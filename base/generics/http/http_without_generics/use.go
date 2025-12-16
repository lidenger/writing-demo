package http_without_generics

import (
	"encoding/json"
	"generics/models"
)

func GetAccount() ([]models.Account, error) {
	url := "http://127.0.0.1:6666/api/v1/account"

	data, err := HttpGet(url)
	if err != nil {
		return nil, err
	}

	resp := models.AccountResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func GetOrder() ([]models.Order, error) {
	url := "http://127.0.0.1:6666/api/v1/order"

	data, err := HttpGet(url)
	if err != nil {
		return nil, err
	}

	resp := models.OrderResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func GetGoods() (models.Goods, error) {
	url := "http://127.0.0.1:6666/api/v1/goods"

	data, err := HttpGet(url)
	if err != nil {
		return models.Goods{}, err
	}

	resp := models.GoodsResp{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return models.Goods{}, err
	}

	return resp.Data, nil
}
