package models

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Response events.APIGatewayProxyResponse

// RespondWithJSON ...
func RespondWithJSON(code int, payload interface{}) (Response, error) {
	var buf bytes.Buffer

	result, _ := json.Marshal(map[string]interface{}{
		"data": payload,
		"err":  nil,
	})

	json.HTMLEscape(&buf, result)

	resp := Response{
		StatusCode:      code,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}
