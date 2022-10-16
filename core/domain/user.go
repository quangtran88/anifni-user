package domain

const UserTokenDomain = "U"
const UserTokenLength = 5

type User struct {
	Id            ID
	PId           PID
	Name          string
	LastName      string
	FirstName     string
	Email         string
	Password      string
	Blocked       bool
	BlockedReason string
	Kyc           KYC
	Investment    UserInvestment
	Staff         Staff
}

type RegisterUserInput struct {
	OTPCode   string
	Email     string
	Password  string
	LastName  string
	FirstName string
}
