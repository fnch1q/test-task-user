package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Success struct {
	StatusCode int `json:"-"`
	Result     interface{}
	Total      int64 `json:"total,omitempty"`
}

func NewSuccess(result interface{}, status int) Success {
	return Success{
		StatusCode: status,
		Result:     result,
	}
}

func (r Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	if r.Result != nil {
		return json.NewEncoder(w).Encode(r.Result)
	}
	return nil
}

// func NewSuccessList(result interface{}, total int64, status int) Success {
// 	return Success{
// 		StatusCode: status,
// 		Result:     result,
// 		Total:      total,
// 	}
// }

func NewSuccessFile(w http.ResponseWriter, filename string, content []byte) error {
	contentType := http.DetectContentType(content)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Length", strconv.Itoa(len(content)))

	_, err := w.Write(content)
	return err
}
