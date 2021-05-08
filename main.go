package main

import (
	"fmt"
	"net/http/cookiejar"

	cclient "github.com/IHaveNothingg/cclientwtf"
	tls "github.com/refraction-networking/utls"
)

func main() {
	cookieJar, err := cookiejar.New(nil)
	handleErr(err)
	client, err := cclient.NewClient(tls.HelloChrome_Auto, "http://localhost:8866")
	handleErr(err)
	client.Jar = cookieJar

	//Add to cart
	AddToCart(client)
	//Goto checkoutpage for dwCont
	checkOutPage := getDwCont(GetCheckout(client))

	//click 'proceed to checkout'
	prooceedToCheckoutPage := ProceedToCheckout(client, checkOutPage)

	//get dwCont and secureKey from the checkout page
	dwcont := getDwCont(prooceedToCheckoutPage)
	secureKey := getStringInBetweenTwoString(prooceedToCheckoutPage, "dwfrm_login_securekey\" value=\"", "\"/>")

	//click Guest checkout and get dwCont and secureKey from the guest checkout page
	guestCheckoutPage := GuestCheckout(client, dwcont, secureKey)
	dwcont = getDwCont(guestCheckoutPage)
	secureKey = getStringInBetweenTwoString(guestCheckoutPage, "dwfrm_login_securekey\" value=\"", "\"/>")

	//click submit shipping
	submitShippingage := SubmitShipping(client, dwcont, secureKey)
	fmt.Println(submitShippingage)

}
