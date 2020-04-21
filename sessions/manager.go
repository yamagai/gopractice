package sessions

import (
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
    "net/http"
)

type Manager struct {
    database map[string]interface{}
}

var kvs Manager

func init() {
    kvs.database = map[string]interface{}{}
}

func NewManager() *Manager {
    return &kvs
}

func (s *Manager) NewSessionID() string {
  b := make([]byte, 64)
  if _, err := io.ReadFull(rand.Reader, b); err != nil {
      return ""
  }
  return base64.URLEncoding.EncodeToString(b)
}

func (s *Manager) Exists(sessionID string) bool {
    _, r := s.database[sessionID]
    return r
}

func (s *Manager) Flush() {
    s.database = map[string]interface{}{}
}

func (s *Manager) Get(r *http.Request, cookieName string) (*Session, error) {
    cookie, err := r.Cookie(cookieName)
    if err != nil {
        // No cookies in the request.
        return nil, err
    }

    sessionID := cookie.Value
    // restore session
    buffer, exists := s.database[sessionID]
    if !exists {
        return nil, errors.New("Invalid sessionID")
    }

    session := buffer.(*Session)
    session.request = r
    return session, nil
}

func (s *Manager) New(r *http.Request, cookieName string) (*Session, error) {
    cookie, err := r.Cookie(cookieName)
    if err == nil && s.Exists(cookie.Value) {
        return nil, errors.New("sessionID already exists")
    }

    session := NewSession(s, cookieName)
    session.ID = s.NewSessionID()
    session.request = r

    return session, nil
}

func (s *Manager) Save(r *http.Request, w http.ResponseWriter, session *Session) error {
    s.database[session.ID] = session

    c := &http.Cookie{
        Name: session.Name(),
        Value: session.ID,
        Path: "/",
    }

    http.SetCookie(session.writer, c)
    return nil
}

func (s *Manager) Delete(sessionID string) {
    delete(s.database, sessionID)
}
