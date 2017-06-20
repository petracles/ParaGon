package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"encoding/base64"
	"log"

	"github.com/gorilla/mux"
)


// TODO: temporary vars until config.json reader is done
var epicApiKey, clientId, clientSecret, baseUri string =
"", "", "", "https://developer-paragon.epicgames.com"


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - - - STRUCTS - - - - - - - - - - -

// A person with an Epic account using this API
type Account struct {
	ID string
	Name string
}

// A hero from Paragon
type Hero struct {
	ID string
	Name []string
	Attack string
	Traits []string
	Scale string
	ReleaseDate string
	Affinities []string
	Difficulty string
	//Stats []
	//Images []
	Abilities []Abilities
}

type Abilities struct {
	Name []string
	Description []string
	ShortDescription []string
	Type string
	Binding string
	DamageType string
	MaxLevel int
	//ModifiersByLevel []
	//Images []
}

type Card struct {
	ID string
	Name []string
	SlotType string
	//Images []
	Rarity string
	Affinities []string
	Cost int
	UpgradeSlots int
	Effects []Effect
	MaxedEffects []Effect
}

type Effect struct {
	Description string
	//Value []
	Passive bool
	Unique bool
	Cooldown int
}


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - - - HELPERS - - - - - - - - - - -

// TODO: temporary html for index page
const loginhtml = `<html><body>
<a href="/prime/epic/login">Log in to your Epic account</a>
</body></html>`

/* Instantiate the client here, as net/http recommends using one client for reuse with
handler func requests to improve effeciency. */
var client = &http.Client{}

/* Makes a net/http Request, returns the request object or an error */
func makeRequest(method string, url string, headers map[string][]string, body io.Reader) *http.Response {
	// Make the request
	req, err1 := http.NewRequest(method, url, body)
	if err1 != nil {
		log.Println("http.NewRequest() FAILURE")
		//log.Panic()
		//log.Fatal()
	}

	// Add any headers passed in
	for key, values := range headers {
		for _, value := range values {
		        req.Header[key] = []string{value}
		    }
	}

	// Have the http.Client run the request
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Println("client.Do() FAILURE")
		//log.Panic()
		//log.Fatal()
	}

	// Return the response generated by the http.Client
	return resp
}


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - - INDEX ROUTE - - - - - - - - - -

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, loginhtml)
}


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - - AUTH ROUTES - - - - - - - - - -

func getLogin(w http.ResponseWriter, r *http.Request) {
	headers := make(map[string][]string)
	headers["X-Epic-ApiKey"] = []string{epicApiKey}
	headers["Accept"] = []string{"application/json"}

	resp := makeRequest(
		"GET", 
		baseUri + "/v1/auth/login/" + clientId,
		headers,
		nil)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		w.WriteHeader(http.StatusOK)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ioutil.ReadAll() FAILURE")
		}
		fmt.Fprintln(w, string(bodyBytes))
	}
}

func getAuthToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	b64 := []byte(clientId + ":" + clientSecret)

	headers := make(map[string][]string)
	headers["X-Epic-ApiKey"] = []string{epicApiKey}
	headers["Authorization"] = []string{"Basic " + base64.StdEncoding.EncodeToString(b64)}
	headers["Accept"] = []string{"application/json"}
	
	resp := makeRequest(
		"GET",
		baseUri + "/v1/auth/token/" + code,
		headers,
		nil)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		w.WriteHeader(http.StatusOK)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ioutil.ReadAll() FAILURE")
		}
		fmt.Fprintln(w, string(bodyBytes))
	}
}


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - ACCOUNT ROUTES - - - - - - - - - -

func getAccount(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// accountId := vars["accountId"]
	// req, err1 := http.NewRequest("GET", baseUri + "/v1/account/" + accountId, nil)
	// if err1 != nil {
	// 	log.Println("http.NewRequest() FAILURE")
	// }
	// req.Header["X-Epic-ApiKey"] = []string{epicApiKey}
	// req.Header["Accept"] = []string{"application/json"}
	// w.WriteHeader(http.StatusOK)

	// resp, err2 := client.Do(req)
	// if err2 != nil {
	// 	log.Println("client.Do() FAILURE")
	// }

	// defer resp.Body.Close()
	// if resp.StatusCode == 200 {
	// 	bodyBytes, err3 := ioutil.ReadAll(resp.Body)
	// 	if err3 != nil {
	// 		log.Println("ioutil.ReadAll() FAILURE")
	// 	}
	// 	fmt.Fprintln(w, string(bodyBytes))
	// }
}


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - CARDS ROUTES - - - - - - - - - -


// - - - - - - - - - - - - - - - - - - - - - - - - - -
// - - - - - - - - - HEROES ROUTES - - - - - - - - - -