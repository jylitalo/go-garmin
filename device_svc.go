package garmin

import "fmt"

type DeviceService service

type Device struct {
	AppSupport                                   bool     `json:"appSupport"`
	ApplicationKey                               string   `json:"applicationKey"`
	DeviceTypePk                                 int      `json:"deviceTypePk"`
	BestInClassVideoLink                         any      `json:"bestInClassVideoLink"`
	BluetoothClassicDevice                       bool     `json:"bluetoothClassicDevice"`
	BluetoothLowEnergyDevice                     bool     `json:"bluetoothLowEnergyDevice"`
	DeviceCategories                             []string `json:"deviceCategories"`
	DeviceEmbedVideoLink                         any      `json:"deviceEmbedVideoLink"`
	DeviceSettingsFile                           string   `json:"deviceSettingsFile"`
	GcmSettingsFile                              any      `json:"gcmSettingsFile"`
	DeviceVideoPageLink                          any      `json:"deviceVideoPageLink"`
	DisplayOrder                                 int      `json:"displayOrder"`
	GolfDisplayOrder                             int      `json:"golfDisplayOrder"`
	HasOpticalHeartRate                          bool     `json:"hasOpticalHeartRate"`
	Highlighted                                  bool     `json:"highlighted"`
	Hybrid                                       bool     `json:"hybrid"`
	ImageURL                                     string   `json:"imageUrl"`
	MinGCMAndroidVersion                         int      `json:"minGCMAndroidVersion"`
	MinGCMWindowsVersion                         int      `json:"minGCMWindowsVersion"`
	MinGCMiOSVersion                             int      `json:"minGCMiOSVersion"`
	MinGolfAppiOSVersion                         int      `json:"minGolfAppiOSVersion"`
	MinGolfAppAndroidVersion                     int      `json:"minGolfAppAndroidVersion"`
	PartNumber                                   string   `json:"partNumber"`
	Primary                                      bool     `json:"primary"`
	ProductDisplayName                           string   `json:"productDisplayName"`
	DeviceTags                                   any      `json:"deviceTags"`
	ProductSku                                   string   `json:"productSku"`
	Wasp                                         bool     `json:"wasp"`
	WeightScale                                  bool     `json:"weightScale"`
	Wellness                                     bool     `json:"wellness"`
	Wifi                                         bool     `json:"wifi"`
	HasPowerButton                               bool     `json:"hasPowerButton"`
	SupportsSecondaryUsers                       bool     `json:"supportsSecondaryUsers"`
	PrimaryApplication                           string   `json:"primaryApplication"`
	IncompatibleApplications                     []any    `json:"incompatibleApplications"`
	AbnormalHeartRateAlertCapable                bool     `json:"abnormalHeartRateAlertCapable"`
	ActivitySummFitFileCapable                   bool     `json:"activitySummFitFileCapable"`
	AerobicTrainingEffectCapable                 bool     `json:"aerobicTrainingEffectCapable"`
	AlarmDaysCapable                             bool     `json:"alarmDaysCapable"`
	AllDayStressCapable                          bool     `json:"allDayStressCapable"`
	AnaerobicTrainingEffectCapable               bool     `json:"anaerobicTrainingEffectCapable"`
	AtpWorkoutCapable                            bool     `json:"atpWorkoutCapable"`
	BodyBatteryCapable                           bool     `json:"bodyBatteryCapable"`
	BrickWorkoutCapable                          bool     `json:"brickWorkoutCapable"`
	CardioCapable                                bool     `json:"cardioCapable"`
	CardioOptionCapable                          bool     `json:"cardioOptionCapable"`
	CardioSportsCapable                          bool     `json:"cardioSportsCapable"`
	CardioWorkoutCapable                         bool     `json:"cardioWorkoutCapable"`
	CellularCapable                              bool     `json:"cellularCapable"`
	ChangeLogCapable                             bool     `json:"changeLogCapable"`
	ContactManagementCapable                     bool     `json:"contactManagementCapable"`
	CourseCapable                                bool     `json:"courseCapable"`
	CourseFileType                               string   `json:"courseFileType"`
	CoursePromptCapable                          bool     `json:"coursePromptCapable"`
	CustomIntensityMinutesCapable                bool     `json:"customIntensityMinutesCapable"`
	CustomWorkoutCapable                         bool     `json:"customWorkoutCapable"`
	CyclingSegmentCapable                        bool     `json:"cyclingSegmentCapable"`
	CyclingSportsCapable                         bool     `json:"cyclingSportsCapable"`
	CyclingWorkoutCapable                        bool     `json:"cyclingWorkoutCapable"`
	DefaultSettingCapable                        bool     `json:"defaultSettingCapable"`
	DeviceSettingCapable                         bool     `json:"deviceSettingCapable"`
	DeviceSettingFileType                        any      `json:"deviceSettingFileType"`
	DisplayFieldsExtCapable                      bool     `json:"displayFieldsExtCapable"`
	DivingCapable                                bool     `json:"divingCapable"`
	EllipticalOptionCapable                      bool     `json:"ellipticalOptionCapable"`
	FloorsClimbedGoalCapable                     bool     `json:"floorsClimbedGoalCapable"`
	FtpCapable                                   bool     `json:"ftpCapable"`
	Gcj02CourseCapable                           bool     `json:"gcj02CourseCapable"`
	GlonassCapable                               bool     `json:"glonassCapable"`
	GoalCapable                                  bool     `json:"goalCapable"`
	GoalFileType                                 string   `json:"goalFileType"`
	GolfAppSyncCapable                           bool     `json:"golfAppSyncCapable"`
	GpsRouteCapable                              bool     `json:"gpsRouteCapable"`
	HandednessCapable                            bool     `json:"handednessCapable"`
	HrZoneCapable                                bool     `json:"hrZoneCapable"`
	HrvStressCapable                             bool     `json:"hrvStressCapable"`
	IntensityMinutesGoalCapable                  bool     `json:"intensityMinutesGoalCapable"`
	LactateThresholdCapable                      bool     `json:"lactateThresholdCapable"`
	LanguageSettingCapable                       bool     `json:"languageSettingCapable"`
	LanguageSettingFileType                      any      `json:"languageSettingFileType"`
	LowHrAlertCapable                            bool     `json:"lowHrAlertCapable"`
	MaxHRCapable                                 bool     `json:"maxHRCapable"`
	MaxWorkoutCount                              int      `json:"maxWorkoutCount"`
	MetricsFitFileReceiveCapable                 bool     `json:"metricsFitFileReceiveCapable"`
	MetricsUploadCapable                         bool     `json:"metricsUploadCapable"`
	MilitaryTimeCapable                          bool     `json:"militaryTimeCapable"`
	ModerateIntensityMinutesGoalCapable          bool     `json:"moderateIntensityMinutesGoalCapable"`
	NfcCapable                                   bool     `json:"nfcCapable"`
	OtherOptionCapable                           bool     `json:"otherOptionCapable"`
	OtherSportsCapable                           bool     `json:"otherSportsCapable"`
	PersonalRecordCapable                        bool     `json:"personalRecordCapable"`
	PersonalRecordFileType                       string   `json:"personalRecordFileType"`
	PoolSwimOptionCapable                        bool     `json:"poolSwimOptionCapable"`
	PowerCurveCapable                            bool     `json:"powerCurveCapable"`
	PowerZonesCapable                            bool     `json:"powerZonesCapable"`
	PulseOxAllDayCapable                         bool     `json:"pulseOxAllDayCapable"`
	PulseOxOnDemandCapable                       bool     `json:"pulseOxOnDemandCapable"`
	PulseOxSleepCapable                          bool     `json:"pulseOxSleepCapable"`
	RemCapable                                   bool     `json:"remCapable"`
	ReminderAlarmCapable                         bool     `json:"reminderAlarmCapable"`
	ReorderablePagesCapable                      bool     `json:"reorderablePagesCapable"`
	RestingHRCapable                             bool     `json:"restingHRCapable"`
	RideOptionsCapable                           bool     `json:"rideOptionsCapable"`
	RunOptionIndoorCapable                       bool     `json:"runOptionIndoorCapable"`
	RunOptionsCapable                            bool     `json:"runOptionsCapable"`
	RunningSegmentCapable                        bool     `json:"runningSegmentCapable"`
	RunningSportsCapable                         bool     `json:"runningSportsCapable"`
	RunningWorkoutCapable                        bool     `json:"runningWorkoutCapable"`
	ScheduleCapable                              bool     `json:"scheduleCapable"`
	ScheduleFileType                             string   `json:"scheduleFileType"`
	SegmentCapable                               bool     `json:"segmentCapable"`
	SegmentPointCapable                          bool     `json:"segmentPointCapable"`
	SettingCapable                               bool     `json:"settingCapable"`
	SettingFileType                              string   `json:"settingFileType"`
	SleepTimeCapable                             bool     `json:"sleepTimeCapable"`
	SmallFitFileOnlyCapable                      bool     `json:"smallFitFileOnlyCapable"`
	SportCapable                                 bool     `json:"sportCapable"`
	SportFileType                                string   `json:"sportFileType"`
	StairStepperOptionCapable                    bool     `json:"stairStepperOptionCapable"`
	StrengthOptionsCapable                       bool     `json:"strengthOptionsCapable"`
	StrengthWorkoutCapable                       bool     `json:"strengthWorkoutCapable"`
	SupportedHrZones                             []string `json:"supportedHrZones"`
	SwimWorkoutCapable                           bool     `json:"swimWorkoutCapable"`
	TrainingPlanCapable                          bool     `json:"trainingPlanCapable"`
	TrainingStatusCapable                        bool     `json:"trainingStatusCapable"`
	TrainingStatusPauseCapable                   bool     `json:"trainingStatusPauseCapable"`
	UserProfileCapable                           bool     `json:"userProfileCapable"`
	UserProfileFileType                          any      `json:"userProfileFileType"`
	UserTcxExportCapable                         bool     `json:"userTcxExportCapable"`
	Vo2MaxBikeCapable                            bool     `json:"vo2MaxBikeCapable"`
	Vo2MaxRunCapable                             bool     `json:"vo2MaxRunCapable"`
	WalkOptionCapable                            bool     `json:"walkOptionCapable"`
	WalkingSportsCapable                         bool     `json:"walkingSportsCapable"`
	WeatherAlertsCapable                         bool     `json:"weatherAlertsCapable"`
	WeatherSettingsCapable                       bool     `json:"weatherSettingsCapable"`
	WorkoutCapable                               bool     `json:"workoutCapable"`
	WorkoutFileType                              string   `json:"workoutFileType"`
	YogaCapable                                  bool     `json:"yogaCapable"`
	YogaOptionCapable                            bool     `json:"yogaOptionCapable"`
	HeatAndAltitudeAcclimationCapable            bool     `json:"heatAndAltitudeAcclimationCapable"`
	TrainingLoadBalanceCapable                   bool     `json:"trainingLoadBalanceCapable"`
	IndoorTrackOptionsCapable                    bool     `json:"indoorTrackOptionsCapable"`
	IndoorBikeOptionsCapable                     bool     `json:"indoorBikeOptionsCapable"`
	IndoorWalkOptionsCapable                     bool     `json:"indoorWalkOptionsCapable"`
	TrainingEffectLabelCapable                   bool     `json:"trainingEffectLabelCapable"`
	PacebandCapable                              bool     `json:"pacebandCapable"`
	RespirationCapable                           bool     `json:"respirationCapable"`
	OpenWaterSwimOptionCapable                   bool     `json:"openWaterSwimOptionCapable"`
	PhoneVerificationCheckRequired               bool     `json:"phoneVerificationCheckRequired"`
	WeightGoalCapable                            bool     `json:"weightGoalCapable"`
	YogaWorkoutCapable                           bool     `json:"yogaWorkoutCapable"`
	PilatesWorkoutCapable                        bool     `json:"pilatesWorkoutCapable"`
	ConnectedGPSCapable                          bool     `json:"connectedGPSCapable"`
	DiveAppSyncCapable                           bool     `json:"diveAppSyncCapable"`
	GolfLiveScoringCapable                       bool     `json:"golfLiveScoringCapable"`
	SolarPanelUtilizationCapable                 bool     `json:"solarPanelUtilizationCapable"`
	SweatLossCapable                             bool     `json:"sweatLossCapable"`
	DiveAlertCapable                             bool     `json:"diveAlertCapable"`
	RequiresInitialDeviceNickname                bool     `json:"requiresInitialDeviceNickname"`
	DefaultSettingsHbaseMigrated                 bool     `json:"defaultSettingsHbaseMigrated"`
	SleepScoreCapable                            bool     `json:"sleepScoreCapable"`
	FitnessAgeV2Capable                          bool     `json:"fitnessAgeV2Capable"`
	IntensityMinutesV2Capable                    bool     `json:"intensityMinutesV2Capable"`
	CollapsibleControlMenuCapable                bool     `json:"collapsibleControlMenuCapable"`
	MeasurementUnitSettingCapable                bool     `json:"measurementUnitSettingCapable"`
	OnDeviceSleepCalculationCapable              bool     `json:"onDeviceSleepCalculationCapable"`
	HiitWorkoutCapable                           bool     `json:"hiitWorkoutCapable"`
	RunningHeartRateZoneCapable                  bool     `json:"runningHeartRateZoneCapable"`
	CyclingHeartRateZoneCapable                  bool     `json:"cyclingHeartRateZoneCapable"`
	SwimmingHeartRateZoneCapable                 bool     `json:"swimmingHeartRateZoneCapable"`
	DefaultHeartRateZoneCapable                  bool     `json:"defaultHeartRateZoneCapable"`
	CyclingPowerZonesCapable                     bool     `json:"cyclingPowerZonesCapable"`
	XcSkiPowerZonesCapable                       bool     `json:"xcSkiPowerZonesCapable"`
	SwimAlgorithmCapable                         bool     `json:"swimAlgorithmCapable"`
	BenchmarkExerciseCapable                     bool     `json:"benchmarkExerciseCapable"`
	SpectatorMessagingCapable                    bool     `json:"spectatorMessagingCapable"`
	EcgCapable                                   bool     `json:"ecgCapable"`
	LteLiveEventSharingCapable                   bool     `json:"lteLiveEventSharingCapable"`
	SleepFitFileReceiveCapable                   bool     `json:"sleepFitFileReceiveCapable"`
	SecondaryWorkoutStepTargetCapable            bool     `json:"secondaryWorkoutStepTargetCapable"`
	AssistancePlusCapable                        bool     `json:"assistancePlusCapable"`
	PowerGuidanceCapable                         bool     `json:"powerGuidanceCapable"`
	AirIntegrationCapable                        bool     `json:"airIntegrationCapable"`
	HealthSnapshotCapable                        bool     `json:"healthSnapshotCapable"`
	RacePredictionsRunCapable                    bool     `json:"racePredictionsRunCapable"`
	VivohubCompatible                            bool     `json:"vivohubCompatible"`
	StepsTrueUpChartCapable                      bool     `json:"stepsTrueUpChartCapable"`
	SportingEventCapable                         bool     `json:"sportingEventCapable"`
	SolarChargeCapable                           bool     `json:"solarChargeCapable"`
	RealTimeSettingsCapable                      bool     `json:"realTimeSettingsCapable"`
	EmergencyCallingCapable                      bool     `json:"emergencyCallingCapable"`
	PersonalRepRecordCapable                     bool     `json:"personalRepRecordCapable"`
	HrvStatusCapable                             bool     `json:"hrvStatusCapable"`
	TrainingReadinessCapable                     bool     `json:"trainingReadinessCapable"`
	PublicBetaSoftwareCapable                    bool     `json:"publicBetaSoftwareCapable"`
	WorkoutAudioPromptsCapable                   bool     `json:"workoutAudioPromptsCapable"`
	ActualStepRecordingCapable                   bool     `json:"actualStepRecordingCapable"`
	GroupTrack2Capable                           bool     `json:"groupTrack2Capable"`
	GolfAppPairingCapable                        bool     `json:"golfAppPairingCapable"`
	LocalWindConditionsCapable                   bool     `json:"localWindConditionsCapable"`
	MultipleGolfCourseCapable                    bool     `json:"multipleGolfCourseCapable"`
	BeaconTrackingCapable                        bool     `json:"beaconTrackingCapable"`
	BatteryStatusCapable                         bool     `json:"batteryStatusCapable"`
	RunningPowerZonesCapable                     bool     `json:"runningPowerZonesCapable"`
	AcuteTrainingLoadCapable                     bool     `json:"acuteTrainingLoadCapable"`
	CriticalSwimSpeedCapable                     bool     `json:"criticalSwimSpeedCapable"`
	PrimaryTrainingCapable                       bool     `json:"primaryTrainingCapable"`
	DayOfWeekSleepWindowCapable                  bool     `json:"dayOfWeekSleepWindowCapable"`
	GolfCourseDownloadCapable                    bool     `json:"golfCourseDownloadCapable"`
	LaunchMonitorEventSharingCapable             bool     `json:"launchMonitorEventSharingCapable"`
	LhaBackupCapable                             bool     `json:"lhaBackupCapable"`
	JetlagCapable                                bool     `json:"jetlagCapable"`
	BloodPressureCapable                         bool     `json:"bloodPressureCapable"`
	BbiRecordingCapable                          bool     `json:"bbiRecordingCapable"`
	WheelchairCapable                            bool     `json:"wheelchairCapable"`
	PrimaryActivityTrackerSettingCapable         bool     `json:"primaryActivityTrackerSettingCapable"`
	SetBodyCompositionCapable                    bool     `json:"setBodyCompositionCapable"`
	AcuteChronicWorkloadRatioCapable             bool     `json:"acuteChronicWorkloadRatioCapable"`
	SleepNeedCapable                             bool     `json:"sleepNeedCapable"`
	WearableBackupRestoreCapable                 bool     `json:"wearableBackupRestoreCapable"`
	CyclingComputerBackupRestoreCapable          bool     `json:"cyclingComputerBackupRestoreCapable"`
	DescriptiveTrainingEffectCapable             bool     `json:"descriptiveTrainingEffectCapable"`
	SleepSkinTemperatureCapable                  bool     `json:"sleepSkinTemperatureCapable"`
	RunningLactateThresholdCapable               bool     `json:"runningLactateThresholdCapable"`
	AltitudeAcclimationPercentageCapable         bool     `json:"altitudeAcclimationPercentageCapable"`
	HillScoreAndEnduranceScoreCapable            bool     `json:"hillScoreAndEnduranceScoreCapable"`
	SwimWorkout2Capable                          bool     `json:"swimWorkout2Capable"`
	EnhancedWorkoutStepCapable                   bool     `json:"enhancedWorkoutStepCapable"`
	PrimaryTrainingBackupCapable                 bool     `json:"primaryTrainingBackupCapable"`
	HideSoftwareUpdateVersionCapable             bool     `json:"hideSoftwareUpdateVersionCapable"`
	AdaptiveCoachingScheduleCapable              bool     `json:"adaptiveCoachingScheduleCapable"`
	OpenWaterSwimWorkoutCapable                  bool     `json:"openWaterSwimWorkoutCapable"`
	MultisportWorkoutCapable                     bool     `json:"multisportWorkoutCapable"`
	FullNameSettingsCapable                      bool     `json:"fullNameSettingsCapable"`
	BodyBatteryTrueUpCapable                     bool     `json:"bodyBatteryTrueUpCapable"`
	SkipLastWorkoutStepInRepeatBlockCapable      bool     `json:"skipLastWorkoutStepInRepeatBlockCapable"`
	AdaptiveCoachingStrengthCapable              bool     `json:"adaptiveCoachingStrengthCapable"`
	MobilityWorkoutCapable                       bool     `json:"mobilityWorkoutCapable"`
	CoursePlannerCapable                         bool     `json:"coursePlannerCapable"`
	TrainableTriathlonEventCapable               bool     `json:"trainableTriathlonEventCapable"`
	PasscodeLockCapable                          bool     `json:"passcodeLockCapable"`
	StrengthTrainingPlanCapable                  bool     `json:"strengthTrainingPlanCapable"`
	BreathingDisruptionsCapable                  bool     `json:"breathingDisruptionsCapable"`
	StandaloneFirstbeatActivityProcessingCapable bool     `json:"standaloneFirstbeatActivityProcessingCapable"`
	Rts20Capable                                 bool     `json:"rts20Capable"`
	Datasource                                   string   `json:"datasource"`
	DeviceStatus                                 string   `json:"deviceStatus"`
	RegisteredDate                               int64    `json:"registeredDate"`
	ActualProductSku                             string   `json:"actualProductSku"`
	VivohubConfigurable                          any      `json:"vivohubConfigurable"`
	CorporateDevice                              bool     `json:"corporateDevice"`
	PrePairedWithHRM                             bool     `json:"prePairedWithHRM"`
	UnRetirable                                  bool     `json:"unRetirable"`
	SerialNumber                                 string   `json:"serialNumber"`
	ShortName                                    any      `json:"shortName"`
	DisplayName                                  string   `json:"displayName"`
	DeviceID                                     int64    `json:"deviceId"`
	UnitID                                       int64    `json:"unitId"`
	WifiSetup                                    bool     `json:"wifiSetup"`
	CurrentFirmwareVersion                       string   `json:"currentFirmwareVersion"`
	CurrentFirmwareVersionMajor                  int      `json:"currentFirmwareVersionMajor"`
	CurrentFirmwareVersionMinor                  int      `json:"currentFirmwareVersionMinor"`
	ActiveInd                                    int      `json:"activeInd"`
	PrimaryActivityTrackerIndicator              bool     `json:"primaryActivityTrackerIndicator"`
	IsPrimaryUser                                bool     `json:"isPrimaryUser"`
	OtherAssociation                             bool     `json:"otherAssociation"`
}

