package model

import "time"

// Para Obtener el numero total de mensajes
type Stats struct {
	DocNum int `json:"doc_num"`
}

type ResponseIndexData struct {
	Stats Stats `json:"stats"`
}

type ResponseSearchData struct {
	Hits Hits `json:"hits"`
}

type ResponseDocData struct {
	Index string `json:"_index"`
	ID    string `json:"_id"`
	Mail  Mail   `json:"_source"`
}

type Hit struct {
	To      string    `json:"To"`
	From    string    `json:"From"`
	Subject string    `json:"Subject"`
	Date    time.Time `json:"Date"`
}

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

type RequestFindMail struct {
	Terms string `json:"Terms"`
	From  int    `json:"From"`
	Max   int    `json:"Max"`
}
