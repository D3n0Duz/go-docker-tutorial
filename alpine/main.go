package main

import (
	"net/http"
)

func init() {

}

const PORT_DEFAULT = ":8080"

func main() {

	r := ChiRouter().InitRouter()

	http.ListenAndServe(PORT_DEFAULT, r)
}
