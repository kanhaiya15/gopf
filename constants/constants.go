package constants

import "time"

//
const (
	Status = "401 Unauthorized"

	AuthFailed          = "Unauthenticated"
	InternalServerError = "Oops! Looks like something just broke."
	ValidationError     = "Validation Failed"
)

//
const (
	AuthFailedEvent = "AUTH_FAILED_EVENT"
)

//
const (
	MysqlMaxIdleConnection     = 10
	MysqlMaxOpenConnection     = 150
	MysqlMaxConnectionLifetime = time.Hour
	RedisConnectionPool        = 40
	CacheAuthTTL               = 60 * 60 * 4
)
