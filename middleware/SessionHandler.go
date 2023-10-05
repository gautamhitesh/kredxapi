package middleware

import "time"

type Session struct {
	username string
	expiry   time.Time
}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

var Sessions = map[string]Session{}
