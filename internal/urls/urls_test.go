package urls

import (
	"errors"
	"testing"
)

func TestValidateURLPayload(t *testing.T) {
	cases := []struct {
		name    string
		payload createURLPayload
		wantErr error
	}{
		{
			name: "valid payload",
			payload: createURLPayload{
				ShortCode:   "abc123",
				Destination: "https://example.com",
			},
			wantErr: nil,
		},
		{
			name: "missing short code",
			payload: createURLPayload{
				ShortCode:   "",
				Destination: "https://example.com",
			},
			wantErr: ErrShortCodeRequired,
		},
		{
			name: "missing destination",
			payload: createURLPayload{
				ShortCode:   "abc123",
				Destination: "",
			},
			wantErr: ErrDestinationRequired,
		},
		{
			name: "missing both",
			payload: createURLPayload{
				ShortCode:   "",
				Destination: "",
			},
			wantErr: ErrShortCodeRequired,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := validateURLPayload(c.payload)
			if !errors.Is(err, c.wantErr) {
				t.Errorf("ValidateURLPayload() error = %v, wantErr = %v", err, c.wantErr)
			}
		})
	}
}
