package garmin

type UserFocusService service

type UserFocus struct {
	UserFocus       string `json:"userFocus"`
	CreateTimestamp string `json:"createTimestamp"`
}

func (uf *UserFocusService) Focus() (*UserFocus, error) {
	var res UserFocus
	return &res, uf.c.apiGet(&res, "/userfocus-service/focus", nil)
}

type SuggestedUserFocus struct {
	UserFocus   string `json:"userFocus"`
	Recommended bool   `json:"recommended"`
}

func (uf *UserFocusService) Suggested() (res []SuggestedUserFocus, e error) {
	return res, uf.c.apiGet(&res, "/userfocus-service/focus/suggestedFocuses", nil)
}

type UserFocusDashboard struct {
	UserFocus struct {
		UserFocus       string `json:"userFocus"`
		CreateTimestamp string `json:"createTimestamp"`
	} `json:"userFocus"`
	Sections []struct {
		SectionType string `json:"sectionType"`
	} `json:"sections"`
	YesterdayStats []struct {
		CardType string `json:"cardType"`
	} `json:"yesterdayStats"`
	LastSevenDayStats []struct {
		CardType string `json:"cardType"`
	} `json:"lastSevenDayStats"`
	DailyStatsConfig struct {
		KeyStats []struct {
			CardType string `json:"cardType"`
		} `json:"keyStats"`
		PrimaryStats []struct {
			CardType string `json:"cardType"`
		} `json:"primaryStats"`
		ActivityTrends []struct {
			CardType string `json:"cardType"`
		} `json:"activityTrends"`
		SecondaryStats []struct {
			CardType string `json:"cardType"`
		} `json:"secondaryStats"`
		SecondaryStatFavorites []struct {
			CardType string `json:"cardType"`
		} `json:"secondaryStatFavorites"`
		SecondaryStatOthers []struct {
			CardType string `json:"cardType"`
		} `json:"secondaryStatOthers"`
		DefaultSecondaryStats bool `json:"defaultSecondaryStats"`
		DefaultPrimaryStats   bool `json:"defaultPrimaryStats"`
	} `json:"dailyStatsConfig"`
}

func (uf *UserFocusService) Dashboard() (*UserFocusDashboard, error) {
	var res UserFocusDashboard
	return &res, uf.c.apiGet(&res, "/userfocus-service/dashboard", nil)
}

type UserFocusAvailablePrimaryStat struct {
	CardType string `json:"cardType"`
}

func (uf *UserFocusService) AvailablePrimaryStats() (res []UserFocusAvailablePrimaryStat, e error) {
	return res, uf.c.apiGet(&res, "/userfocus-service/dashboard/availablePrimaryStats", nil)
}
