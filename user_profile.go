package garmin

import (
	"fmt"
	"net/http"
	"net/url"
)

type UserProfileService service

type UserProfileBase struct {
	UserProfilePk                int    `json:"userProfilePk"`
	UserName                     string `json:"userName"`
	FirstName                    string `json:"firstName"`
	LastName                     string `json:"lastName"`
	BirthDate                    string `json:"birthDate"`
	Gender                       string `json:"gender"`
	EmailAddress                 string `json:"emailAddress"`
	CreateDate                   string `json:"createDate"`
	MeasurementSystemPk          int    `json:"measurementSystemPk"`
	GlucoseMeasurementUnitID     int    `json:"glucoseMeasurementUnitId"`
	HydrationMeasurementUnitID   int    `json:"hydrationMeasurementUnitId"`
	TimeZonePk                   int    `json:"timeZonePk"`
	DecimalFormat                int    `json:"decimalFormat"`
	TimeFormat                   int    `json:"timeFormat"`
	FormatLocalePk               int    `json:"formatLocalePk"`
	DayOfWeekPk                  int    `json:"dayOfWeekPk"`
	GarminGlobalID               string `json:"garminGlobalId"`
	DisplayName                  string `json:"displayName"`
	TocAcceptedDate              string `json:"tocAcceptedDate"`
	AccessDeletedDate            any    `json:"accessDeletedDate"`
	GarminGUID                   string `json:"garminGUID"`
	CountryCode                  string `json:"countryCode"`
	CountryCodeVerified          bool   `json:"countryCodeVerified"`
	CountryCodeVerifiedTimestamp string `json:"countryCodeVerifiedTimestamp"`
	GolfDistanceUnitID           int    `json:"golfDistanceUnitId"`
	GolfElevationUnitID          any    `json:"golfElevationUnitId"`
	GolfSpeedUnitID              any    `json:"golfSpeedUnitId"`
}

func (up *UserProfileService) UserProfileBase() (*UserProfileBase, error) {
	var upb UserProfileBase
	return &upb, up.c.apiGet(&upb, "/userprofile-service/userprofile/userProfileBase", nil)
}

type UserSettings struct {
	ID               int                        `json:"id"`
	UserData         UserSettingsUserData       `json:"userData"`
	UserSleep        UserSettingsSleep          `json:"userSleep"`
	ConnectDate      any                        `json:"connectDate"`
	SourceType       any                        `json:"sourceType"`
	UserSleepWindows []UserSettingsSleepWindows `json:"userSleepWindows"`
}

type UserSettingsSleep struct {
	SleepTime        int  `json:"sleepTime"`
	DefaultSleepTime bool `json:"defaultSleepTime"`
	WakeTime         int  `json:"wakeTime"`
	DefaultWakeTime  bool `json:"defaultWakeTime"`
}

type UserSettingsSleepWindows struct {
	SleepWindowFrequency              string `json:"sleepWindowFrequency"`
	StartSleepTimeSecondsFromMidnight int    `json:"startSleepTimeSecondsFromMidnight"`
	EndSleepTimeSecondsFromMidnight   int    `json:"endSleepTimeSecondsFromMidnight"`
}

