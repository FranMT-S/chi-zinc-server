package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	constants_log "github.com/FranMT-S/chi-zinc-server/src/constants/log"
	myDatabase "github.com/FranMT-S/chi-zinc-server/src/db"
	_logs "github.com/FranMT-S/chi-zinc-server/src/logs"
	"github.com/FranMT-S/chi-zinc-server/src/model"
	"github.com/go-chi/chi/v5"
)

/*
GetAllMailsSummary return all emails, takes two URL query parameters, "from" and "max", which are used for result pagination.

  - from (int): the "from" parameter represents the index from which the search will begin.
  - max (int): The "max" parameter specifies the maximum number of elements to display.

If the request is successful return:

	{
			"status":code,
			"msg":"message",
			"data":Hits
	}

if failed then return a ResponseError object
*/
func GetAllMailsSummary(res http.ResponseWriter, req *http.Request) {

	from, errFrom := strconv.Atoi(chi.URLParam(req, "from"))
	max, errMax := strconv.Atoi(chi.URLParam(req, "max"))
	code := 0

	if errFrom != nil || errMax != nil {
		err := model.NewResponseError(
			http.StatusBadRequest,
			constants_log.ERROR_FROM_MAX_IS_NOT_NUMBER,
			fmt.Errorf(constants_log.ERROR_INVALID_PARAMS),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_MAILS_REQUEST,
			constants_log.ERROR_FROM_MAX_IS_NOT_NUMBER,
			err.Err,
		)

		res.WriteHeader(err.Status)
		res.Write([]byte(err.Error()))

		return
	}

	if from < 0 || max < 0 {
		err := model.NewResponseError(
			http.StatusBadRequest,
			constants_log.ERROR_VALUE_LESS_ZERO,
			fmt.Errorf(constants_log.ERROR_INVALID_PARAMS),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_MAILS_REQUEST,
			constants_log.ERROR_VALUE_LESS_ZERO,
			err.Err,
		)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))

		return
	}

	dbHits, responseError := myDatabase.ZincDatabase().GetAllMailsSummary(from, max)
	if responseError != nil {
		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_DATABASE,
			responseError.Msg,
			responseError.Err,
		)

		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		return
	}

	data, err := json.Marshal(dbHits)
	if err != nil {

		err := model.NewResponseError(
			http.StatusInternalServerError,
			constants_log.ERROR_JSON_PARSE,
			fmt.Errorf(constants_log.ERROR_JSON_PARSE),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_DATABASE,
			constants_log.ERROR_JSON_PARSE,
			err,
		)

		res.WriteHeader(err.Status)
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

// FindMailsSummary return all emails that match the request in the terms parameter
// The function takes two URL query parameters, "from" and "max", which are used for result pagination.
//
//   - from (int): the "from" parameter represents the index from which the search will begin.
//
//   - max (int): The "max" parameter specifies the maximum number of elements to display.
//
//   - terms (string): the words or query that will be used for the search.
//
// If the request is successful return:
//
//	{
//			"status":code,
//			"msg":"message",
//			"data":Hits
//	}
//
// where Hits is:
//
// if failed then return a ResponseError object
//
// The searches in Terms are composed this way:
//
//  1. %20 instead of blank space = search for any match of the terms
//
//  2. + = returns all data where both terms appear
//
//  3. - = returns all data where the terms do not appear
//
//  4. * = returns all the data where it starts with the term
//
// # example:
//
//   - susan = find all matches of susan in all fields
//   - susan%20bianca (instead of "susan bianca") = find all matches of susan or bianca in all fields
//   - -susan = all matches where susan is not in all fields
//   - susan.bailey +bianca.ornelas = all matches where this susan and bianca.ornelas in all fields
//   - susan* = all matches starting with susan in all fields
//   - -susan*=all matches you start that do not start with susan in all fields
//   - From:susan = all susan matches in the From field
//   - -From:susan = all non-susan matches in the field
//   - From:susan* = all matches in From that start with susan
//   - -From:susan* = all matches in From that do not start with susan
//   - +From:susan.bailey%20+To:bianca.ornelas = all matches in From de susan.bailey and in To de bianca.ornelas
//
// The fields where you can search are:
//  1. Message_ID,
//  2. From
//  3. To
//  4. Subject
//  5. Cc
//  6. Mime_Version
//  7. Content_Type
//  8. Content_Transfer_Encoding
//  9. Bcc
//  10. X_From
//  11. X_To
//  12. X_cc
//  13. X_bcc
//  14. X_Folder
//  15. X_Origin
//  16. X_FileName
//  17. Content
func FindMailsSummary(res http.ResponseWriter, req *http.Request) {

	from, errFrom := strconv.Atoi(chi.URLParam(req, "from"))
	max, errMax := strconv.Atoi(chi.URLParam(req, "max"))
	terms := chi.URLParam(req, "terms")

	code := 0

	decodedString, err := url.QueryUnescape(terms)
	if err != nil {
		errorResponse := model.NewResponseError(
			http.StatusInternalServerError,
			constants_log.ERROR_QUERY_DECODE,
			err,
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_MAILS_REQUEST,
			constants_log.ERROR_FROM_MAX_IS_NOT_NUMBER,
			err,
		)

		res.WriteHeader(errorResponse.Status)
		res.Write([]byte(errorResponse.Error()))

		return
	}

	terms = decodedString

	if errFrom != nil || errMax != nil {
		err := model.NewResponseError(
			http.StatusBadRequest,
			constants_log.ERROR_FROM_MAX_IS_NOT_NUMBER,
			fmt.Errorf(constants_log.ERROR_INVALID_PARAMS),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_MAILS_REQUEST,
			constants_log.ERROR_FROM_MAX_IS_NOT_NUMBER,
			err.Err,
		)

		res.WriteHeader(err.Status)
		res.Write([]byte(err.Error()))

		return
	}

	if from < 0 || max < 0 {
		err := model.NewResponseError(
			http.StatusBadRequest,
			constants_log.ERROR_VALUE_LESS_ZERO,
			fmt.Errorf(constants_log.ERROR_INVALID_PARAMS),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_MAILS_REQUEST,
			constants_log.ERROR_VALUE_LESS_ZERO,
			err.Err,
		)

		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))

		return
	}

	dbHits, responseError := myDatabase.ZincDatabase().FindMailsSummary(terms, from, max)

	if responseError != nil {
		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_DATABASE,
			responseError.Msg,
			responseError.Err,
		)

		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		return
	}

	data, err := json.Marshal(dbHits)
	if err != nil {
		err := model.NewResponseError(
			http.StatusInternalServerError,
			constants_log.ERROR_JSON_PARSE,
			fmt.Errorf(constants_log.ERROR_JSON_PARSE),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_DATABASE,
			constants_log.ERROR_JSON_PARSE,
			err,
		)

		res.WriteHeader(err.Status)
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

// GetMail return all emails that match with id param,
//
// query parameters:.
//
//   - id (string): the id of the searched email.
//
// If the request is successful return:
//
//	{
//			"status":code,
//			"msg":"message",
//			"data":Mail
//	}
//
// if failed then return is ResponseError
func GetMail(res http.ResponseWriter, req *http.Request) {

	code := 0

	id := chi.URLParam(req, "id")

	dbMail, responseError := myDatabase.ZincDatabase().GetMail(id)
	if responseError != nil {
		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_DATABASE,
			responseError.Msg,
			responseError.Err,
		)

		res.WriteHeader(responseError.Status)
		res.Write([]byte(responseError.Error()))
		return
	}

	data, err := dbMail.ToJsonBytes()
	if err != nil {
		err := model.NewResponseError(
			http.StatusInternalServerError,
			constants_log.ERROR_JSON_PARSE,
			fmt.Errorf(constants_log.ERROR_JSON_PARSE),
		)

		_logs.LogSVG(
			constants_log.FILE_NAME_ERROR_GENERAL,
			constants_log.OPERATION_DATABASE,
			constants_log.ERROR_JSON_PARSE,
			err,
		)

		res.WriteHeader(err.Status)
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
