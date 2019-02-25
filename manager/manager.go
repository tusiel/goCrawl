package manager

import (
	"sync"
)

// LinkManager is the overall container
type LinkManager struct {
	sync.WaitGroup
	ProcessedLinks ProcessedLinks
	LinkReport     LinkReport
}

// ProcessedLinks will hold all of the links that have been processed
// to prevent double processing and a potential never-ending loop
type ProcessedLinks struct {
	sync.RWMutex
	Links map[string]bool
}

// LinkReport will hold each page that has been scraped, all the links
// that page contains and the number of times it appears
type LinkReport struct {
	sync.RWMutex
	Links map[Report]uint32
}

// Report is a key for each page scraped and the links found on the page
type Report struct {
	ScrapedLink string
	FoundLink   string
}

// NewLinkManager returns a new instance of LinkManager
func NewLinkManager() *LinkManager {
	return &LinkManager{
		ProcessedLinks: ProcessedLinks{
			Links: make(map[string]bool),
		},
		LinkReport: LinkReport{
			Links: make(map[Report]uint32),
		},
	}
}

// SetLinkProcessed will add the link to the ProcessedLinks map
// with a boolean value of true
func (lm *LinkManager) SetLinkProcessed(link string) {
	lm.ProcessedLinks.Lock()
	defer lm.ProcessedLinks.Unlock()

	lm.ProcessedLinks.Links[link] = true
}

// IsProcessed returns a boolean indicating whether the
// link has already been processed
func (lm *LinkManager) IsProcessed(link string) bool {
	lm.ProcessedLinks.RLock()
	defer lm.ProcessedLinks.RUnlock()

	return lm.ProcessedLinks.Links[link]
}

// SetReportCount will set the count of the passed report
// in the LinkReport object
func (lm *LinkManager) SetReportCount(scrapedLink string, foundLink string) {
	r := Report{ScrapedLink: scrapedLink, FoundLink: foundLink}

	lm.LinkReport.Lock()
	defer lm.LinkReport.Unlock()

	_, ok := lm.LinkReport.Links[r]
	if !ok {
		lm.LinkReport.Links[r] = 0
	}

	lm.LinkReport.Links[r]++
}

// GetReport will return a map that consists of the scapedLink,
// the links found on that page, and how many times the link was found
// on the page
func (lm *LinkManager) GetReport() map[string]map[string]uint32 {
	report := make(map[string]map[string]uint32)

	for key, value := range lm.LinkReport.Links {
		_, ok := report[key.ScrapedLink]
		if !ok {
			report[key.ScrapedLink] = make(map[string]uint32)
		}

		_, ok = report[key.ScrapedLink][key.FoundLink]
		if !ok {
			report[key.ScrapedLink][key.FoundLink] = 0
		}

		report[key.ScrapedLink][key.FoundLink] += value
	}

	return report
}
