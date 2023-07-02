package validator

import "testing"

func TestValidMailAddress(t *testing.T) {
	type want struct {
		email string
		valid bool
	}
	tests := []struct {
		name  string
		email string
		want  want
	}{
		{
			name:  "Valid public email",
			email: "taras@schevchenko.com",
			want:  want{email: "taras@schevchenko.com", valid: true},
		},
		{
			name:  "Valid email with local domain name",
			email: "taras@3.com",
			want:  want{email: "taras@3.com", valid: true},
		},
		{
			name:  "Invalid email with out domain name",
			email: "taras@.com",
			want:  want{email: "", valid: false},
		},
		{
			name:  "Invalid email with out pre domain name",
			email: "@schevchenko.com",
			want:  want{email: "", valid: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, ok := ValidMailAddress(tt.email)
			if email != tt.want.email || ok != tt.want.valid {
				t.Errorf("ValidMailAddress() got = %s, %v want =  %v, %v", email, ok, tt.want.email, tt.want.valid)
			}
		})
	}
}

func TestValidURLWithError(t *testing.T) {
}
