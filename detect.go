package detect

import (
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
	for i := 0; i < len(PlatForms); i++ {
		match := regexp.MustCompile(`(?i)` + PlatFormKeys[i])
		platformMatch := match.FindString(u.Agent)
		if len(platformMatch) > 0 {
			u.PlatForm = PlatForms[PlatFormKeys[i]]
			return true
		}
	}

	u.PlatForm = "Unknown Platform"
	return false
}
