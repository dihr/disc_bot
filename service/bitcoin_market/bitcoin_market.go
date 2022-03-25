package bitcoin_market

import (
	"disc_bot/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type (
	BitcoinMarketSvc interface {
		GetCoins(coin string) map[string]string
		GetCoinTicker(coin string) (model.Ticker, error)
		GetOrderBook(coin string) (model.OrderBook, error)
	}

	bitcoinMarketImp struct {
		baseURL string
	}
)

func NewBitcoinMarket(baseURL string) BitcoinMarketSvc {
	return &bitcoinMarketImp{
		baseURL: baseURL,
	}
}

func (b *bitcoinMarketImp) GetOrderBook(coin string) (model.OrderBook, error) {
	client := http.Client{
		Timeout: 20 * time.Second,
	}
	url := fmt.Sprintf("%s/%s/orderbook", b.baseURL, coin)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.OrderBook{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return model.OrderBook{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return model.OrderBook{}, errors.New(fmt.Sprintf("fail to get ticker data [status_code %d]",
			resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.OrderBook{}, err
	}

	orderBook := model.OrderBook{}
	if err := json.Unmarshal(body, &orderBook); err != nil {
		fmt.Println(err.Error())
		return model.OrderBook{}, err
	}
	return orderBook, nil
}

func (b *bitcoinMarketImp) GetCoinTicker(coin string) (model.Ticker, error) {
	client := http.Client{
		Timeout: 20 * time.Second,
	}
	url := fmt.Sprintf("%s/%s/ticker", b.baseURL, coin)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Ticker{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return model.Ticker{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return model.Ticker{}, errors.New(fmt.Sprintf("fail to get ticker data [status_code %d]",
			resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.Ticker{}, err
	}

	tickerW := model.TickerWrapper{}
	if err := json.Unmarshal(body, &tickerW); err != nil {
		fmt.Println(err.Error())
		return model.Ticker{}, err
	}
	return tickerW.Ticker, nil
}

func (b *bitcoinMarketImp) GetCoins(coin string) map[string]string {
	result := make(map[string]string)
	for key, value := range coins {
		parameter := fmt.Sprintf("^(?i)%s", coin)
		if ok, _ := regexp.MatchString(parameter, key); ok {
			result[key] = value
		}
	}
	return result
}
