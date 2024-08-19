package utils

import (
	"database/sql"
	"math/rand"
	"time"
)

// ConvertSQLNullStringToString converts a sql.NullString to a string
func ConvertSQLNullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

// ConvertStringToSQLNullString converts a string to a sql.NullString
func ConvertStringToSQLNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// ConvertTotimeToSQLNullTime converts a time.Time to a sql.NullTime
func ConvertTotimeToSQLNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

// ConvertSQLNullTimeToTime converts a sql.NullTime to a time.Time
func ConvertSQLNullTimeToTime(t sql.NullTime) time.Time {
	if t.Valid {
		return t.Time
	}
	return time.Time{}
}

// ConvertBoolToSQLNullBool converts a bool to a sql.NullBool
func ConvertBoolToSQLNullBool(b bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: true,
	}
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890asdfghjklqwertyuiozxcvbnm@")

// RandomString generates random string
func RandomString(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
