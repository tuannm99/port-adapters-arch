package httpx

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(r *http.Request, out any) error {
	return json.NewDecoder(r.Body).Decode(out)
}
