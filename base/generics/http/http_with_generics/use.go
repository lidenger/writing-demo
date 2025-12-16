package http_with_generics

import "generics/models"

func GetAccount() ([]models.Account, error) {
	url := "http://127.0.0.1:6666/api/v1/account"
	resp, err := HttpGet[models.Resp[[]models.Account]](url)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func GetOrder() ([]models.Order, error) {
	url := "http://127.0.0.1:6666/api/v1/order"
	resp, err := HttpGet[models.Resp[[]models.Order]](url)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func GetGoods() (models.Goods, error) {
	url := "http://127.0.0.1:6666/api/v1/goods"
	resp, err := HttpGet[models.Resp[models.Goods]](url)
	if err != nil {
		return models.Goods{}, err
	}
	return resp.Data, nil
}
