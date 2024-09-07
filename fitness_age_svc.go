package garmin

import (
	"fmt"
	"time"
)

type FitnessAgeService service

type FitnessAge struct {
	ChronologicalAge     int                  `json:"chronologicalAge"`
	FitnessAge           float64              `json:"fitnessAge"`
	AchievableFitnessAge float64              `json:"achievableFitnessAge"`
	PreviousFitnessAge   float64              `json:"previousFitnessAge"`
	Components           FitnessAgeComponents `json:"components"`
	LastUpdated          string               `json:"lastUpdated"`
}

type FitnessAgeComponents struct {
	VigorousDaysAvg    FitnessAgeVigorousDaysAvg    `json:"vigorousDaysAvg"`
	RHR                FitnessAgeRHR                `json:"rhr"`
	VigorousMinutesAvg FitnessAgeVigorousMinutesAvg `json:"vigorousMinutesAvg"`
	BMI                FitnessAgeBMI                `json:"bmi"`
}

type FitnessAgeVigorousDaysAvg struct {
	Value           float64 `json:"value"`
	TargetValue     int     `json:"targetValue"`
	PotentialAge    float64 `json:"potentialAge"`
	Priority        int     `json:"priority"`
	Stale           bool    `json:"stale"`
	NumOfWeeksForIm int     `json:"numOfWeeksForIm"`
}

type FitnessAgeRHR struct {
	Value int  `json:"value"`
	Stale bool `json:"stale"`
}

type FitnessAgeVigorousMinutesAvg struct {
	Value           float64 `json:"value"`
	Stale           bool    `json:"stale"`
	NumOfWeeksForIm int     `json:"numOfWeeksForIm"`
}

type FitnessAgeBMI struct {
	Value               float64 `json:"value"`
	TargetValue         float64 `json:"targetValue"`
	ImprovementValue    float64 `json:"improvementValue"`
	PotentialAge        float64 `json:"potentialAge"`
	Priority            int     `json:"priority"`
	Stale               bool    `json:"stale"`
	LastMeasurementDate string  `json:"lastMeasurementDate"`
}

func (fas *FitnessAgeService) FitnessAge(date time.Time) (*FitnessAge, error) {
	var fa FitnessAge
	path := fmt.Sprintf("/fitnessage-service/fitnessage/%s", date.Format(time.DateOnly))
	return &fa, fas.c.apiGet(&fa, path, nil)
}

type DailyFitnessAge struct {
	AchievableFitnessAge float64 `json:"achievableFitnessAge"`
	VigorousDaysAvg      float64 `json:"vigorousDaysAvg"`
	FitnessAge           float64 `json:"fitnessAge"`
	RHR                  int     `json:"rhr"`
	BMI                  float64 `json:"bmi"`
}

func (fas *FitnessAgeService) Daily(start, end time.Time) (res []Stat[DailyFitnessAge], e error) {
	p := datepath("/fitnessage-service/stats/daily", start, end)
	return res, fas.c.apiGet(&res, p, nil)
}

type WeeklyFitnessAge struct {
	VigorousDaysAvg float64 `json:"vigorousDaysAvg"`
	FitnessAge      float64 `json:"fitnessAge"`
	RHR             float64 `json:"rhr"`
	BMI             float64 `json:"bmi"`
}

func (fas *FitnessAgeService) Weekly(start time.Time, weeks int) (res []Stat[WeeklyFitnessAge], e error) {
	p := fmt.Sprintf("/fitnessage-service/stats/weekly/%s/%d", start.Format(time.DateOnly), weeks)
	return res, fas.c.apiGet(&res, p, nil)
}
