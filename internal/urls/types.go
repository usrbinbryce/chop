package urls

import "errors"

var (
	ErrShortCodeRequired   = errors.New("short code is required")
	ErrDestinationRequired = errors.New("destination is required")
)

type createUrlRequest struct {
	ShortCode   string `json:"short_code"`
	Destination string `json:"destination"`
}
