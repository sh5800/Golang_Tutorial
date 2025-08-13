package utils

import(
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        return
    }
    defer r.Body.Close()

    if len(body) > 0 {
        err = json.Unmarshal(body, x)
        if err != nil {
            return
        }
    }
}
