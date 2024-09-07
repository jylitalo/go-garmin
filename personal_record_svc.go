package garmin

import "fmt"

type PersonalRecordService service

type PersonalRecord struct {
	ID                                  int64   `json:"id"`
	TypeID                              int     `json:"typeId"`
	ActivityID                          int64   `json:"activityId"`
	ActivityName                        string  `json:"activityName"`
	ActivityType                        string  `json:"activityType"`
	ActivityStartDateTimeInGMT          int64   `json:"activityStartDateTimeInGMT"`
	ActStartDateTimeInGMTFormatted      string  `json:"actStartDateTimeInGMTFormatted"`
	ActivityStartDateTimeLocal          int64   `json:"activityStartDateTimeLocal"`
	ActivityStartDateTimeLocalFormatted string  `json:"activityStartDateTimeLocalFormatted"`
	Value                               float64 `json:"value"`
	PrStartTimeGmt                      int64   `json:"prStartTimeGmt"`
	PrStartTimeGmtFormatted             string  `json:"prStartTimeGmtFormatted"`
	PrStartTimeLocal                    int64   `json:"prStartTimeLocal"`
	PrStartTimeLocalFormatted           string  `json:"prStartTimeLocalFormatted"`
	PrTypeLabelKey                      any     `json:"prTypeLabelKey"`
	PoolLengthUnit                      any     `json:"poolLengthUnit"`
}

func (prs *PersonalRecordService) PRs(userUUID string) (res []PersonalRecord, e error) {
	p := fmt.Sprintf("/personalrecord-service/personalrecord/prs/%s", userUUID)
	return res, prs.c.apiGet(&res, p, nil)
}

func (prs *PersonalRecordService) Candidate(userUUID string) (res []PersonalRecord, e error) {
	p := fmt.Sprintf("/personalrecord-service/personalrecordcandidate/%s", userUUID)
	return res, prs.c.apiGet(&res, p, nil)
}

type PersonalRecordType struct {
	ID       int     `json:"id"`
	Key      string  `json:"key"`
	Visible  bool    `json:"visible"`
	Sport    string  `json:"sport"`
	MinValue float64 `json:"minValue"`
	MaxValue float64 `json:"maxValue"`
}

func (prs *PersonalRecordService) PersonalRecordTypes(userUUID string) (res []PersonalRecordType, err error) {
	p := fmt.Sprintf("/personalrecord-service/personalrecordtype/prtypes/%s", userUUID)
	return res, prs.c.apiGet(&res, p, nil)
}
