package models

type Email struct {
	MessageId               string `json:"message_id"`
	Date                    string `json:"date"`
	From                    string `json:"from"`
	Subject                 string `json:"subject"`
	MimeVersion             string `json:"mime_version"`
	ContentType             string `json:"content_type"`
	ContentTransferEncoding string `json:"content_transfer_encoding"`
	XFrom                   string `json:"x_from"`
	XTo                     string `json:"x_to"`
	XCc                     string `json:"x_cc"`
	XBcc                    string `json:"x_bcc"`
	XFolder                 string `json:"x_folder"`
	XOrigin                 string `json:"x_origin"`
	XFileName               string `json:"x_fileName"`
	Body                    string `json:"body"`
}

type SearchResult struct {
	Emails []Email `json:"emails"`
	Total  int     `json:"total"`
	Page   int     `json:"page"`
	Limit  int     `json:"limit"`
	Pages  int     `json:"pages"`
}

type searchResponse struct {
	Total int `json:"total"`
	Hits  []struct {
		MessageId               string `json:"message_id"`
		Date                    string `json:"date"`
		From                    string `json:"from"`
		Subject                 string `json:"subject"`
		MimeVersion             string `json:"mime_version"`
		ContentType             string `json:"content_type"`
		ContentTransferEncoding string `json:"content_transfer_encoding"`
		XFrom                   string `json:"x_from"`
		XTo                     string `json:"x_to"`
		XCc                     string `json:"x_cc"`
		XBcc                    string `json:"x_bcc"`
		XFolder                 string `json:"x_folder"`
		XOrigin                 string `json:"x_origin"`
		XFileName               string `json:"x_fileName"`
		Body                    string `json:"body"`
	} `json:"hits"`
}