func (d *DeviceService) Devices() (res []Device, e error) {
	return res, d.c.apiGet(&res, "/device-service/deviceregistration/devices", nil)
}

type DeviceLastUsed struct {
	UserDeviceID                 int64  `json:"userDeviceId"`
	UserProfileNumber            int    `json:"userProfileNumber"`
	ApplicationNumber            int    `json:"applicationNumber"`
	LastUsedDeviceApplicationKey string `json:"lastUsedDeviceApplicationKey"`
	LastUsedDeviceName           string `json:"lastUsedDeviceName"`
	LastUsedDeviceUploadTime     int64  `json:"lastUsedDeviceUploadTime"`
	ImageURL                     string `json:"imageUrl"`
	Released                     bool   `json:"released"`
}

func (d *DeviceService) LastUsed() (*DeviceLastUsed, error) {
	var lu DeviceLastUsed
	return &lu, d.c.apiGet(&lu, "/device-service/deviceservice/mylastused", nil)
}

type DeviceMessages struct {
	ServiceHost   string `json:"serviceHost"`
	NumOfMessages int    `json:"numOfMessages"`
	Messages      []any  `json:"messages"`
}

func (d *DeviceService) DeviceMessages() (*DeviceMessages, error) {
	var dm DeviceMessages
	return &dm, d.c.apiGet(&dm, "/device-service/devicemessage/messages", nil)
}

