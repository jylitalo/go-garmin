package garmin

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type SleepService service

type DailySleep struct {
	DailySleepDTO                       DailySleepDTO                         `json:"dailySleepDTO"`
	RemSleepData                        bool                                  `json:"remSleepData"`
	SleepMovement                       []DurationLevel                       `json:"sleepMovement"`
	SleepLevels                         []DurationLevel                       `json:"sleepLevels"`
	SleepRestlessMoments                []TimestampedValue                    `json:"sleepRestlessMoments"`
	RestlessMomentsCount                int                                   `json:"restlessMomentsCount"`
	WellnessEpochRespirationDataDTOList []WellnessEpochRespirationDataDTOList `json:"wellnessEpochRespirationDataDTOList"`
	SleepHeartRate                      []TimestampedValue                    `json:"sleepHeartRate"`
	SleepStress                         []TimestampedValue                    `json:"sleepStress"`
	SleepBodyBattery                    []TimestampedValue                    `json:"sleepBodyBattery"`
	HrvData                             []TimestampedValue                    `json:"hrvData"`
	SkinTempDataExists                  bool                                  `json:"skinTempDataExists"`
	AvgOvernightHrv                     float64                               `json:"avgOvernightHrv"`
	HrvStatus                           string                                `json:"hrvStatus"`
	BodyBatteryChange                   int                                   `json:"bodyBatteryChange"`
	RestingHeartRate                    int                                   `json:"restingHeartRate"`
}

type WellnessEpochRespirationDataDTOList struct {
	StartTimeGMT     int64   `json:"startTimeGMT"`
	RespirationValue float64 `json:"respirationValue"`
}

// DurationLevel holds the value of the level and the start/end timestamp
type DurationLevel struct {
	StartGMT      string  `json:"startGMT"`
	EndGMT        string  `json:"endGMT"`
	ActivityLevel float64 `json:"activityLevel"`
}

type TimestampedValue struct {
	Value    int   `json:"value"`
	StartGMT int64 `json:"startGMT"`
}

type DailySleepDTO struct {
	ID                            int64       `json:"id"`
	UserProfilePK                 int         `json:"userProfilePK"`
	CalendarDate                  string      `json:"calendarDate"`
	SleepTimeSeconds              int         `json:"sleepTimeSeconds"`
	NapTimeSeconds                int         `json:"napTimeSeconds"`
	SleepWindowConfirmed          bool        `json:"sleepWindowConfirmed"`
	SleepWindowConfirmationType   string      `json:"sleepWindowConfirmationType"`
	SleepStartTimestampGMT        int64       `json:"sleepStartTimestampGMT"`
	SleepEndTimestampGMT          int64       `json:"sleepEndTimestampGMT"`
	SleepStartTimestampLocal      int64       `json:"sleepStartTimestampLocal"`
	SleepEndTimestampLocal        int64       `json:"sleepEndTimestampLocal"`
	AutoSleepStartTimestampGMT    any         `json:"autoSleepStartTimestampGMT"`
	AutoSleepEndTimestampGMT      any         `json:"autoSleepEndTimestampGMT"`
	SleepQualityTypePK            any         `json:"sleepQualityTypePK"`
	SleepResultTypePK             any         `json:"sleepResultTypePK"`
	UnmeasurableSleepSeconds      int         `json:"unmeasurableSleepSeconds"`
	DeepSleepSeconds              int         `json:"deepSleepSeconds"`
	LightSleepSeconds             int         `json:"lightSleepSeconds"`
	RemSleepSeconds               int         `json:"remSleepSeconds"`
	AwakeSleepSeconds             int         `json:"awakeSleepSeconds"`
	DeviceRemCapable              bool        `json:"deviceRemCapable"`
	Retro                         bool        `json:"retro"`
	SleepFromDevice               bool        `json:"sleepFromDevice"`
	AverageRespirationValue       float64     `json:"averageRespirationValue"`
	LowestRespirationValue        float64     `json:"lowestRespirationValue"`
	HighestRespirationValue       float64     `json:"highestRespirationValue"`
	AwakeCount                    int         `json:"awakeCount"`
	AvgSleepStress                float64     `json:"avgSleepStress"`
	AgeGroup                      string      `json:"ageGroup"`
	SleepScoreFeedback            string      `json:"sleepScoreFeedback"`
	SleepScoreInsight             string      `json:"sleepScoreInsight"`
	SleepScorePersonalizedInsight string      `json:"sleepScorePersonalizedInsight"`
	SleepScores                   SleepScores `json:"sleepScores"`
	SleepVersion                  int         `json:"sleepVersion"`
	SleepNeed                     SleepNeed   `json:"sleepNeed"`
	NextSleepNeed                 SleepNeed   `json:"nextSleepNeed"`
}

