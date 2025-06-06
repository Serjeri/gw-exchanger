package models

type ExchangeRates struct {
	USDRUB int `json:"usd_rub"`
	USDEUR int `json:"usd_eur"`
	EURRUB int `json:"eur_rub"`
	EURUSD int `json:"eur_usd"`
	RUBEUR int `json:"rub_eur"`
	RUBUSD int `json:"rub_usd"`
}
