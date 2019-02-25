package utils

import (
	"testing"
)

func TestIsExternalDomain(t *testing.T) {
	external := IsExternalDomain("https://www.bbc.co.uk")
	internal := IsExternalDomain("/subdomain")

	if !external {
		t.Errorf("Expected link to return true, got %t", external)
	}

	if internal {
		t.Errorf("Expected link to return false, got %t", internal)
	}
}
func TestGetRelativeURL(t *testing.T) {
	relativeURL := GetRelativeURL("/subdomain", "https://maindomain.com")
	expected := "https://maindomain.com/subdomain"

	if relativeURL != expected {
		t.Errorf("Expected URL to %s be but got %s", expected, relativeURL)
	}
}
