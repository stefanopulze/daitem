package session

import (
	"time"
)

// Session store information about user and authentication token
type Session struct {
	accessToken  string
	refreshToken string
	expireAt     time.Time
	username     string
	password     string
	masterCode   string
	userId       int
}

// New create new session with user information
func New(username string, password string, masterCode string) *Session {
	return &Session{
		username:   username,
		password:   password,
		masterCode: masterCode,
		expireAt:   time.Now(),
	}
}

func (s *Session) GetUsername() string {
	return s.username
}

func (s *Session) GetPassword() string {
	return s.password
}

func (s *Session) GetAccessToken() string {
	return s.accessToken
}

func (s *Session) GetRefreshToken() string {
	return s.refreshToken
}

func (s *Session) GetUserId() int {
	return s.userId
}

func (s *Session) GetMasterCode() string {
	return s.masterCode
}

func (s *Session) IsValid() bool {
	return time.Now().Before(s.expireAt)
}

// Update authentication session after login or refresh token call
func (s *Session) Update(accessToken string, refreshToken string, expireIn int, userId int) {
	s.accessToken = accessToken
	s.refreshToken = refreshToken
	s.expireAt = time.Now().Add(time.Duration(expireIn) * time.Second)
	s.userId = userId
}
