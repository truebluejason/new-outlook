package resp

import (
	"fmt"
	"github.com/truebluejason/new-outlook/internal/data"
)

func FormatSSML(data data.DBResponse) string {
	var (
		response string
	)
	response = fmt.Sprint("<speak>", data.Description, "</speak>")
	return response
}