package manager

import (
	"testing"
)

func TestSetLinkProcessed(t *testing.T) {
	lm := NewLinkManager()

	lm.SetLinkProcessed("https://example.com/test")

	processed := lm.IsProcessed("https://example.com/test")

	if !processed {
		t.Errorf("Expected IsProcessed to return true, got %t", processed)
	}
}
func TestGetReport(t *testing.T) {
	lm := NewLinkManager()

	lm.SetReportCount("https://example.com/test", "https://example.com/anotherTest")
	lm.SetReportCount("https://example.com/test", "https://example.com/anotherTest")
	lm.SetReportCount("https://example.com/test", "https://example.com/anotherTest")
	lm.SetReportCount("https://example.com/test", "https://example.com/anotherTest")
	lm.SetReportCount("https://example.com/test", "https://example.com/differentTest")
	lm.SetReportCount("https://example.com/test", "https://example.com/differentTest")
	lm.SetReportCount("https://example.com/test", "https://example.com/oneMoreTest")

	report := lm.GetReport()

	anotherTest := report["https://example.com/test"]["https://example.com/anotherTest"]
	differentTest := report["https://example.com/test"]["https://example.com/differentTest"]
	oneMoreTest := report["https://example.com/test"]["https://example.com/oneMoreTest"]

	if anotherTest != 4 {
		t.Errorf("Expected anotherTest to return 4, got %d", anotherTest)
	}

	if differentTest != 2 {
		t.Errorf("Expected differentTest to return 2, got %d", differentTest)
	}

	if oneMoreTest != 1 {
		t.Errorf("Expected oneMoreTest to return 1, got %d", oneMoreTest)
	}

}