func (d *DeviceService) DeviceMessageCount() (c int, e error) {
	err := d.c.apiGet(&c, "/device-service/devicemessage/message/count", nil)
	if err != nil {
		return 0, err
	}
	return c, nil
}

// TODO GET https://connect.garmin.com/device-service/deviceregistration/devices/historical

type UserDevice struct {
	DeviceID              int64 `json:"deviceId"`
	UserID                int   `json:"userId"`
	ApplicationVersionID  int   `json:"applicationVersionId"`
	ApplicationID         int   `json:"applicationId"`
	LastUploadTimestamp   int64 `json:"lastUploadTimestamp"`
	LastDownloadTimestamp int64 `json:"lastDownloadTimestamp"`
	DeviceStatus          any   `json:"deviceStatus"`
}

func (d *DeviceService) UserDevice(deviceID int64) (*UserDevice, error) {
	var ud UserDevice
	p := fmt.Sprintf("/device-service/deviceservice/user-device/%d", deviceID)
	return &ud, d.c.apiGet(&ud, p, nil)
}

func (d *DeviceService) DevicesByUser(userUUID string) (res []Device, e error) {
	p := fmt.Sprintf("/device-service/deviceregistration/devices/all/%s", userUUID)
	return res, d.c.apiGet(&res, p, nil)
}

