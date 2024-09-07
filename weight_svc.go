package garmin

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type WeightService service

type WeightUnit string

const (
	WeightUnitLbs WeightUnit = "lbs"
)

const (
	gramsInAPound = 453.59237
	poundsInAGram = 0.00220462262185
)

func GramsToPounds(g float64) float64 { return g / gramsInAPound }

func PoundsToGrams(lbs float64) float64 { return lbs / poundsInAGram }

func (ws *WeightService) UpdateWeight(weight float64, unit WeightUnit) error {
	const dateFormat = "2006-01-02T15:04:05.99"
	// POST /weight-service/user-weight
	// Authorization: Bearer ...
	//
	// {"dateTimestamp":"2024-08-15T17:18:00.00","gmtTimestamp":"2024-08-16T00:18:00.00","unitKey":"lbs","value":172}
	// {"dateTimestamp":"2024-08-15T18:08:00.00","gmtTimestamp":"2024-08-16T01:08:00.00","unitKey":"lbs","value":172}
	// =========================================================
	//
	// NOTE:
	// Might need "GET /gdprconsent-service/feature/UPLOAD?_=1723767405770"
	// first, I'm not sure...
	type WeightUpdateRequest struct {
		Date  string  `json:"dateTimestamp"`
		GMT   string  `json:"gmtTimestamp"`
		Unit  string  `json:"unitKey"` // usually "lbs"
		Value float64 `json:"value"`
	}
	now := ws.c.Clock.Now()
	payload := WeightUpdateRequest{
		Date:  now.Format(dateFormat),
		GMT:   now.In(time.UTC).Format(dateFormat),
		Unit:  string(unit),
		Value: weight,
	}
	status, err := ws.c.api(
		nil, // output
		"POST",
		"/weight-service/user-weight",
		nil, // url params
		&payload,
	)
	if err != nil {
		return err
	}
	switch status {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("invalid status code %d", status)
	}
}

func (ws *WeightService) First() (*WeighIn, error) {
	var w WeighIn
	err := ws.c.apiGet(&w, "/weight-service/weight/first", nil)
	return &w, err
}

func (ws *WeightService) Latest(date time.Time) (*WeighIn, error) {
	// One of:
	// GET https://connect.garmin.com/weight-service/weight/latest?date=2024-08-16&ignorePriority=true
	// GET https://connect.garmin.com/weight-service/weight/latest?date=2024-08-16T02:17:50.0&ignorePriority=true
	// GET https://connect.garmin.com/weight-service/weight/latest?date=<date>&ignorePriority=true
	const dateFormat = "2006-01-02T15:04:05.99"
	var w WeighIn
	err := ws.c.apiGet(
		&w,
		"/weight-service/weight/latest",
		url.Values{
			"date":           []string{date.Format(dateFormat)},
			"ignorePriority": []string{"true"},
		},
	)
	return &w, err
}

// DeleteWeight calls "DELETE /weight-service/weight/<date>/byversion/<weight_version>"
//
// You can get the version either from the WeightDated object's `Version` field
// or the `SamplePk` field.
func (ws *WeightService) DeleteWeight(date time.Time, version int64) error {
	path := fmt.Sprintf(
		"/weight-service/weight/%s/byversion/%d",
		date.Format(time.DateOnly),
		version,
	)
	status, err := ws.c.api(nil, "DELETE", path, nil, nil)
	if err != nil {
		return err
	}
	switch status {
	case http.StatusOK, http.StatusNoContent, http.StatusAccepted:
		return nil
	default:
		return fmt.Errorf("invalid status code %d", status)
	}
}

type DailyWeightSummary struct {
	SummaryDate        string    `json:"summaryDate"`
	NumOfWeightEntries int       `json:"numOfWeightEntries"`
	MinWeight          float64   `json:"minWeight"`
	MaxWeight          float64   `json:"maxWeight"`
	LatestWeight       WeighIn   `json:"latestWeight"`
	AllWeightMetrics   []WeighIn `json:"allWeightMetrics"`
}

type WeighIn struct {
	// SamplePk is the same as the `Version`. It depends on the endpoint but one
	// of the two fields will be present.
	SamplePk int64  `json:"samplePk"` // not sure what this is, its only included in some responses
	Date     UnixTS `json:"date"`
	// Date           `json:"date"`
	CalendarDate string `json:"calendarDate"`
	// Version is not included in the range endpoint
	Version *int64 `json:"version"`
	// Weight in grams
	Weight float64 `json:"weight"`
	// BMI is weight / height squared
	BMI               *float64 `json:"bmi"`
	BodyFatPercentage *float64 `json:"bodyFat"`
	// BodyWater is body water in in kilograms
	BodyWater float64 `json:"bodyWater"`
	// BoneMass in grams
	BoneMass int `json:"boneMass"`
	// MuscleMass in grams
	MuscleMass     *int     `json:"muscleMass"`
	PhysiqueRating any      `json:"physiqueRating"`
	VisceralFat    any      `json:"visceralFat"`
	MetabolicAge   any      `json:"metabolicAge"`
	SourceType     string   `json:"sourceType"`
	TimestampGMT   int64    `json:"timestampGMT"`
	WeightDelta    *float64 `json:"weightDelta"`
	CaloricIntake  any      `json:"caloricIntake"`
}

func (w *WeighIn) WeightLbs() float64 {
	return GramsToPounds(w.Weight)
}

type WeightAverage struct {
	From  int64 `json:"from"`
	Until int64 `json:"until"`
	WeighIn
}

type WeightRange struct {
	DailyWeightSummaries []DailyWeightSummary `json:"dailyWeightSummaries"`
	TotalAverage         WeightAverage        `json:"totalAverage"`
	PreviousDateWeight   WeighIn              `json:"previousDateWeight"`
	NextDateWeight       WeighIn              `json:"nextDateWeight"`
}

func (ws *WeightService) Range(start, end time.Time) (*WeightRange, error) {
	// GET /weight-service/weight/range/<start>/<end>?includeAll=true
	path := fmt.Sprintf(
		"/weight-service/weight/range/%s/%s",
		start.Format(time.DateOnly),
		end.Format(time.DateOnly),
	)
	var wr WeightRange
	err := ws.c.apiGet(&wr, path, url.Values{"includeAll": []string{"true"}})
	return &wr, err
}

type WeightDayView struct {
	StartDate      string        `json:"startDate"`
	EndDate        string        `json:"endDate"`
	DateWeightList []WeighIn     `json:"dateWeightList"`
	TotalAverage   WeightAverage `json:"totalAverage"`
}

func (ws *WeightService) DayView(date time.Time) (*WeightDayView, error) {
	var w WeightDayView
	path := fmt.Sprintf("/weight-service/weight/dayview/%s", date.Format(time.DateOnly))
	err := ws.c.apiGet(&w, path, nil)
	return &w, err
}
