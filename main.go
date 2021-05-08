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

	//Go to checkout as guest using dwCont

	//res := strings.Contains(string(resp), "")

	//After adding to cart we need to 'checkout as guest'
	//Go to www.hottopic.com/cart.
	//find ""

	//then we submit shipping info to same cart address with this dwordcont
	//not sure yet if they're the same (i assme they are)
	//dwcont: C251111390

	//dwcont could be in the html, seems to change from each request.

	//	Checkout As a Guest

	//HTML := Checkout(client)
	//doc, err := goquery.NewDocumentFromReader(HTML.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
