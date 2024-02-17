package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// declare const
	TOKEN := "kjdlgfjdofjel"
	CHAT_ID := "34322344432"

	//  Parse command line flags.
	msg := flag.String("msg", "Hello World!", "Message to be sent.")
	flag.Parse()

	// Send message to the user
	url := "https://api.telegram.org/bot" + TOKEN + "/sendMessage?chat_id=" + CHAT_ID + "&text=" + *msg
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
