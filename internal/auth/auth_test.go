package auth

import (
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("should return api key", func(t *testing.T) {
		headers := map[string][]string{
			"Authorization": {"ApiKey 1234567890"},
		}
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if apiKey != "1234567890" {
			t.Errorf("expected api key to be 1234567890, got %s", apiKey)
		}
	})
}