type UserSettingsUserData struct {
	Gender string `json:"gender"`
	// Weight in grams.
	Weight float64 `json:"weight"`
	// Height in centimeters.
	Height     float64 `json:"height"`
	TimeFormat string  `json:"timeFormat"`
	// BirthDate in YYYY-MM-DD format.
	BirthDate         string             `json:"birthDate"`
	MeasurementSystem string             `json:"measurementSystem"`
	ActivityLevel     any                `json:"activityLevel"`
	Handedness        string             `json:"handedness"`
	PowerFormat       UserSettingsFormat `json:"powerFormat"`
	HeartRateFormat   UserSettingsFormat `json:"heartRateFormat"`
	FirstDayOfWeek    struct {
		DayID              int    `json:"dayId"`
		DayName            string `json:"dayName"`
		SortOrder          int    `json:"sortOrder"`
		IsPossibleFirstDay bool   `json:"isPossibleFirstDay"`
	} `json:"firstDayOfWeek"`
	Vo2MaxRunning                  float64 `json:"vo2MaxRunning"`
	Vo2MaxCycling                  any     `json:"vo2MaxCycling"`
	LactateThresholdSpeed          any     `json:"lactateThresholdSpeed"`
	LactateThresholdHeartRate      any     `json:"lactateThresholdHeartRate"`
	DiveNumber                     any     `json:"diveNumber"`
	IntensityMinutesCalcMethod     string  `json:"intensityMinutesCalcMethod"`
	ModerateIntensityMinutesHrZone int     `json:"moderateIntensityMinutesHrZone"`
	VigorousIntensityMinutesHrZone int     `json:"vigorousIntensityMinutesHrZone"`
	HydrationMeasurementUnit       string  `json:"hydrationMeasurementUnit"`
	HydrationContainers            []struct {
		Name   any    `json:"name"`
		Volume int    `json:"volume"`
		Unit   string `json:"unit"`
	} `json:"hydrationContainers"`
	HydrationAutoGoalEnabled       bool `json:"hydrationAutoGoalEnabled"`
	FirstbeatMaxStressScore        any  `json:"firstbeatMaxStressScore"`
	FirstbeatCyclingLtTimestamp    any  `json:"firstbeatCyclingLtTimestamp"`
	FirstbeatRunningLtTimestamp    any  `json:"firstbeatRunningLtTimestamp"`
	ThresholdHeartRateAutoDetected any  `json:"thresholdHeartRateAutoDetected"`
	FtpAutoDetected                any  `json:"ftpAutoDetected"`
	TrainingStatusPausedDate       any  `json:"trainingStatusPausedDate"`
	WeatherLocation                struct {
		UseFixedLocation any `json:"useFixedLocation"`
		Latitude         any `json:"latitude"`
		Longitude        any `json:"longitude"`
		LocationName     any `json:"locationName"`
		IsoCountryCode   any `json:"isoCountryCode"`
		PostalCode       any `json:"postalCode"`
	} `json:"weatherLocation"`
	GolfDistanceUnit          string   `json:"golfDistanceUnit"`
	GolfElevationUnit         any      `json:"golfElevationUnit"`
	GolfSpeedUnit             any      `json:"golfSpeedUnit"`
	ExternalBottomTime        any      `json:"externalBottomTime"`
	AvailableTrainingDays     []string `json:"availableTrainingDays"`
	PreferredLongTrainingDays []string `json:"preferredLongTrainingDays"`
}

type UserSettingsFormat struct {
	FormatID      int    `json:"formatId"`
	FormatKey     string `json:"formatKey"`
	MinFraction   int    `json:"minFraction"`
	MaxFraction   int    `json:"maxFraction"`
	GroupingUsed  bool   `json:"groupingUsed"`
	DisplayFormat any    `json:"displayFormat"`
}

func (up *UserProfileService) UserSettings() (*UserSettings, error) {
	var us UserSettings
	return &us, up.c.apiGet(&us, "/userprofile-service/userprofile/user-settings", nil)
}

