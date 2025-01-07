package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func createHitsQuery(query string, page, limit int) (map[string]interface{}, error) {
	if page < 1 {
		return nil, fmt.Errorf("page debe ser mayor que 0")
	}
	if limit < 1 {
		return nil, fmt.Errorf("limit debe ser mayor que 0")
	}

	from := (page - 1) * limit

	sq := map[string]interface{}{
		"query": map[string]interface{}{
			"sql":        query,
			"start_time": time.Now().AddDate(-3, 0, 0).UnixMicro(),
			"end_time":   time.Now().UnixMicro(),
			"from":       from,
			"size":       limit,
		},
	}

	return sq, nil
}

func createTotalQuery(query string) (map[string]interface{}, error) {
	sq := map[string]interface{}{
		"query": map[string]interface{}{
			"sql":              query,
			"start_time":       time.Now().AddDate(-3, 0, 0).UnixMicro(),
			"end_time":         time.Now().UnixMicro(),
			"track_total_hits": true,
			"size":             0,
		},
	}

	return sq, nil
}

func buildRequest(queryData []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", openObserveSearchURL+"/_search", bytes.NewBuffer(queryData))
	if err != nil {
		return nil, fmt.Errorf("error creando solicitud: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(openObserveAuthUser, openObserveAuthPass)
	return req, nil
}

func doRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando solicitud: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return resp, nil
}

func decodeResponse(resp *http.Response) (searchResponse, error) {
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return searchResponse{}, fmt.Errorf("error leyendo respuesta: %v", err)
	}

	var sr searchResponse
	if err := json.Unmarshal(bodyBytes, &sr); err != nil {
		return searchResponse{}, fmt.Errorf("error decodificando respuesta: %v", err)
	}

	return sr, nil
}

func transformToSearchResult(sr searchResponse, page, limit int) SearchResult {
	emails := make([]Email, len(sr.Hits))
	for i, h := range sr.Hits {
		emails[i] = Email{
			MessageId:               h.MessageId,
			Body:                    h.Body,
			Date:                    h.Date,
			From:                    h.From,
			Subject:                 h.Subject,
			MimeVersion:             h.MimeVersion,
			ContentType:             h.ContentType,
			ContentTransferEncoding: h.ContentTransferEncoding,
			XFrom:                   h.XFrom,
			XTo:                     h.XTo,
			XCc:                     h.XCc,
			XBcc:                    h.XBcc,
			XFolder:                 h.XFolder,
			XOrigin:                 h.XOrigin,
			XFileName:               h.XFileName,
		}
	}

	return SearchResult{
		Emails: emails,
		Total:  0,
		Page:   page,
		Limit:  limit,
		Pages:  0,
	}
}
