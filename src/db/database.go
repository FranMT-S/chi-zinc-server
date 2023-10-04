package myDatabase

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	constants_log "github.com/FranMT-S/chi-zinc-server/src/constants/log"
	myMiddleware "github.com/FranMT-S/chi-zinc-server/src/middleware"
	"github.com/FranMT-S/chi-zinc-server/src/model"
)

var z_database *zincDatabase

type zincDatabase struct {
	client *http.Client
}

// returns a single instance of the database
func ZincDatabase() *zincDatabase {
	if z_database == nil {
		z_database = &zincDatabase{client: &http.Client{}}
	}

	return z_database
}

// GetAllMailsSummary return a Hits object that containst information of mails request
// The function takes two URL query parameters, "from" and "max", which are used for result pagination.
//
// - from (int): the "from" parameter represents the index from which the search will begin.
// - max (int): The "max" parameter specifies the maximum number of elements to display.
//
// if the request failed return a ResponseError object
func (db zincDatabase) GetAllMailsSummary(from, max int) (*model.Hits, *model.ResponseError) {

	var ResponseData model.ResponseSearchData

	query := fmt.Sprintf(`
		{
			"search_type": "matchall",
			"sort_fields": ["-Date"],
			"from": %v,
			"max_results": %v,
			"_source": [
			"To", "From","Date", "Subject"
			]
		}`, from, max)

	url := os.Getenv("URL") + os.Getenv("INDEX") + "/_search"

	dbResp, errResponse := db.doRequest("POST", url, strings.NewReader(query))

	if errResponse != nil {
		return nil, errResponse
	}

	defer dbResp.Body.Close()

	err := json.NewDecoder(dbResp.Body).Decode(&ResponseData)

	if err != nil {
		return nil, model.NewResponseError(http.StatusInternalServerError, constants_log.ERROR_JSON_PARSE, err)

	}

	return &ResponseData.Hits, nil
}

// GetAllMailsSummary return a Hits object that containst information of mails that match the request in the terms parameter
//
// The function takes two URL query parameters, "from" and "max", which are used for result pagination.
//
// if the request failed return a ResponseError object
func (db zincDatabase) FindMailsSummary(term string, from, max int) (*model.Hits, *model.ResponseError) {

	var ResponseData model.ResponseSearchData

	query := fmt.Sprintf(`
		{
		"search_type": "querystring",
		"query": {
			"term": "%v",
			"field":"_all"
		},
		"sort_fields": ["-Date"],
		"from": %v,
		"max_results": %v,
		"_source": [
			"Date", "From","Subject", "To"
		]
	}`, term, from, max)

	url := os.Getenv("URL") + os.Getenv("INDEX") + "/_search"

	dbResp, errResponse := db.doRequest("POST", url, strings.NewReader(query))

	if errResponse != nil {
		return nil, errResponse
	}

	defer dbResp.Body.Close()

	err := json.NewDecoder(dbResp.Body).Decode(&ResponseData)

	if err != nil {

		return nil, model.NewResponseError(http.StatusInternalServerError, constants_log.ERROR_JSON_PARSE, err)

	}

	return &ResponseData.Hits, nil
}

// GetMail return a ResponseDocData object that containt the mail that match with id param,
// query parameters:.
//
//   - id (string): the id of the searched email.
//
// if the request failed return a ResponseError object
func (db zincDatabase) GetMail(id string) (*model.Mail, *model.ResponseError) {

	var ResponseData *model.ResponseDocData

	url := os.Getenv("URL") + os.Getenv("INDEX") + "/_doc/" + id

	dbResp, errResponse := db.doRequest("GET", url, nil)

	if errResponse != nil {
		return nil, errResponse
	}

	defer dbResp.Body.Close()

	err := json.NewDecoder(dbResp.Body).Decode(&ResponseData)

	if err != nil {

		return nil, model.NewResponseError(http.StatusInternalServerError, constants_log.ERROR_JSON_PARSE, err)

	}

	return &ResponseData.Mail, nil
}

// doRequest Execute a request to database.
//
// if the request failed return a ResponseError
func (db zincDatabase) doRequest(method string, url string, body io.Reader) (*http.Response, *model.ResponseError) {

	dbReq, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(constants_log.ERROR_DATA_BASE_REQUEST+": ", err)
		return nil, model.NewResponseError(http.StatusBadRequest, constants_log.ERROR_DATA_BASE_REQUEST, err)
	}

	myMiddleware.ZincHeader(dbReq)

	dbResp, err := db.client.Do(dbReq)
	if err != nil {
		log.Println(constants_log.ERROR_DATA_BASE_REQUEST+": ", err)
		return nil, model.NewResponseError(http.StatusBadRequest, constants_log.ERROR_DATA_BASE_REQUEST, err)
	}

	if dbResp.StatusCode != http.StatusOK {
		return nil, model.NewResponseError(dbResp.StatusCode, constants_log.ERROR_INVALID_DATA_ENTERED, fmt.Errorf(constants_log.ERROR_INVALID_DATA_ENTERED))
	}

	return dbResp, nil
}
