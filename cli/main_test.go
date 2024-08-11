package main

import (
	"strings"
	"testing"
)

func Test_generatePassword(t *testing.T) {
	type args struct {
		length       int
		useSymbols   bool
		useNumbers   bool
		useUppercase bool
	}
	tests := []struct {
		name          string
		args          args
		wantLength    int
		wantSymbol    bool
		wantNumber    bool
		wantUppercase bool
	}{
		{
			name:       "Test length 8",
			args:       args{8, false, false, false},
			wantLength: 8,
		},
		{
			name:       "Test length 16 and use symbols",
			args:       args{16, true, false, false},
			wantLength: 16,
			wantSymbol: true,
		},
		{
			name:       "Test length 16 and use numbers",
			args:       args{16, false, true, false},
			wantLength: 16,
			wantNumber: true,
		},
		{
			name:          "Test length 16 and use uppercase",
			args:          args{16, false, false, true},
			wantLength:    16,
			wantUppercase: true,
		},
		{
			name:          "Test length 16 and use all",
			args:          args{16, true, true, true},
			wantLength:    16,
			wantSymbol:    true,
			wantNumber:    true,
			wantUppercase: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePassword(tt.args.length, tt.args.useSymbols, tt.args.useNumbers, tt.args.useUppercase)
			if len(got) != tt.wantLength {
				t.Errorf("generatePassword() length = %v, want %v", len(got), tt.wantLength)
			}
			if tt.wantSymbol {
				if !strings.ContainsAny(got, symbols) {
					t.Errorf("generatePassword() = %v, want symbol", got)
				}
			}
			if tt.wantNumber {
				if !strings.ContainsAny(got, numbers) {
					t.Errorf("generatePassword() = %v, want number", got)
				}
			}
			if tt.wantUppercase {
				if !strings.ContainsAny(got, uppercase) {
					t.Errorf("generatePassword() = %v, want uppercase", got)
				}
			}
		})
	}
}
