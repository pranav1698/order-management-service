package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CheckDuplicateOrderTrue(t *testing.T) {
	order := Order{
		Id:     "abcdef-123456",
		Status: "PENDING_INVOICE",
		Items: []Item{
			Item{
				Id:          "123456",
				Description: "a product description",
				Price:       12.40,
				Quantity:    1,
			},
		},
		Total:        12.40,
		CurrencyUnit: "USD",
	}
	Orders = append(Orders, order)

	check := CheckDuplicateOrder(order)
	assert.True(t, check)
}

func Test_CheckDuplicateOrderFalse(t *testing.T) {
	order := Order{
		Id:     "abcdef-123456",
		Status: "PENDING_INVOICE",
		Items: []Item{
			Item{
				Id:          "123456",
				Description: "a product description",
				Price:       12.40,
				Quantity:    1,
			},
		},
		Total:        12.40,
		CurrencyUnit: "USD",
	}
	Orders = append(Orders, order)

	newOrder := Order{
		Id:     "abcdef-12345689",
		Status: "PENDING_INVOICE",
		Items: []Item{
			Item{
				Id:          "123456",
				Description: "a product description",
				Price:       12.40,
				Quantity:    1,
			},
		},
		Total:        12.40,
		CurrencyUnit: "USD",
	}

	check := CheckDuplicateOrder(newOrder)
	assert.False(t, check)
}

func Test_GetOrders(t *testing.T) {
	req, err := http.NewRequest("GET", "/orders", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOrders)

	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
}
