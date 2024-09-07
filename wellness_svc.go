package garmin

import (
	"fmt"
	"net/url"
	"time"
)

type WellnessService service

type DailyHeartRate struct {
	UserProfilePK                    int    `json:"userProfilePK"`
	CalendarDate                     string `json:"calendarDate"`
	StartTimestampGMT                string `json:"startTimestampGMT"`
	EndTimestampGMT                  string `json:"endTimestampGMT"`
	StartTimestampLocal              string `json:"startTimestampLocal"`
	EndTimestampLocal                string `json:"endTimestampLocal"`
	MaxHeartRate                     int    `json:"maxHeartRate"`
	MinHeartRate                     int    `json:"minHeartRate"`
	RestingHeartRate                 int    `json:"restingHeartRate"`
	LastSevenDaysAvgRestingHeartRate int    `json:"lastSevenDaysAvgRestingHeartRate"`
	HeartRateValueDescriptors        []struct {
		Key   string `json:"key"`
		Index int    `json:"index"`
	} `json:"heartRateValueDescriptors"`
	// HeartRateValues is an array of integer pairs that look like [timestamp, heartrate]
	HeartRateValues [][2]uint64 `json:"heartRateValues"`
}

func (ws *WellnessService) DailyHeartRate(date time.Time) (hr *DailyHeartRate, err error) {
	hr = new(DailyHeartRate)
	return hr, ws.c.apiGet(
		hr,
		"/wellness-service/wellness/dailyHeartRate",
		url.Values{"date": []string{date.Format(time.DateOnly)}},
	)
}

func (ws *WellnessService) DailySleep(userUUID string, date time.Time) (*DailySleep, error) {
	var sd DailySleep
	p := fmt.Sprintf("/wellness-service/wellness/dailySleepData/%s", userUUID)
	return &sd, ws.c.apiGet(&sd, p, url.Values{"date": []string{date.Format(time.DateOnly)}})
}

type DailyStress struct {
	UserProfilePK                 int    `json:"userProfilePK"`
	CalendarDate                  string `json:"calendarDate"`
	StartTimestampGMT             string `json:"startTimestampGMT"`
	EndTimestampGMT               string `json:"endTimestampGMT"`
	StartTimestampLocal           string `json:"startTimestampLocal"`
	EndTimestampLocal             string `json:"endTimestampLocal"`
	MaxStressLevel                int    `json:"maxStressLevel"`
	AvgStressLevel                int    `json:"avgStressLevel"`
	StressChartValueOffset        int    `json:"stressChartValueOffset"`
	StressChartYAxisOrigin        int    `json:"stressChartYAxisOrigin"`
	StressValueDescriptorsDTOList []struct {
		Key   string `json:"key"`
		Index int    `json:"index"`
	} `json:"stressValueDescriptorsDTOList"`
	StressValuesArray                  [][]any `json:"stressValuesArray"`
	BodyBatteryValueDescriptorsDTOList []struct {
		BodyBatteryValueDescriptorIndex int    `json:"bodyBatteryValueDescriptorIndex"`
		BodyBatteryValueDescriptorKey   string `json:"bodyBatteryValueDescriptorKey"`
	} `json:"bodyBatteryValueDescriptorsDTOList"`
	BodyBatteryValuesArray [][]any `json:"bodyBatteryValuesArray"`
}

func (w *WellnessService) DailyStress(date time.Time) (*DailyStress, error) {
	var ds DailyStress
	p := fmt.Sprintf("/wellness-service/wellness/dailyStress/%s", date.Format(time.DateOnly))
	return &ds, w.c.apiGet(&ds, p, nil)
}

type BodyBatteryMessagingToday struct {
	UserProfilePK              int     `json:"userProfilePK"`
	CalendarDate               string  `json:"calendarDate"`
	StartTimestampGMT          string  `json:"startTimestampGMT"`
	EndTimestampGMT            string  `json:"endTimestampGMT"`
	StartTimestampLocal        string  `json:"startTimestampLocal"`
	EndTimestampLocal          string  `json:"endTimestampLocal"`
	DeltaValue                 int     `json:"deltaValue"`
	TimeOfDay                  string  `json:"timeOfDay"`
	ConfirmedTotalSleepSeconds int     `json:"confirmedTotalSleepSeconds"`
	BodyBatteryVersion         float64 `json:"bodyBatteryVersion"`
}

func (w *WellnessService) BodyBatteryMessagingToday() (*BodyBatteryMessagingToday, error) {
	var bbm BodyBatteryMessagingToday
	return &bbm, w.c.apiGet(&bbm, "/wellness-service/wellness/bodyBattery/messagingToday", nil)
}

type BodyBatteryEvent struct {
	Event struct {
		EventType              string `json:"eventType"`
		EventStartTimeGmt      string `json:"eventStartTimeGmt"`
		TimezoneOffset         int    `json:"timezoneOffset"`
		DurationInMilliseconds int    `json:"durationInMilliseconds"`
		BodyBatteryImpact      int    `json:"bodyBatteryImpact"`
		FeedbackType           string `json:"feedbackType"`
		ShortFeedback          string `json:"shortFeedback"`
	} `json:"event"`
	ActivityName                  any     `json:"activityName"`
	ActivityType                  any     `json:"activityType"`
	ActivityID                    any     `json:"activityId"`
	AverageStress                 float64 `json:"averageStress"`
	StressValueDescriptorsDTOList []struct {
		Key   string `json:"key"`
		Index int    `json:"index"`
	} `json:"stressValueDescriptorsDTOList"`
	StressValuesArray                  [][]any `json:"stressValuesArray"`
	BodyBatteryValueDescriptorsDTOList []struct {
		BodyBatteryValueDescriptorIndex int    `json:"bodyBatteryValueDescriptorIndex"`
		BodyBatteryValueDescriptorKey   string `json:"bodyBatteryValueDescriptorKey"`
	} `json:"bodyBatteryValueDescriptorsDTOList"`
	BodyBatteryValuesArray [][]any `json:"bodyBatteryValuesArray"`
}

