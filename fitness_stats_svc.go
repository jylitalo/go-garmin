package garmin

import (
	"net/url"
	"time"
)

type FitnessStatsService service

// GET https://connect.garmin.com/fitnessstats-service/activity/availableMetrics?startDate=1970-01-01&endDate=2024-08-16&activityType=running&activityType=cycling&activityType=swimming&activityType=fitness_equipment&activityType=other&activityType=winter_sports&activityType=para_sports&_=1723851082967
// GET https://connect.garmin.com/fitnessstats-service/activity?aggregation=daily&userFirstDay=sunday&startDate=2024-08-10&endDate=2024-08-16&groupByActivityType=true&metric=duration&_=1723851082977
// GET https://connect.garmin.com/fitnessstats-service/activity/availableMetrics?startDate=1970-01-01&endDate=2024-08-16&activityType=running&activityType=cycling&activityType=swimming&activityType=fitness_equipment&activityType=other&activityType=winter_sports&activityType=para_sports&_=1723851082967
// GET https://connect.garmin.com/fitnessstats-service/activity/availableMetrics?startDate=2023-08-16&endDate=2024-08-16&_=1723851082974
// GET https://connect.garmin.com/fitnessstats-service/activity?aggregation=daily&userFirstDay=sunday&startDate=2024-08-10&endDate=2024-08-16&groupByActivityType=false&activityType=running&metric=maxHr&_=1723851083009

func (fs *FitnessStatsService) AvailableMetrics(activities []string) (map[string][]string, error) {
	now := fs.c.Clock.Now()
	start := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	out := make(map[string][]string)
	return out, fs.c.apiGet(&out, "/fitnessstats-service/activity/availableMetrics", url.Values{
		"startDate":    []string{start.Format(time.DateOnly)},
		"endDate":      []string{now.Format(time.DateOnly)},
		"activityType": activities,
	})
}

func (fs *FitnessStatsService) Activity(metric, activityType string, start, end time.Time) ([]map[string]any, error) {
	// GET https://connect.garmin.com/fitnessstats-service/activity?aggregation=daily&userFirstDay=sunday&startDate=2024-08-10&endDate=2024-08-16&groupByActivityType=true&metric=<metric>
	res := make([]map[string]any, 0)
	return res, fs.c.apiGet(&res, "/fitnessstats-service/activity", url.Values{
		"aggregation":         []string{"daily"},
		"userFirstDay":        []string{"sunday"},
		"startDate":           []string{start.Format(time.DateOnly)},
		"endDate":             []string{end.Format(time.DateOnly)},
		"groupByActivityType": []string{"true"},
		// duration
		// elevationGain
		// maxhr
		"metric": []string{metric},
		// "metric":              []string{"duration"},
	})
}
