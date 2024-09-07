package garmin

import (
	"fmt"
	"net/url"
	"strconv"
)

type BadgeService service

type Badge struct {
	ID                    int           `json:"badgeId"`
	Key                   string        `json:"badgeKey"`
	Name                  string        `json:"badgeName"`
	UUID                  any           `json:"badgeUuid"`
	CategoryID            int           `json:"badgeCategoryId"`
	DifficultyID          int           `json:"badgeDifficultyId"`
	Points                int           `json:"badgePoints"`
	TypeIds               []int         `json:"badgeTypeIds"`
	SeriesID              int           `json:"badgeSeriesId"`
	StartDate             string        `json:"badgeStartDate"`
	EndDate               any           `json:"badgeEndDate"`
	UserProfileID         int           `json:"userProfileId"`
	FullName              string        `json:"fullName"`
	DisplayName           string        `json:"displayName"`
	EarnedDate            string        `json:"badgeEarnedDate"`
	EarnedNumber          int           `json:"badgeEarnedNumber"`
	LimitCount            any           `json:"badgeLimitCount"`
	IsViewed              bool          `json:"badgeIsViewed"`
	ProgressValue         float64       `json:"badgeProgressValue"`
	TargetValue           any           `json:"badgeTargetValue"`
	UnitID                any           `json:"badgeUnitId"`
	AssocTypeID           int           `json:"badgeAssocTypeId"`
	AssocDataID           string        `json:"badgeAssocDataId"`
	AssocDataName         any           `json:"badgeAssocDataName"`
	EarnedByMe            bool          `json:"earnedByMe"`
	CurrentPlayerType     any           `json:"currentPlayerType"`
	UserJoined            any           `json:"userJoined"`
	ChallengeStatusID     any           `json:"badgeChallengeStatusId"`
	PromotionCodeTypeList []any         `json:"badgePromotionCodeTypeList"`
	PromotionCodeStatus   any           `json:"promotionCodeStatus"`
	CreateDate            string        `json:"createDate"`
	RelatedBadges         []BadgeSparse `json:"relatedBadges"`
	ConnectionNumber      any           `json:"connectionNumber"`
	Connections           any           `json:"connections"`
}

type BadgeSparse struct {
	ID           int    `json:"badgeId"`
	Key          string `json:"badgeKey"`
	UUID         any    `json:"badgeUuid"`
	Name         string `json:"badgeName"`
	DifficultyID int    `json:"badgeDifficultyId"`
	Points       int    `json:"badgePoints"`
	TypeIds      []int  `json:"badgeTypeIds"`
	EarnedByMe   bool   `json:"earnedByMe"`
	CategoryID   int    `json:"badgeCategoryId"`
}

func (b *BadgeService) Earned() (res []Badge, e error) {
	return res, b.c.apiGet(&res, "/badge-service/badge/earned", nil)
}

func (b *BadgeService) Badge(id int64) (*Badge, error) {
	var res Badge
	p := fmt.Sprintf("/badge-service/badge/detail/v2/%d", id)
	return &res, b.c.apiGet(&res, p, nil)
}

func (b *BadgeService) Available() (res []Badge, e error) {
	return res, b.c.apiGet(&res, "/badge-service/badge/available", nil)
}

func (b *BadgeService) ActivityBadges(userUUID string, activityID int64) (res []BadgeSparse, e error) {
	p := fmt.Sprintf("/badge-service/badge/%s/earned/activity/%d", userUUID, activityID)
	return res, b.c.apiGet(&res, p, nil)
}

type BadgeLeaderboard struct {
	Connections            []BadgeLeaderboardConnection `json:"connections"`
	PublicConnectionCount  int                          `json:"publicConnectionCount"`
	PrivateConnectionCount int                          `json:"privateConnectionCount"`
}

type BadgeLeaderboardConnection struct {
	UserProfileID         int    `json:"userProfileId"`
	FullName              string `json:"fullName"`
	DisplayName           string `json:"displayName"`
	UserPro               bool   `json:"userPro"`
	ProfileImageURLLarge  any    `json:"profileImageUrlLarge"`
	ProfileImageURLMedium string `json:"profileImageUrlMedium"`
	ProfileImageURLSmall  string `json:"profileImageUrlSmall"`
	UserLevel             int    `json:"userLevel"`
	UserPoint             int    `json:"userPoint"`
	LevelPointThreshold   any    `json:"levelPointThreshold"`
	LevelUpdateDate       any    `json:"levelUpdateDate"`
	LevelIsViewed         any    `json:"levelIsViewed"`
	HasPrivate            bool   `json:"hasPrivate"`
	Badges                any    `json:"badges"`
}

func (b *BadgeService) Leaderboard(limit int) (*BadgeLeaderboard, error) {
	var bl BadgeLeaderboard
	return &bl, b.c.apiGet(
		&bl,
		"/badge-service/badge/leaderboard",
		url.Values{"limit": []string{strconv.FormatInt(int64(limit), 10)}},
	)
}

type BadgeAttributes struct {
	BadgeTypes []struct {
		BadgeTypeID  int    `json:"badgeTypeId"`
		BadgeTypeKey string `json:"badgeTypeKey"`
	} `json:"badgeTypes"`
	BadgeCategories []struct {
		BadgeCategoryID  int    `json:"badgeCategoryId"`
		BadgeCategoryKey string `json:"badgeCategoryKey"`
	} `json:"badgeCategories"`
	BadgeDifficulties []struct {
		BadgeDifficultyID  int    `json:"badgeDifficultyId"`
		BadgeDifficultyKey string `json:"badgeDifficultyKey"`
		BadgePoints        int    `json:"badgePoints"`
	} `json:"badgeDifficulties"`
	BadgeUnits []struct {
		BadgeUnitID  int    `json:"badgeUnitId"`
		BadgeUnitKey string `json:"badgeUnitKey"`
	} `json:"badgeUnits"`
	BadgeAssocTypes []struct {
		BadgeAssocTypeID  int    `json:"badgeAssocTypeId"`
		BadgeAssocTypeKey string `json:"badgeAssocTypeKey"`
	} `json:"badgeAssocTypes"`
}

func (b *BadgeService) Attributes() (*BadgeAttributes, error) {
	var ba BadgeAttributes
	return &ba, b.c.apiGet(&ba, "/badge-service/badge/attributes", nil)
}
