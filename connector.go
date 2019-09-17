package odoo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Init initializes an Odoo Connection
func Init(protocol, host string, port int, db, username, password string) Connection {
	conn := Connection{
		Protocol: protocol,
		Host:     host,
		Port:     port,
		DB:       db,
		Username: username,
		Password: password,
	}

	conn.connect()

	return conn
}

// Fetch sends requests that does not apply to a specific record or set of records.
func (connection *Connection) Fetch(requestParams *RequestParams) []byte {
	buildKwargs := func() map[string]interface{} {
		kwargs := map[string]interface{}{
			"context": connection.Session.Context,
		}

		if 0 < len(requestParams.Fields) {
			kwargs["fields"] = requestParams.Fields
		} else {
			if requestParams.Method != MethodSearchCount {
				kwargs["fields"] = [][]string{}
			}
		}

		if 0 < len(requestParams.Domains) {
			searches := make([]interface{}, 0)

			for _, domain := range requestParams.Domains {
				searches = append(searches, []interface{}{
					domain.Field,
					domain.Operator,
					domain.Value,
				})
			}

			kwargs["domain"] = searches
		}

		if requestParams.Order != "" {
			kwargs["order"] = requestParams.Order
		} else {
			if requestParams.Method != MethodSearchCount {
				kwargs["order"] = "id desc"
			}
		}

		if requestParams.Limit != 0 {
			kwargs["offset"] = requestParams.Offset
			kwargs["limit"] = requestParams.Limit
		} else {
			if requestParams.Method != MethodSearchCount {
				kwargs["offset"] = 0
				kwargs["limit"] = 10
			}
		}

		return kwargs
	}

	buildParams := func() map[string]interface{} {
		params := map[string]interface{}{
			"model":  requestParams.Model,
			"method": requestParams.Method,
			"kwargs": buildKwargs(),
		}

		if requestParams.ID != 0 {
			params["id"] = requestParams.ID
		}

		if 0 < len(requestParams.Args) {
			params["args"] = requestParams.Args
		} else {
			if requestParams.Method != MethodSearchCount {
				params["args"] = map[string]interface{}{}
			}
		}

		return params
	}

	params := map[string]interface{}{
		"jsonrpc": 2.0,
		"method":  "call",
		"params":  buildParams(),
	}

	jsonParams, _ := json.Marshal(params)

	res, error := connection.authenticatedRequest(connection.URL, http.MethodPost, jsonParams)

	if error != nil {
		log.Fatalln(error)
	}

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	return data
}

func (connection *Connection) authenticatedRequest(url, method string, jsonValue []byte) (*http.Response, error) {
	return connection.request(url, method, jsonValue, map[string]string{
		"Cookie": connection.Session.SessionString,
	})
}

func (connection *Connection) unAuthenticatedRequest(url, method string, jsonValue []byte) (*http.Response, error) {
	return connection.request(url, method, jsonValue, map[string]string{})
}

func (connection *Connection) request(url, method string, jsonValue []byte, headers map[string]string) (*http.Response, error) {
	req, error := http.NewRequest(method, url, bytes.NewBuffer(jsonValue))

	if error != nil {
		log.Fatalln(error)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Length", string(jsonValue))

	for header := range headers {
		req.Header.Add(header, headers[header])
	}

	client := &http.Client{}

	return client.Do(req)
}

func (connection *Connection) connect() {
	url := connection.Protocol + "://" + connection.Host + ":" + strconv.Itoa(connection.Port) + "/web/session/authenticate"
	authResult := AuthResult{}
	jsonValue, _ := json.Marshal(map[string]map[string]string{
		"params": {
			"login":    connection.Username,
			"password": connection.Password,
			"db":       connection.DB,
		},
	})

	if res, error := connection.unAuthenticatedRequest(url, http.MethodPost, jsonValue); error == nil {
		defer res.Body.Close()
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &authResult)
		cookie := res.Header.Get("set-cookie")
		connection.Session = authResult.Result
		connection.Session.SessionString = strings.Split(cookie, ";")[0]
		connection.URL = fmt.Sprintf(
			"%v://%v:%v/web/dataset/call_kw",
			connection.Protocol,
			connection.Host,
			connection.Port,
		)
	} else {
		log.Fatalln(error)
	}
}