type PrimaryTrainingDevice struct {
	PrimaryTrainingDevice struct {
		DeviceID int64 `json:"deviceId"`
	} `json:"PrimaryTrainingDevice"`
	WearableDevices struct {
		DeviceWeights []struct {
			DisplayName            string `json:"displayName"`
			DeviceID               int64  `json:"deviceId"`
			ImageURL               string `json:"imageUrl"`
			Weight                 int    `json:"weight"`
			PrimaryTrainingCapable bool   `json:"primaryTrainingCapable"`
			LhaBackupCapable       bool   `json:"lhaBackupCapable"`
			PrimaryWearableDevice  bool   `json:"primaryWearableDevice"`
		} `json:"deviceWeights"`
		WearableDeviceCount int `json:"wearableDeviceCount"`
	} `json:"WearableDevices"`
	TrainingStatusOnlyDevices struct {
		DeviceWeights []any `json:"deviceWeights"`
	} `json:"TrainingStatusOnlyDevices"`
	PrimaryTrainingDevices struct {
		DeviceWeights []struct {
			DisplayName            string `json:"displayName"`
			DeviceID               int64  `json:"deviceId"`
			ImageURL               string `json:"imageUrl"`
			Weight                 int    `json:"weight"`
			PrimaryTrainingCapable bool   `json:"primaryTrainingCapable"`
			LhaBackupCapable       bool   `json:"lhaBackupCapable"`
			PrimaryWearableDevice  bool   `json:"primaryWearableDevice"`
		} `json:"deviceWeights"`
		PrimaryTrainingDeviceCount int `json:"primaryTrainingDeviceCount"`
	} `json:"PrimaryTrainingDevices"`
	RegisteredDevices []Device `json:"RegisteredDevices"`
}