type SleepNeed struct {
	UserProfilePk            int    `json:"userProfilePk"`
	CalendarDate             string `json:"calendarDate"`
	DeviceID                 int64  `json:"deviceId"`
	TimestampGmt             string `json:"timestampGmt"`
	Baseline                 int    `json:"baseline"`
	Actual                   int    `json:"actual"`
	Feedback                 string `json:"feedback"`
	TrainingFeedback         string `json:"trainingFeedback"`
	SleepHistoryAdjustment   string `json:"sleepHistoryAdjustment"`
	HrvAdjustment            string `json:"hrvAdjustment"`
	NapAdjustment            string `json:"napAdjustment"`
	DisplayedForTheDay       bool   `json:"displayedForTheDay"`
	PreferredActivityTracker bool   `json:"preferredActivityTracker"`
}

type SleepScores struct {
	TotalDuration   SleepScoreRangedRating `json:"totalDuration"`
	Stress          SleepScoreRangedRating `json:"stress"`
	AwakeCount      SleepScoreRangedRating `json:"awakeCount"`
	Overall         SleepScoreValueRating  `json:"overall"`
	RemPercentage   SleepScorePercentage   `json:"remPercentage"`
	Restlessness    SleepScoreRangedRating `json:"restlessness"`
	LightPercentage SleepScorePercentage   `json:"lightPercentage"`
	DeepPercentage  SleepScorePercentage   `json:"deepPercentage"`
}

type SleepScorePercentage struct {
	Value               int     `json:"value"`
	QualifierKey        string  `json:"qualifierKey"`
	OptimalStart        float64 `json:"optimalStart"`
	OptimalEnd          float64 `json:"optimalEnd"`
	IdealStartInSeconds float64 `json:"idealStartInSeconds"`
	IdealEndInSeconds   float64 `json:"idealEndInSeconds"`
}

type SleepScoreRangedRating struct {
	QualifierKey string  `json:"qualifierKey"`
	OptimalStart float64 `json:"optimalStart"`
	OptimalEnd   float64 `json:"optimalEnd"`
}

type SleepScoreValueRating struct {
	Value        int    `json:"value"`
	QualifierKey string `json:"qualifierKey"`
}

func (ss *SleepService) Daily(date time.Time, nonSleepBufferMinutes int) (*DailySleep, error) {
	var ds DailySleep
	return &ds, ss.c.apiGet(&ds, "/sleep-service/sleep/dailySleepData", url.Values{
		"date":                  []string{date.Format(time.DateOnly)},
		"nonSleepBufferMinutes": []string{strconv.FormatInt(int64(nonSleepBufferMinutes), 10)},
	})
}

type DailySleepStats struct {
	OverallStats    DailySleepAverages     `json:"overallStats"`
	IndividualStats []Stat[DailySleepStat] `json:"individualStats"`
}