type PersonalInformation struct {
	UserInfo struct {
		BirthDate   string `json:"birthDate"`
		GenderType  string `json:"genderType"`
		Email       string `json:"email"`
		Locale      string `json:"locale"`
		TimeZone    string `json:"timeZone"`
		Age         int    `json:"age"`
		CountryCode string `json:"countryCode"`
	} `json:"userInfo"`
	BiometricProfile struct {
		UserID                    int     `json:"userId"`
		Height                    float64 `json:"height"`
		Weight                    float64 `json:"weight"`
		Vo2Max                    float64 `json:"vo2Max"`
		Vo2MaxCycling             any     `json:"vo2MaxCycling"`
		LactateThresholdHeartRate any     `json:"lactateThresholdHeartRate"`
		ActivityClass             any     `json:"activityClass"`
		FunctionalThresholdPower  any     `json:"functionalThresholdPower"`
		CriticalSwimSpeed         any     `json:"criticalSwimSpeed"`
	} `json:"biometricProfile"`
	TimeZone  string `json:"timeZone"`
	Locale    string `json:"locale"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

func (up *UserProfileService) PersonalInformation(userUUID string) (*PersonalInformation, error) {
	var pi PersonalInformation
	p := fmt.Sprintf("/userprofile-service/userprofile/personal-information/%s", userUUID)
	return &pi, up.c.apiGet(&pi, p, nil)
}

type SocialProfile struct {
	ID                            int64    `json:"id"`
	ProfileID                     int64    `json:"profileId"`
	GarminGUID                    string   `json:"garminGUID"`
	DisplayName                   string   `json:"displayName"`
	FullName                      string   `json:"fullName"`
	UserName                      string   `json:"userName"`
	ProfileImageType              string   `json:"profileImageType"`
	ProfileImageURLLarge          string   `json:"profileImageUrlLarge"`
	ProfileImageURLMedium         string   `json:"profileImageUrlMedium"`
	ProfileImageURLSmall          string   `json:"profileImageUrlSmall"`
	Location                      string   `json:"location"`
	FacebookURL                   string   `json:"facebookUrl"`
	TwitterURL                    string   `json:"twitterUrl"`
	PersonalWebsite               *string  `json:"personalWebsite"`
	Motivation                    *string  `json:"motivation"`
	Bio                           *string  `json:"bio"`
	PrimaryActivity               *string  `json:"primaryActivity"`
	FavoriteActivityTypes         []any    `json:"favoriteActivityTypes"`
	RunningTrainingSpeed          float64  `json:"runningTrainingSpeed"`
	CyclingTrainingSpeed          float64  `json:"cyclingTrainingSpeed"`
	FavoriteCyclingActivityTypes  []any    `json:"favoriteCyclingActivityTypes"`
	CyclingClassification         any      `json:"cyclingClassification"`
	CyclingMaxAvgPower            float64  `json:"cyclingMaxAvgPower"`
	SwimmingTrainingSpeed         float64  `json:"swimmingTrainingSpeed"`
	ProfileVisibility             string   `json:"profileVisibility"`
	ActivityStartVisibility       string   `json:"activityStartVisibility"`
	ActivityMapVisibility         string   `json:"activityMapVisibility"`
	CourseVisibility              string   `json:"courseVisibility"`
	ActivityHeartRateVisibility   string   `json:"activityHeartRateVisibility"`
	ActivityPowerVisibility       string   `json:"activityPowerVisibility"`
	BadgeVisibility               string   `json:"badgeVisibility"`
	ShowAge                       bool     `json:"showAge"`
	ShowWeight                    bool     `json:"showWeight"`
	ShowHeight                    bool     `json:"showHeight"`
	ShowWeightClass               bool     `json:"showWeightClass"`
	ShowAgeRange                  bool     `json:"showAgeRange"`
	ShowGender                    bool     `json:"showGender"`
	ShowActivityClass             bool     `json:"showActivityClass"`
	ShowVO2Max                    bool     `json:"showVO2Max"`
	ShowPersonalRecords           bool     `json:"showPersonalRecords"`
	ShowLast12Months              bool     `json:"showLast12Months"`
	ShowLifetimeTotals            bool     `json:"showLifetimeTotals"`
	ShowUpcomingEvents            bool     `json:"showUpcomingEvents"`
	ShowRecentFavorites           bool     `json:"showRecentFavorites"`
	ShowRecentDevice              bool     `json:"showRecentDevice"`
	ShowRecentGear                bool     `json:"showRecentGear"`
	ShowBadges                    bool     `json:"showBadges"`
	OtherActivity                 *string  `json:"otherActivity"`
	OtherPrimaryActivity          *string  `json:"otherPrimaryActivity"`
	OtherMotivation               *string  `json:"otherMotivation"`
	UserRoles                     []string `json:"userRoles"`
	NameApproved                  bool     `json:"nameApproved"`
	UserProfileFullName           string   `json:"userProfileFullName"`
	MakeGolfScorecardsPrivate     bool     `json:"makeGolfScorecardsPrivate"`
	AllowGolfLiveScoring          bool     `json:"allowGolfLiveScoring"`
	AllowGolfScoringByConnections bool     `json:"allowGolfScoringByConnections"`
	UserLevel                     int      `json:"userLevel"`
	UserPoint                     int      `json:"userPoint"`
	LevelUpdateDate               string   `json:"levelUpdateDate"`
	LevelIsViewed                 bool     `json:"levelIsViewed"`
	LevelPointThreshold           int      `json:"levelPointThreshold"`
	UserPointOffset               int      `json:"userPointOffset"`
	UserPro                       bool     `json:"userPro"`
}

// SocialProfile will return the user social profile given the user's
// 'displayName' (the displayName is a UUID).
func (up *UserProfileService) SocialProfile(displayName string) (*SocialProfile, error) {
	var updated SocialProfile
	// TODO is it possible to exclude the displayName UUID???
	p := fmt.Sprintf("/userprofile-service/socialProfile/%s", displayName)
	return &updated, up.c.apiGet(&updated, p, nil)
}

type PublicSocialProfile struct {
	ID                          int    `json:"id"`
	ProfileID                   int    `json:"profileId"`
	DisplayName                 string `json:"displayName"`
	ImageType                   string `json:"profileImageType"`
	ImageURLLarge               string `json:"profileImageUrlLarge"`
	ImageURLMedium              string `json:"profileImageUrlMedium"`
	ImageURLSmall               string `json:"profileImageUrlSmall"`
	ProfileVisibility           string `json:"profileVisibility"`
	ActivityHeartRateVisibility string `json:"activityHeartRateVisibility"`
	ActivityPowerVisibility     string `json:"activityPowerVisibility"`
	FullName                    string `json:"fullName"`
	Location                    any    `json:"location"`
	UserLevel                   int    `json:"userLevel"`
	UserPoint                   int    `json:"userPoint"`
	LevelUpdateDate             string `json:"levelUpdateDate"`
	LevelIsViewed               bool   `json:"levelIsViewed"`
	LevelPointThreshold         int    `json:"levelPointThreshold"`
	IsBlocked                   bool   `json:"isBlocked"`
}

func (up *UserProfileService) PublicSocialProfile(displayName string) (*PublicSocialProfile, error) {
	var psp PublicSocialProfile
	p := fmt.Sprintf("/userprofile-service/socialProfile/public/%s", displayName)
	return &psp, up.c.apiGet(&psp, p, nil)
}

type ProfileStatus struct {
	UserProfileID         int      `json:"userProfileId"`
	DisplayName           string   `json:"displayName"`
	ConnectionCount       int      `json:"connectionCount"`
	ConnectionRequestID   any      `json:"connectionRequestId"`
	ConnectionRequestorID any      `json:"connectionRequestorId"`
	UserConnectionStatus  int      `json:"userConnectionStatus"`
	FollowerCount         int      `json:"followerCount"`
	UserRoles             []string `json:"userRoles"`
	UserPro               bool     `json:"userPro"`
}

func (up *UserProfileService) ProfileStatus(displayName string) (*ProfileStatus, error) {
	var ps ProfileStatus
	p := fmt.Sprintf("/userprofile-service/connection/profileStatus/%s", displayName)
	return &ps, up.c.apiGet(&ps, p, url.Values{
		"displayMutedStatus": []string{"true"},
	})
}

// SocialProfile will update the user social profile given the user's
// 'displayName' (the displayName is a UUID).
func (up *UserProfileService) UpdateSocialProfile(displayName string, profile *SocialProfile) (*SocialProfile, error) {
	// PUT https://connect.garmin.com/userprofile-service/socialProfile/<user_uuid>
	var updated SocialProfile
	status, err := up.c.api(
		&updated,
		"PUT",
		fmt.Sprintf("/userprofile-service/socialProfile/%s", displayName),
		nil,
		profile,
	)
	if status != http.StatusOK {
		return nil, fmt.Errorf("bad status code %d", status)
	}
	return &updated, err
}

// UserSettingsUpdate is the payload sent in order to update the user's
// settings. It is similar to the UserSettings struct except that all of its
// fields are nilable so that only a limited number of fields are updated.
type UserSettingsUpdate struct {
	UserData *UserDataUpdate `json:"userData,omitempty"`
}

func (usu *UserSettingsUpdate) Weight(w float64) *UserSettingsUpdate {
	if usu.UserData == nil {
		usu.UserData = new(UserDataUpdate)
	}
	usu.UserData.Weight = &w
	return usu
}

func (usu *UserSettingsUpdate) Height(h float64) *UserSettingsUpdate {
	if usu.UserData == nil {
		usu.UserData = new(UserDataUpdate)
	}
	usu.UserData.Height = &h
	return usu
}

func (usu *UserSettingsUpdate) Gender(g string) *UserSettingsUpdate {
	if usu.UserData == nil {
		usu.UserData = new(UserDataUpdate)
	}
	usu.UserData.Gender = &g
	return usu
}

// LeftHanded sets the request to update user settings to be left handed.
func (usu *UserSettingsUpdate) LeftHanded() *UserSettingsUpdate { return usu.hand("LEFT") }

// RightHanded sets the request to update user settings to be right handed.
func (usu *UserSettingsUpdate) RightHanded() *UserSettingsUpdate { return usu.hand("RIGHT") }

func (usu *UserSettingsUpdate) hand(hand string) *UserSettingsUpdate {
	if usu.UserData == nil {
		usu.UserData = new(UserDataUpdate)
	}
	usu.UserData.Handedness = &hand
	return usu
}

type UserDataUpdate struct {
	Gender *string `json:"gender,omitempty"`
	// Weight in grams.
	Weight *float64 `json:"weight,omitempty"`
	// Height in centimeters.
	Height     *float64 `json:"height,omitempty"`
	Handedness *string  `json:"handedness,omitempty"`
}

// UpdateSettings will send a partial user settings object with only the fields
// that should be changed.
//
//	func TestUpdates() {
//	    req := new(garmin.UserSettingsUpdate).Weight(79786.8328)
//	    err := api.UserProfile.UpdateSettings(req)
//	    if err != nil {
//	        panic(err)
//	    }
//	}
func (up *UserProfileService) UpdateSettings(usu *UserSettingsUpdate) error {
	// Entering 175.8 lbs triggers this request (converted to grams):
	//
	// PUT https://connect.garmin.com/userprofile-service/userprofile/user-settings/
	//
	// {"userData":{"weight":79741.4736}}
	//
	// Or to update both weight (g) and height (cm), use this payload:
	//  {"userData":{"weight":79786.8328,"height":182.87999972202238}}
	_, err := up.c.api(nil, "PUT", "/userprofile-service/userprofile/user-settings", nil, usu)
	return err
}

type PulseOxCapable struct {
	CapableName   string `json:"capableName"`
	CapableEnable bool   `json:"capableEnable"`
}

func (up *UserProfileService) PulseOxCapable() (*PulseOxCapable, error) {
	var po PulseOxCapable
	return &po, up.c.apiGet(&po, "/userprofile-service/userprofile/capableEnable/pulseOxCapable", nil)
}

type SegmentLeaderboard struct {
	UserOptionalFeaturePK string `json:"userOptionalFeaturePK"`
	UserProfilePK         string `json:"userProfilePK"`
	OptionalFeatureType   string `json:"optionalFeatureType"`
	UserOptionType        string `json:"userOptionType"`
	CreateDate            string `json:"createDate"`
	UpdateDate            string `json:"updateDate"`
}

func (up *UserProfileService) SegmentLeaderboard() (*SegmentLeaderboard, error) {
	var sl SegmentLeaderboard
	return &sl, up.c.apiGet(&sl, "/userprofile-service/userprofile/optional-feature/segment-leaderboard", nil)
}

func (up *UserProfileService) StravaSegments() (*SegmentLeaderboard, error) {
	var sl SegmentLeaderboard
	return &sl, up.c.apiGet(&sl, "/userprofile-service/userprofile/optional-feature/strava-segments", nil)
}

type Settings struct {
	DisplayName       string `json:"displayName"`
	PreferredLocale   string `json:"preferredLocale"`
	MeasurementSystem string `json:"measurementSystem"`
	FirstDayOfWeek    struct {
		DayID              int    `json:"dayId"`
		DayName            string `json:"dayName"`
		SortOrder          int    `json:"sortOrder"`
		IsPossibleFirstDay bool   `json:"isPossibleFirstDay"`
	} `json:"firstDayOfWeek"`
	NumberFormat string `json:"numberFormat"`
	TimeFormat   struct {
		FormatID      int    `json:"formatId"`
		FormatKey     string `json:"formatKey"`
		MinFraction   int    `json:"minFraction"`
		MaxFraction   int    `json:"maxFraction"`
		GroupingUsed  bool   `json:"groupingUsed"`
		DisplayFormat string `json:"displayFormat"`
	} `json:"timeFormat"`
	DateFormat struct {
		FormatID      int    `json:"formatId"`
		FormatKey     string `json:"formatKey"`
		MinFraction   int    `json:"minFraction"`
		MaxFraction   int    `json:"maxFraction"`
		GroupingUsed  bool   `json:"groupingUsed"`
		DisplayFormat string `json:"displayFormat"`
	} `json:"dateFormat"`
	PowerFormat struct {
		FormatID      int    `json:"formatId"`
		FormatKey     string `json:"formatKey"`
		MinFraction   int    `json:"minFraction"`
		MaxFraction   int    `json:"maxFraction"`
		GroupingUsed  bool   `json:"groupingUsed"`
		DisplayFormat any    `json:"displayFormat"`
	} `json:"powerFormat"`
	HeartRateFormat struct {
		FormatID      int    `json:"formatId"`
		FormatKey     string `json:"formatKey"`
		MinFraction   int    `json:"minFraction"`
		MaxFraction   int    `json:"maxFraction"`
		GroupingUsed  bool   `json:"groupingUsed"`
		DisplayFormat any    `json:"displayFormat"`
	} `json:"heartRateFormat"`
	TimeZone                 string `json:"timeZone"`
	HydrationMeasurementUnit string `json:"hydrationMeasurementUnit"`
	HydrationContainers      []struct {
		Name   any    `json:"name"`
		Volume int    `json:"volume"`
		Unit   string `json:"unit"`
	} `json:"hydrationContainers"`
	GolfDistanceUnit          string `json:"golfDistanceUnit"`
	GolfElevationUnit         any    `json:"golfElevationUnit"`
	GolfSpeedUnit             any    `json:"golfSpeedUnit"`
	AvailableTrainingDays     any    `json:"availableTrainingDays"`
	PreferredLongTrainingDays any    `json:"preferredLongTrainingDays"`
}

func (up *UserProfileService) Settings() (*Settings, error) {
	var s Settings
	return &s, up.c.apiGet(&s, "/userprofile-service/userprofile/settings", nil)
}