func (d *DeviceService) PrimaryTrainingDevice() (*PrimaryTrainingDevice, error) {
	var pd PrimaryTrainingDevice
	return &pd, d.c.apiGet(&pd, "/web-gateway/device-info/primary-training-device", nil)
}

type DeviceMessage struct {
	DeviceID    int64  `json:"deviceId"`
	MessageURL  string `json:"messageUrl"`
	MessageType string `json:"messageType"`
	MessageName string `json:"messageName"`
	GroupName   any    `json:"groupName"`
	Priority    int    `json:"priority"`
	FileType    string `json:"fileType"`
	MetaDataID  int64  `json:"metaDataId"`
}

type UploadedDeviceMessage struct {
	DeviceID          int64  `json:"deviceId"`
	DeviceName        string `json:"deviceName,omitempty"`
	MessageID         int64  `json:"messageId"`
	MessageType       string `json:"messageType"`
	MessageStatus     string `json:"messageStatus,omitempty"`
	ApplicationKey    any    `json:"applicationKey"`
	FirmwareVersion   any    `json:"firmwareVersion"`
	WifiSetup         bool   `json:"wifiSetup"`
	DeviceXMLDataType any    `json:"deviceXmlDataType,omitempty"`
	Hidden            bool   `json:"hidden,omitempty"`
	CreatedTimeStamp  any    `json:"createdTimeStamp,omitempty"`
	UpdatedTimeStamp  any    `json:"updatedTimeStamp,omitempty"`
	FileType          string `json:"fileType"`
	MessageURL        string `json:"messageUrl"`
	UniqueIdentifier  any    `json:"uniqueIdentifier,omitempty"`
	MessageName       string `json:"messageName"`
	GroupName         any    `json:"groupName"`
	Priority          int    `json:"priority"`
	MetaDataID        int    `json:"metaDataId"`
	AppDetails        any    `json:"appDetails,omitempty"`
}

func (d *DeviceService) SendDeviceMessages(msgs []DeviceMessage) (res []UploadedDeviceMessage, err error) {
	// POST https://connect.garmin.com/device-service/devicemessage/messages
	// Content-Type: application/json
	//
	// [{ ... }]
	_, err = d.c.api(&res, "POST", "/device-service/devicemessage/messages", nil, msgs)
	return res, err
}

func (d *DeviceService) SendCourceToDevice(deviceID, courseID int64, courseName string) error {
	_, err := d.SendDeviceMessages([]DeviceMessage{{
		DeviceID:    deviceID,
		MessageURL:  fmt.Sprintf("course-service/course/fit/%d/%d?elevation=true", courseID, deviceID),
		FileType:    "FIT",
		MessageType: "courses",
		MessageName: courseName,
		GroupName:   nil,
		MetaDataID:  courseID,
	}})
	return err
}
