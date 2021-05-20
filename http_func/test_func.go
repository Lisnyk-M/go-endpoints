package http_func

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func send() {
	client := &http.Client{}
	req, err := http.NewRequest(
		"GET", "https://health-fsoff2.herokuapp.com/products/?%D1%81%D1%8B%D1%80%D0%BE%D0%BA", nil,
	)
	// добавляем заголовки
	req.Header.Add("Accept", "text/html")     // добавляем заголовок Accept
	req.Header.Add("User-Agent", "MSIE/15.0") // добавляем заголовок User-Agent

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	// body, err := io.ReadAll(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	//
	// var result map[string]interface{}
	// json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result)

}
