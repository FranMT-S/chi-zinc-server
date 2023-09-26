# Chi-zinc-server

This project works in conjunction with two other projects. the [Fronted](https://github.com/FranMT-S/fronted) which is responsible for reading the files and  the [indexer](https://github.com/FranMT-S/Challenge-Go) used to index the data and upload it to the database

This project endpoints to request the emails indexed in the database.

The database used in this project is the [Enron Corp](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)

# Enviroment Variable
In the path `src/constanst/env.const.go` there is the environment variable INDEX which is the name of the index used in the database make sure it is the same in the [indexer](https://github.com/FranMT-S/Challenge-Go)


# Run
Execute Command:
``` Go run main.go```
  

# End points

 
## Get Mails

Return a mailing list

 
 <summary><code>GET</code> <code><b>/</b></code> <code>http://localhost:3000/v1/api/from-{from}-max-{max}</code></summary>

The parameters are:
- **from (int)**: index from where the search would start
- **max (int)**: the total number of emails that will be returned

Response Success return a [ResponseHits](#response-hits) :
`Code:200`
```  
interface ResponseHits {
	"status":  int,
	"msg":  string,
	"data":  {
		"total":{
			"value":  int
		},
		"hits":  Hit[]
	}
}
```
[Hit](#hit)

Failed response returns [Response Error Interface](#response-error) 
`Code:400`

## Find Mails

Find emails that match the requested query

 
 <summary><code>GET</code> <code><b>/</b></code> <code>http://localhost:3000/v1/api/from-{from}-max-{max}-terms-{terms}</code></summary>

The parameters are:
- **from (int)**: index from where the search would start
- **max (int)**: the total number of emails that will be returned
- **terms (string)**: the query used for search

The searches in Terms are composed this way:

1)  `%20` instead of blank space = search for any match of the terms.
2)  `+` used to returns all data where both terms appear.
3)  `-` used to returns all data where the terms do not appear.
4) `*` used to returns all the data where it starts with the term.

#### example terms:
 - `susan`  find all matches of susan in all fields
 - `susan%20bianca` (instead of "susan bianca")  find all matches of susan or bianca in all fields
 - `-susan`  all matches where susan is not in all fields
 - `susan.bailey +bianca.ornelas`  all matches where this susan and bianca.ornelas in all fields
 - `susan*`  all matches starting with susan in all fields
 - `-susan*` all matches you start that do not start with susan in all fields
 - `From:susan`   all susan matches in the From field
 - `-From:susan`   all non-susan matches in the field
 - `From:susan*`  all matches in From that start with susan
 - `-From:susan*`  all matches in From that do not start with susan
 - `+From:susan.bailey%20+To:bianca.ornelas`  all matches in From de susan.bailey and in To de bianca.ornelas

Response Success return a [ResponseHits](#response-hits) :
`Code:200`
```  
interface ResponseHits {
	"status":  int,
	"msg":  string,
	"data":  {
		"total":{
			"value":  int
		},
		"hits":  Hit[]
	}
}
```
[Hit](#hit)

Failed response returns [Response Error Interface](#response-error) 
`Code:400`

## Get Mail

Return a mail

 
 <summary><code>GET</code> <code><b>/</b></code> <code>http://localhost:3000/v1/api/{id}</code></summary>

The parameters are:
- **id (string)**: ID of the requested email

Response Success return a response with a [Mail](#mail) :
```
interface ResponseMail {
	"status":  int,
	"msg":  string,
	"data": Mail
}
```

`Code:200`
 

# Interfaces Responses

## Response Hits

```
interface ResponseHits {
	"status":  int,
	"msg":  string,
	"data":  {
		"total":{
			"value":  int
		},
		"hits":  Hit[]
	}
}

```

## Hit

One Source is equivalent to Mail Resummary
```
interface Hit {
    _index:  string;
    _id:     string;
    _source: Source;
}

interface Source {
    To:      string;
    From:    string;
    Subject: string;
    Date:    Date;
}
```

## Mail
```
interface Mail {
    Message_ID:                string;
    Date:                      Date;
    From:                      string;
    To:                        string;
    Subject:                   string;
    Cc:                        string;
    Mime_Version:              string;
    Content_Type:              string;
    Content_Transfer_Encoding: string;
    Bcc:                       string;
    X_From:                    string;
    X_To:                      string;
    X_cc:                      string;
    X_bcc:                     string;
    X_Folder:                  string;
    X_Origin:                  string;
    X_FileName:                string;
    Content:                   string;
}
```

## Response Error

```
interface ResponseError {
    status: number;
    msg:    string;
    error:  string;
}

```

