package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/Vedjw/lensvault/rand"
)

const MinBytesPerToken = 32

type Session struct {
	ID        uint
	UserID    uint
	Token     string // Token is only set when creating a new session. When looking up a session
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken uint
}

func (ss *SessionService) Create(userID uint) (*Session, error) {
	session := &Session{
		UserID: userID,
	}
	err := ss.setTokenHash(ss.setToken(session))
	if err != nil {
		return nil, fmt.Errorf("setting session token: %w", err)
	}
	err = ss.updateSession(session)
	if err == sql.ErrNoRows {
		err = ss.createSession(session)
	}
	if err != nil {
		return nil, fmt.Errorf("models: create session: %w", err)
	}
	return session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)
	var userId uint
	err := ss.getUserIdFromSession(tokenHash, &userId)
	if err != nil {
		return nil, fmt.Errorf("user not in sessions: %w", err)
	}
	user := &User{
		PK: userId,
	}
	err = ss.getUser(user)
	if err != nil {
		return nil, fmt.Errorf("user session: %w", err)
	}
	return user, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}

func (ss *SessionService) setToken(s *Session) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create token: %w", err)
	}
	s.Token = token
	return s, nil
}

func (ss *SessionService) setTokenHash(s *Session, err error) error {
	s.TokenHash = ss.hash(s.Token)
	return err
}

func (ss *SessionService) updateSession(s *Session) error {
	row := ss.DB.QueryRow(`
		UPDATE sessions
		SET token_hash=$2
		WHERE user_id=$1
		RETURNING id;`,
		s.UserID, s.TokenHash)
	err := row.Scan(&s.ID)
	return err
}

func (ss *SessionService) createSession(s *Session) error {
	row := ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2) RETURNING id;`,
		s.UserID, s.TokenHash)
	err := row.Scan(&s.ID)
	return err
}

func (ss *SessionService) getUserIdFromSession(th string, uid *uint) error {
	row := ss.DB.QueryRow(`
		SELECT user_id
		FROM sessions WHERE token_hash=$1`, th)
	err := row.Scan(uid)
	return err
}

func (ss *SessionService) getUser(u *User) error {
	row := ss.DB.QueryRow(`
		SELECT id, first_name, last_name, age, password_hash
		FROM users WHERE pk=$1`, u.PK)
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Age, &u.PasswordHash)
	return err
}
