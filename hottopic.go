package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func AddToCart(client http.Client) {
	dataCreate := url.Values{
		"shippingMethod-13249991":  {"shipToHome"},
		"deliveryMsgHome-13249991": {"Backorder:Expected to ship by:05/18/21 - 05/29/21"},
		"atc-13249991":             {"0.0"},
		"storeId-13249991":         {"2536"},
		"deliveryType-13249991":    {""},
		"deliveryMsg-13249991":     {"Unavailable for in-store pickup"},
		"cgid":                     {""},
		"pid":                      {"13249991"},
		"Quantity":                 {"1"},
		"hasColorSelected":         {"notRequired"},
		"hasSizeSelected":          {"true"},
		"hasInseamSelected":        {"notRequired"},
		"cartAction":               {"add"},
		"productColor":             {""},
	}
	data := strings.NewReader(dataCreate.Encode())

	//Create HTTP POST request
	req, err := http.NewRequest("POST", "https://www.hottopic.com/on/demandware.store/Sites-hottopic-Site/default/Cart-AddProduct?format=ajax", data)
	handleErr(err)

	//Add Standard headers + Cookie
	_cookies := []string{""}
	req.Header = addStandardHeaders(Header{cookie: _cookies, content: data})

	//POST
	resp, err := client.Do(req)
	handleErr(err)

	//cleanup
	defer resp.Body.Close()
	defer req.Body.Close()
}
func GetCheckout(client http.Client) string {

	//Create HTTP POST request
	req, err := http.NewRequest("GET", "https://www.hottopic.com/cart", nil)
	handleErr(err)

	//Add Standard headers + Cookie
	_cookies := []string{""}
	req.Header = addStandardHeaders(Header{cookie: _cookies})

	//POST
	resp, err := client.Do(req)
	handleErr(err)

	//cleanup
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
func ProceedToCheckout(client http.Client, dwcont string) string {
	// /https://www.hottopic.com/cart?dwcont=C276350184
	//payload
	//dwfrm_cart_checkoutCart: Checkout
	dataCreate := url.Values{
		"dwfrm_cart_checkoutCart": {"checkout"},
	}
	data := strings.NewReader(dataCreate.Encode())

	//Create HTTP POST request
	req, err := http.NewRequest("POST", "https://www.hottopic.com/cart?dwcont="+dwcont, data)
	handleErr(err)

	//Add Standard headers + Cookie
	_cookies := []string{""}
	req.Header = addStandardHeaders(Header{cookie: _cookies, content: data})

	//POST
	resp, err := client.Do(req)
	handleErr(err)

	//cleanup
	defer resp.Body.Close()
	defer req.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
func GuestCheckout(client http.Client, dwcont string, secureKey string) string {
	// /https://www.hottopic.com/cart?dwcont=C276350184
	//payload
	//dwfrm_cart_checkoutCart: Checkout
	dataCreate := url.Values{
		"dwfrm_login_unregistered": {"Checkout As a Guest"},
		"dwfrm_login_securekey":    {secureKey},
	}
	data := strings.NewReader(dataCreate.Encode())

	//Create HTTP POST request
	req, err := http.NewRequest("POST", "https://www.hottopic.com/cart?dwcont="+dwcont, data)
	handleErr(err)

	//Add Standard headers + Cookie
	_cookies := []string{""}
	req.Header = addStandardHeaders(Header{cookie: _cookies, content: data})

	//POST
	resp, err := client.Do(req)
	handleErr(err)

	//cleanup
	defer resp.Body.Close()
	defer req.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}
func SubmitShipping(client http.Client, dwcont string, secureKey string) string {
	// /https://www.hottopic.com/cart?dwcont=C276350184
	//payload
	//dwfrm_cart_checkoutCart: Checkout
	dataCreate := url.Values{
		"dwfrm_login_unregistered": {"Checkout As a Guest"},
		"dwfrm_login_securekey":    {secureKey},
	}
	data := strings.NewReader(dataCreate.Encode())

	//Create HTTP POST request
	req, err := http.NewRequest("POST", "https://www.hottopic.com/cart?dwcont="+dwcont, data)
	handleErr(err)

	//Add Standard headers + Cookie
	_cookies := []string{""}
	req.Header = addStandardHeaders(Header{cookie: _cookies, content: data})

	//POST
	resp, err := client.Do(req)
	handleErr(err)

	//cleanup
	defer resp.Body.Close()
	defer req.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}

//do GEt request to https://www.hottopic.com/cart
//to get the ID from the 'form action'
