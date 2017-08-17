package detect

import (
	"fmt"
	"regexp"
)

// UserAgent for detect user agent data
type UserAgent struct {
	Agent     string
	IsBrowser bool
	IsRobot   bool
	IsMobile  bool
	PlatForm  string
	Browser   string
	Version   string
	Mobile    string
	Robot     string
}

// New is Constructor
func New(agent string) *UserAgent {
	return &UserAgent{
		Agent: agent,
	}
}

func (u *UserAgent) setPlatform() bool {
	for k, v := range PlatForms {
		match := regexp.MustCompile(`(?i)` + k)
		platformMatch := match.FindString(u.Agent)
		if len(platformMatch) > 0 {
			u.PlatForm = v
			return true
		}
	}

	u.PlatForm = "Unknown Platform"
	return false
}
