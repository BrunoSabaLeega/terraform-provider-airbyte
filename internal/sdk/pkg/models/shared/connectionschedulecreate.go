// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectionScheduleCreate - schedule for when the the connection should run, per the schedule type
type ConnectionScheduleCreate struct {
	CronExpression *string              `json:"cronExpression,omitempty"`
	ScheduleType   ScheduleTypeEnumEnum `json:"scheduleType"`
}