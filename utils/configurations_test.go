package utils

import (
	"testing"
)

func TestInlineAuthConfig(t *testing.T) {
	tests := []struct {
		name              string
		organizationUrl   string
		organizationToken string
		wantErr           bool
	}{
		{
			name:              "Should return invalid URL",
			organizationUrl:   "gju45gn43g",
			organizationToken: "goierngriuognir",
			wantErr:           true,
		},
		{
			name:              "Shuld create file with",
			organizationUrl:   "https://dev.azure.com/sample",
			organizationToken: "goierngriuognir",
			wantErr:           false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InlineAuthConfig(tt.organizationUrl, tt.organizationToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("InlineAuthConfig() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}

}
