package validator

import (
	"errors"
	"testing"
)

func TestValidMailAddress(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr error
	}{
		{
			name:    "Valid public email",
			email:   "taras@schevchenko.com",
			wantErr: nil,
		},
		{
			name:    "Valid email with local domain name",
			email:   "taras@3.com",
			wantErr: nil,
		},
		{
			name:    "Invalid email with out domain name",
			email:   "taras@.com",
			wantErr: ErrInvalidMailAddress,
		},
		{
			name:    "Invalid email with out pre domain name",
			email:   "@schevchenko.com",
			wantErr: ErrInvalidMailAddress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidMailAddress(tt.email); !errors.Is(err, tt.wantErr) {
				t.Errorf("ValidMailAddress() name = %s got = %#v wantErr =  %#v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestValidURLWithError(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		wantErr error
	}{
		{name: "Valid url", rawURL: "https://www.youtube.com/watch?v=U2chxSjrnvk", wantErr: nil},
		{name: "Invalid empty url", rawURL: "", wantErr: ErrInvalidURL},
		{name: "Invalid url with out scheme", rawURL: "www.youtube", wantErr: ErrInvalidURL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidURL(tt.rawURL); !errors.Is(err, tt.wantErr) {
				t.Errorf("ValidMailAddress() name = %s got = %v wantErr =  %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