func (w *WellnessService) BodyBatteryEvents(date time.Time) (res []BodyBatteryEvent, e error) {
	p := fmt.Sprintf("/wellness-service/wellness/bodyBattery/events/%s", date.Format(time.DateOnly))
	return res, w.c.apiGet(&res, p, nil)
}

func (w *WellnessService) DailyEvents(userUUID string, date time.Time) {
	// GET https://connect.garmin.com/wellness-service/wellness/dailyEvents/<userUUID>?calendarDate=2024-08-16
	panic("not implemented")
}

type DailySummaryChartValue struct {
	StartGMT              string `json:"startGMT"`
	EndGMT                string `json:"endGMT"`
	Steps                 int    `json:"steps"`
	Pushes                int    `json:"pushes"`
	PrimaryActivityLevel  string `json:"primaryActivityLevel"`
	ActivityLevelConstant bool   `json:"activityLevelConstant"`
}

func (w *WellnessService) DailySummaryChart(date time.Time) (res []DailySummaryChartValue, e error) {
	return res, w.c.apiGet(&res, "/wellness-service/wellness/dailySummaryChart", url.Values{
		"date": []string{date.Format(time.DateOnly)},
	})
}

type ConsolidatedWellnessGoal struct {
	DeviceGoal       int    `json:"deviceGoal"`
	UserGoal         int    `json:"userGoal"`
	SyncedToDevice   bool   `json:"syncedToDevice"`
	Effective        string `json:"effective"`
	UserGoalCategory string `json:"userGoalCategory"`
	GoalType         string `json:"goalType"`
}

func (w *WellnessService) StepsGoal(date time.Time) (*ConsolidatedWellnessGoal, error) {
	var cw ConsolidatedWellnessGoal
	p := fmt.Sprintf("/wellness-service/wellness/wellness-goals/consolidated/steps/%s", date.Format(time.DateOnly))
	return &cw, w.c.apiGet(&cw, p, nil)
}

func (w *WellnessService) PushesGoal(date time.Time) (*ConsolidatedWellnessGoal, error) {
	var cw ConsolidatedWellnessGoal
	p := fmt.Sprintf("/wellness-service/wellness/wellness-goals/consolidated/pushes/%s", date.Format(time.DateOnly))
	return &cw, w.c.apiGet(&cw, p, nil)
}

type DailyIntensityMinutes struct {
	UserProfilePK             int    `json:"userProfilePK"`
	CalendarDate              string `json:"calendarDate"`
	StartTimestampGMT         string `json:"startTimestampGMT"`
	EndTimestampGMT           string `json:"endTimestampGMT"`
	StartTimestampLocal       string `json:"startTimestampLocal"`
	EndTimestampLocal         string `json:"endTimestampLocal"`
	WeeklyModerate            int    `json:"weeklyModerate"`
	WeeklyVigorous            int    `json:"weeklyVigorous"`
	WeeklyTotal               int    `json:"weeklyTotal"`
	WeekGoal                  int    `json:"weekGoal"`
	DayOfGoalMet              string `json:"dayOfGoalMet"`
	StartDayMinutes           int    `json:"startDayMinutes"`
	EndDayMinutes             int    `json:"endDayMinutes"`
	ModerateMinutes           int    `json:"moderateMinutes"`
	VigorousMinutes           int    `json:"vigorousMinutes"`
	ImValueDescriptorsDTOList any    `json:"imValueDescriptorsDTOList"`
	ImValuesArray             any    `json:"imValuesArray"`
}

func (w *WellnessService) DailyIntensityMinutes(date time.Time) (*DailyIntensityMinutes, error) {
	var dim DailyIntensityMinutes
	p := fmt.Sprintf("/wellness-service/wellness/daily/im/%s", date.Format(time.DateOnly))
	return &dim, w.c.apiGet(&dim, p, nil)
}

type HourlyIntensityMinutes struct {
	WeeklyGoal int `json:"weeklyGoal"`
	DailyStats []struct {
		CalendarDate      string `json:"calendarDate"`
		StartTimestampGMT string `json:"startTimestampGMT"`
		EndTimestampGMT   string `json:"endTimestampGMT"`
		TimezoneOffset    int    `json:"timezoneOffset"`
		VigorousMinutes   int    `json:"vigorousMinutes"`
		ModerateMinutes   int    `json:"moderateMinutes"`
		ValueDescriptors  any    `json:"valueDescriptors"`
		HourlyValues      any    `json:"hourlyValues"`
	} `json:"dailyStats"`
}

func (w *WellnessService) HourlyIntensityMinutes(days int, end time.Time) (*HourlyIntensityMinutes, error) {
	var him HourlyIntensityMinutes
	p := fmt.Sprintf("/wellness-service/stats/hourly/im/%s/%d", end.Format(time.DateOnly), days)
	return &him, w.c.apiGet(&him, p, nil)
}
