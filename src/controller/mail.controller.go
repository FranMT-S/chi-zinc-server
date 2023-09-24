package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	constants_response "github.com/FranMT-S/chi-zinc-server/src/constants/response"
	myDatabase "github.com/FranMT-S/chi-zinc-server/src/db"
	"github.com/FranMT-S/chi-zinc-server/src/model"
	"github.com/go-chi/chi/v5"
)

const ()

func GetTotalMail(res http.ResponseWriter, req *http.Request) {

	code := 0

	dbRespBody, responseError := myDatabase.ZincDatabase().GetIndexData()

	if responseError != nil {
		code = http.StatusInternalServerError
		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		log.Println("Error en la peticion a la base de datos:", responseError.Err)
	}

	defer dbRespBody.Close()

	var ResponseIndexData model.ResponseIndexData
	err := json.NewDecoder(dbRespBody).Decode(&ResponseIndexData)

	if err != nil {
		code = http.StatusInternalServerError
		responseError = model.NewResponseError(code, constants_response.STATUS_ERROR, constants_response.ERROR_SERVER)

		res.WriteHeader(code)
		res.Write([]byte(responseError.Error()))
		log.Println("Error al decodificar la respuesta JSON:", err)
		return
	}

	code = http.StatusOK
	res.WriteHeader(code)
	res.Write([]byte(fmt.Sprintf(`
	{
		"status":%v,
		"msg":"%v",
		"total":%v
	}`, code, constants_response.STATUS_OK, ResponseIndexData.Stats.DocNum)))
}

func GetAllMailsSummary(res http.ResponseWriter, req *http.Request) {

	from, errFrom := strconv.Atoi(chi.URLParam(req, "from"))
	max, errMax := strconv.Atoi(chi.URLParam(req, "max"))
	code := 0

	if errFrom != nil || errMax != nil {
		error := model.ResponseError{
			Status: http.StatusBadRequest,
			Msg:    constants_response.STATUS_ERROR,
			Err:    "Los datos ingresados deben ser numeros"}

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(error.Error()))
		log.Println("Los datos ingresados deben ser numeros")
		return
	}

	if from < 0 || max < 0 {
		error := model.ResponseError{
			Status: http.StatusBadRequest,
			Msg:    constants_response.STATUS_ERROR,
			Err:    "Lo campos no pueden ser menores de 0"}

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(error.Error()))
		log.Println("Lo campos no pueden ser menores de 0")
		return
	}

	dbHits, responseError := myDatabase.ZincDatabase().GetAllMailsSummary(from, max)

	if responseError != nil {
		code = http.StatusInternalServerError
		log.Println("Error en la peticion a la base de datos:", responseError.Err)
		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		return
	}

	data, err := json.Marshal(dbHits)
	if err != nil {
		code = http.StatusInternalServerError
		log.Println("Error en convertir a json la informacion")
		res.WriteHeader(code)
		res.Write([]byte(err.Error()))
		return
	}

	code = http.StatusOK
	res.WriteHeader(code)
	res.Write([]byte(fmt.Sprintf(`
	{
		"status":%v,
		"msg":"%v",
		"data":%v
	}`, code, "OK", string(data))))
}

func FindMailsSummary(res http.ResponseWriter, req *http.Request) {

	from, errFrom := strconv.Atoi(chi.URLParam(req, "from"))
	max, errMax := strconv.Atoi(chi.URLParam(req, "max"))
	terms := chi.URLParam(req, "terms")
	code := 0

	if errFrom != nil || errMax != nil {
		error := model.ResponseError{
			Status: http.StatusBadRequest,
			Msg:    constants_response.STATUS_ERROR,
			Err:    "Los datos ingresados deben ser numeros"}

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(error.Error()))
		log.Println("Los datos ingresados deben ser numeros")
		return
	}

	if from < 0 || max < 0 {
		error := model.ResponseError{
			Status: http.StatusBadRequest,
			Msg:    constants_response.STATUS_ERROR,
			Err:    "Lo campos no pueden ser menores de 0"}

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(error.Error()))
		log.Println("Lo campos no pueden ser menores de 0")
		return
	}

	dbHits, responseError := myDatabase.ZincDatabase().FindMailsSummary(terms, from, max)

	if responseError != nil {
		code = http.StatusInternalServerError
		log.Println("Error en la peticion a la base de datos:", responseError.Err)
		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		return
	}

	data, err := json.Marshal(dbHits)
	if err != nil {
		code = http.StatusInternalServerError
		log.Println("Error en convertir a json la informacion")
		res.WriteHeader(code)
		res.Write([]byte(err.Error()))
		return
	}

	code = http.StatusOK
	res.WriteHeader(code)
	res.Write([]byte(fmt.Sprintf(`
	{
		"status":%v,
		"msg":"%v",
		"data":%v
	}`, code, "OK", string(data))))
}

func GetMail(res http.ResponseWriter, req *http.Request) {

	code := 0

	id := chi.URLParam(req, "id")

	dbMail, responseError := myDatabase.ZincDatabase().GetMail(id)

	if responseError != nil {
		code = http.StatusInternalServerError
		log.Println("Error en la peticion a la base de datos:", responseError.Err)
		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		return
	}

	data, err := dbMail.ToJsonBytes()
	if err != nil {
		code = http.StatusInternalServerError
		log.Println("Error en convertir a json la informacion")
		res.WriteHeader(code)
		res.Write([]byte(err.Error()))
		return
	}

	code = http.StatusOK
	res.WriteHeader(code)
	res.Write([]byte(fmt.Sprintf(`
	{
		"status":%v,
		"msg":"%v",
		"data":%v
	}`, code, "OK", string(data))))
}
