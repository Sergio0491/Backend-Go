package models

import (
	"encoding/json"
	"fmt"
)

func SearchEmails(query string, page int, limit int) (SearchResult, error) {
	sqHits, err := createHitsQuery(query, page, limit)
	if err != nil {
		return SearchResult{}, err
	}

	sqTotal, err := createTotalQuery(query)
	if err != nil {
		return SearchResult{}, err
	}

	hitsData, err := json.Marshal(sqHits)
	if err != nil {
		return SearchResult{}, fmt.Errorf("error serializando query hits: %v", err)
	}

	totalData, err := json.Marshal(sqTotal)
	if err != nil {
		return SearchResult{}, fmt.Errorf("error serializando query total: %v", err)
	}

	reqHits, err := buildRequest(hitsData)
	if err != nil {
		return SearchResult{}, err
	}

	respHits, err := doRequest(reqHits)
	if err != nil {
		return SearchResult{}, err
	}

	srHits, err := decodeResponse(respHits)
	if err != nil {
		return SearchResult{}, err
	}

	reqTotal, err := buildRequest(totalData)
	if err != nil {
		return SearchResult{}, err
	}

	respTotal, err := doRequest(reqTotal)
	if err != nil {
		return SearchResult{}, err
	}

	srTotal, err := decodeResponse(respTotal)
	if err != nil {
		return SearchResult{}, err
	}

	result := transformToSearchResult(srHits, page, limit)
	result.Total = srTotal.Total
	result.Pages = (srTotal.Total / result.Limit) + 1

	return result, nil
}

func SearchEmailByMessageID(messageID string) (Email, error) {
	if messageID == "" {
		return Email{}, fmt.Errorf("messageID no puede estar vacío")
	}

	query := fmt.Sprintf(`SELECT * FROM "email_records" WHERE message_id = '%s'`, messageID)

	result, err := SearchEmails(query, 1, 1)
	if err != nil {
		return Email{}, err
	}

	if len(result.Emails) == 0 {
		return Email{}, fmt.Errorf("no se encontró ningún email con message_id: %s", messageID)
	}

	return result.Emails[0], nil
}
