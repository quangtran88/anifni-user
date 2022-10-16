package domain

import "time"

type KYCStatus string

const (
	Pending KYCStatus = "pending"
	Success KYCStatus = "success"
	Failed  KYCStatus = "failed"
)

type KYC struct {
	IDNumber     string
	IDFrontUrl   string
	IDBackUrl    string
	IDAddress    string
	IDIssueDate  time.Time
	IDIssuePlace string
	status       KYCStatus
}
