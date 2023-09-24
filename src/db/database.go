package myDatabase

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	constants_response "github.com/FranMT-S/chi-zinc-server/src/constants/response"
	myMiddleware "github.com/FranMT-S/chi-zinc-server/src/middleware"
	"github.com/FranMT-S/chi-zinc-server/src/model"
)

var z_database *zincDatabase

type zincDatabase struct {
	client *http.Client
}

func ZincDatabase() *zincDatabase {
	if z_database == nil {
		z_database = &zincDatabase{client: &http.Client{}}
	}

	return z_database
}

func (db zincDatabase) GetIndexData() (io.ReadCloser, *model.ResponseError) {

	url := os.Getenv("URL") + "index/" + os.Getenv("INDEX")

	dbResp, err := db.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return dbResp.Body, nil
}

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

		return nil, model.NewResponseError(http.StatusInternalServerError, constants_response.STATUS_ERROR,
			"Hubo un error en el servidor: "+err.Error())

	}

	return &ResponseData.Hits, nil
}

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

		return nil, model.NewResponseError(http.StatusInternalServerError, constants_response.STATUS_ERROR,
			"Hubo un error en el servidor: "+err.Error())

	}

	return &ResponseData.Hits, nil
}

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

		return nil, model.NewResponseError(http.StatusInternalServerError, constants_response.STATUS_ERROR,
			"Hubo un error en el servidor: "+err.Error())

	}

	return &ResponseData.Mail, nil
}

func (db zincDatabase) doRequest(method string, url string, body io.Reader) (*http.Response, *model.ResponseError) {

	// Realizar la solicitud
	dbReq, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println("Error al crear la solicitud:", err)
		return nil, model.NewResponseError(http.StatusBadRequest, constants_response.STATUS_ERROR, constants_response.ERROR_CREATE_REQUEST)
	}

	myMiddleware.ZincHeader(dbReq)

	dbResp, err := db.client.Do(dbReq)
	if err != nil {
		log.Println("Error al realizar la solicitud:", err)
		return nil, model.NewResponseError(http.StatusBadRequest, constants_response.STATUS_ERROR, constants_response.ERROR_REQUEST)
	}

	// Verificar el código de estado de la respuesta
	if dbResp.StatusCode != http.StatusOK {
		log.Println("Respuesta no exitosa. Código de estado:", dbResp.Status)
		return nil, model.NewResponseError(dbResp.StatusCode, constants_response.STATUS_ERROR, constants_response.ERROR_REQUEST)
	}

	return dbResp, nil
}
