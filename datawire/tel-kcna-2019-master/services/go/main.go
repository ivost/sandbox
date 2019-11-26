package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getValue() string {
	res, err := http.Get("http://base")
	if err != nil {
		return err.Error()
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func getHostname() string {
	res, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return res
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		value := "salty" + getValue()
		hashcode := base64.StdEncoding.EncodeToString([]byte(value))
		lines := []string{
			"[ Hello KubeCon NA 2019! ]",
			"[ Greetings from Go      ]",
			fmt.Sprintf("[ Code: %s ]", hashcode),
			"",
			fmt.Sprintf("Host: %s", getHostname()),
			fmt.Sprintf("Now:  %s", time.Now().Format(time.RFC3339Nano)),
		}
		res := strings.Join(lines, "\n") + "\n"
		fmt.Fprint(w, res)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
