package model

import "time"

type ResponseSearchData struct {
	Hits Hits `json:"hits"`
}

//  used to send a Mail
type ResponseDocData struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	Mail  Mail   `json:"_source"`
}

// A hit represents a match as provided by [Zincsearch] models
// represents a summary format of an email
//
// [Zincsearch]: https://zincsearch-docs.zinc.dev/api/search/search/#response
type Hit struct {
	To      string    `json:"To"`
	From    string    `json:"From"`
	Subject string    `json:"Subject"`
	Date    time.Time `json:"Date"`
}

// Represents a collection of Hits
//
// A Hits represents a match as provided by [Zincsearch] models
//
// [Zincsearch]: https://zincsearch-docs.zinc.dev/api/search/search/#response
type Hits struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`

	Hits []struct {
		Index  string `json:"_index"`
		ID     string `json:"_id"`
		Source Hit    `json:"_source"`
	} `json:"hits"`
}

// Basic structure of a post request
type RequestFindMail struct {
	Terms string `json:"Terms"`
	From  int    `json:"From"`
	Max   int    `json:"Max"`
}
