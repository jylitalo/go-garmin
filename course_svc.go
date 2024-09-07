package garmin

import "fmt"

type CourseService service

type Course struct {
	CourseID      int    `json:"courseId"`
	UserProfileID int    `json:"userProfileId"`
	DisplayName   string `json:"displayName"`
	UserGroupID   any    `json:"userGroupId"`
	GeoRoutePk    any    `json:"geoRoutePk"`
	ActivityType  struct {
		TypeID       int    `json:"typeId"`
		TypeKey      string `json:"typeKey"`
		ParentTypeID int    `json:"parentTypeId"`
		IsHidden     bool   `json:"isHidden"`
		Restricted   bool   `json:"restricted"`
		Trimmable    bool   `json:"trimmable"`
	} `json:"activityType"`
	CourseName        string `json:"courseName"`
	CourseDescription any    `json:"courseDescription"`
	CreatedDate       int64  `json:"createdDate"`
	UpdatedDate       int64  `json:"updatedDate"`
	PrivacyRule       struct {
		TypeID  int    `json:"typeId"`
		TypeKey string `json:"typeKey"`
	} `json:"privacyRule"`
	DistanceInMeters         float64 `json:"distanceInMeters"`
	ElevationGainInMeters    float64 `json:"elevationGainInMeters"`
	ElevationLossInMeters    float64 `json:"elevationLossInMeters"`
	StartLatitude            float64 `json:"startLatitude"`
	StartLongitude           float64 `json:"startLongitude"`
	SpeedInMetersPerSecond   float64 `json:"speedInMetersPerSecond"`
	SourceTypeID             int     `json:"sourceTypeId"`
	SourcePk                 any     `json:"sourcePk"`
	ElapsedSeconds           any     `json:"elapsedSeconds"`
	CoordinateSystem         string  `json:"coordinateSystem"`
	OriginalCoordinateSystem string  `json:"originalCoordinateSystem"`
	Consumer                 any     `json:"consumer"`
	ElevationSource          int     `json:"elevationSource"`
	HasShareableEvent        bool    `json:"hasShareableEvent"`
	HasPaceBand              bool    `json:"hasPaceBand"`
	HasPowerGuide            bool    `json:"hasPowerGuide"`
	Favorite                 bool    `json:"favorite"`
	HasTurnDetectionDisabled bool    `json:"hasTurnDetectionDisabled"`
	CuratedCourseID          any     `json:"curatedCourseId"`
	Public                   bool    `json:"public"`
	ActivityTypeID           struct {
		TypeID       int    `json:"typeId"`
		TypeKey      string `json:"typeKey"`
		ParentTypeID int    `json:"parentTypeId"`
		IsHidden     bool   `json:"isHidden"`
		Restricted   bool   `json:"restricted"`
		Trimmable    bool   `json:"trimmable"`
	} `json:"activityTypeId"`
	CreatedDateFormatted string `json:"createdDateFormatted"`
	UpdatedDateFormatted string `json:"updatedDateFormatted"`
}

type UserCourses struct {
	Courses []Course `json:"coursesForUser"`
}

func (cs *CourseService) List(userDisplayName string) (*UserCourses, error) {
	// GET https://connect.garmin.com/course-service/course/owner/<user_uuid>
	var c UserCourses
	p := fmt.Sprintf("/course-service/course/owner/%s", userDisplayName)
	return &c, cs.c.apiGet(&c, p, nil)
}

func (cs *CourseService) Courses() (*UserCourses, error) {
	var c UserCourses
	return &c, cs.c.apiGet(&c, "/web-gateway/course/owner", nil)
}

type CourseMetadata struct {
	CourseID      int    `json:"courseId"`
	UserProfileID int    `json:"userProfileId"`
	Name          string `json:"name"`
	Description   any    `json:"description"`
	ActivityType  struct {
		TypeID       int    `json:"typeId"`
		TypeKey      string `json:"typeKey"`
		ParentTypeID int    `json:"parentTypeId"`
		IsHidden     bool   `json:"isHidden"`
		Restricted   bool   `json:"restricted"`
		Trimmable    bool   `json:"trimmable"`
	} `json:"activityType"`
	PrivacyRule struct {
		TypeID  int    `json:"typeId"`
		TypeKey string `json:"typeKey"`
	} `json:"privacyRule"`
	DistanceInMeters float64 `json:"distanceInMeters"`
}

func (cs *CourseService) Metadata(id int64) (*CourseMetadata, error) {
	// GET https://connect.garmin.com/course-service/course/metadata/<id>
	var cm CourseMetadata
	p := fmt.Sprintf("/course-service/course/metadata/%d", id)
	return &cm, cs.c.apiGet(&cm, p, nil)
}
