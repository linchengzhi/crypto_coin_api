package huobi

import (
	"crypto_coin_api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var api = New(http.DefaultClient, "b51cb101-c24aa919-22ae1ca0-0dfe6", "7c7be766-20c2aa9b-0b56dc5e-d57f5")

func TestHuobi_GetTicker(t *testing.T) {
	ticker, err := api.GetTicker(coinapi.BTC_CNY)
	assert.Empty(t, err)
	t.Log(ticker)
}

func TestHuobi_GetDepth(t *testing.T) {
	depth, err := api.GetDepth(4, coinapi.BTC_CNY)
	assert.Empty(t, err)
	t.Log(depth)
}

func TestHuobi_GetAccount(t *testing.T) {
	depth, err := api.GetAccount()
	assert.Empty(t, err)
	t.Log(depth)
}