type DailySleepStat struct {
	RemTime                     int     `json:"remTime"`
	RestingHeartRate            int     `json:"restingHeartRate"`
	TotalSleepTimeInSeconds     int     `json:"totalSleepTimeInSeconds"`
	Respiration                 float64 `json:"respiration"`
	LocalSleepEndTimeInMillis   int64   `json:"localSleepEndTimeInMillis"`
	DeepTime                    int     `json:"deepTime"`
	AwakeTime                   int     `json:"awakeTime"`
	SleepScoreQuality           string  `json:"sleepScoreQuality"`
	SpO2                        any     `json:"spO2"`
	LocalSleepStartTimeInMillis int64   `json:"localSleepStartTimeInMillis"`
	SleepNeed                   int     `json:"sleepNeed"`
	BodyBatteryChange           int     `json:"bodyBatteryChange"`
	GmtSleepStartTimeInMillis   int64   `json:"gmtSleepStartTimeInMillis"`
	GmtSleepEndTimeInMillis     int64   `json:"gmtSleepEndTimeInMillis"`
	HrvStatus                   string  `json:"hrvStatus"`
	SkinTempF                   any     `json:"skinTempF"`
	SleepScore                  int     `json:"sleepScore"`
	SkinTempC                   any     `json:"skinTempC"`
	LightTime                   int     `json:"lightTime"`
	Hrv7DAverage                float64 `json:"hrv7dAverage"`
}

type DailySleepAverages struct {
	SpO2                any     `json:"averageSpO2"`
	LocalSleepStartTime float64 `json:"averageLocalSleepStartTime"`
	Respiration         float64 `json:"averageRespiration"`
	BodyBatteryChange   float64 `json:"averageBodyBatteryChange"`
	SkinTempF           any     `json:"averageSkinTempF"`
	SleepScore          float64 `json:"averageSleepScore"`
	LocalSleepEndTime   float64 `json:"averageLocalSleepEndTime"`
	SkinTempC           any     `json:"averageSkinTempC"`
	SleepSeconds        float64 `json:"averageSleepSeconds"`
	SleepNeed           float64 `json:"averageSleepNeed"`
	RestingHeartRate    float64 `json:"averageRestingHeartRate"`
}

func (ss *SleepService) DailySleepStats(start, end time.Time) (*DailySleepStats, error) {
	var s DailySleepStats
	return &s, ss.c.apiGet(&s, datepath("/sleep-service/stats/sleep/daily", start, end), nil)
}

type WeeklySleepStats struct {
	OverallStats    WeeklySleepAverages         `json:"overallStats"`
	IndividualStats []WeeklySleepIndividualStat `json:"individualStats"`
}

type WeeklySleepIndividualStat struct {
	WeekStartDate string               `json:"weekStartDate"`
	WeekEndDate   string               `json:"weekEndDate"`
	Values        WeeklySleepStatValue `json:"values"`
}

type WeeklySleepStatValue struct {
	RemTime                    float64 `json:"remTime"`
	AwakeTime                  float64 `json:"awakeTime"`
	SleepScoreQuality          string  `json:"sleepScoreQuality"`
	AverageLocalSleepStartTime float64 `json:"averageLocalSleepStartTime"`
	AverageSleepScore          float64 `json:"averageSleepScore"`
	SleepDataDaysCount         int     `json:"sleepDataDaysCount"`
	AverageLocalSleepEndTime   float64 `json:"averageLocalSleepEndTime"`
	AverageSleepSeconds        float64 `json:"averageSleepSeconds"`
	AverageSleepNeed           float64 `json:"averageSleepNeed"`
	LightTime                  float64 `json:"lightTime"`
	DeepTime                   float64 `json:"deepTime"`
}

type WeeklySleepAverages struct {
	LocalSleepStartTime float64 `json:"averageLocalSleepStartTime"`
	LowSleepScore       float64 `json:"averageLowSleepScore"`
	LocalSleepEndTime   float64 `json:"averageLocalSleepEndTime"`
	SleepSeconds        float64 `json:"averageSleepSeconds"`
	SleepNeed           float64 `json:"averageSleepNeed"`
	HighSleepScore      float64 `json:"averageHighSleepScore"`
}

func (ss *SleepService) WeeklySleepStats(weeks int, end time.Time) (*WeeklySleepStats, error) {
	var s WeeklySleepStats
	p := fmt.Sprintf("/sleep-service/stats/sleep/weekly/%s/%d", end.Format(time.DateOnly), weeks)
	return &s, ss.c.apiGet(&s, p, nil)
}
