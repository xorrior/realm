package auth

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"realm.pub/tavern/internal/ent"
	"realm.pub/tavern/internal/ent/user"
)

const (
	// SessionCookieName is the name of the cookie used to store the user's session token.
	SessionCookieName = "auth-session"
)

var (
	ErrInvalidCredentials = fmt.Errorf("invalid username or password")
	ErrSetupNotAllowed    = fmt.Errorf("setup is only allowed when no users exist")
	ErrMissingFields      = fmt.Errorf("username and password are required")
	ErrPasswordTooShort   = fmt.Errorf("password must be at least 8 characters")
)

// HashPassword creates a bcrypt hash of the given password.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

// CheckPassword compares a plaintext password against a bcrypt hash.
func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// NewPasswordLoginHandler returns an HTTP handler that authenticates users with username and password.
// On GET, it returns the login page status (whether setup is needed).
// On POST, it validates credentials and sets a session cookie.
func NewPasswordLoginHandler(graph *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			// Check if any users exist to determine if setup is needed
			count := graph.User.Query().CountX(req.Context())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"needsSetup": count == 0,
			})
			return
		}

		if req.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(req.Body).Decode(&creds); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if creds.Username == "" || creds.Password == "" {
			http.Error(w, ErrMissingFields.Error(), http.StatusBadRequest)
			return
		}

		// Lookup user by name
		usr, err := graph.User.Query().
			Where(user.Name(creds.Username)).
			Only(req.Context())
		if err != nil {
			slog.WarnContext(req.Context(), "login failed: user not found", "username", creds.Username)
			http.Error(w, ErrInvalidCredentials.Error(), http.StatusUnauthorized)
			return
		}

		// Check password
		if usr.PasswordHash == nil || !CheckPassword(*usr.PasswordHash, creds.Password) {
			slog.WarnContext(req.Context(), "login failed: invalid password", "username", creds.Username)
			http.Error(w, ErrInvalidCredentials.Error(), http.StatusUnauthorized)
			return
		}

		// Set session cookie
		http.SetCookie(w, &http.Cookie{
			Name:     SessionCookieName,
			Value:    usr.SessionToken,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().AddDate(0, 1, 0),
		})

		slog.InfoContext(req.Context(), "user logged in", "user_id", usr.ID, "user_name", usr.Name, "is_admin", usr.IsAdmin)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
		})
	})
}

// NewSetupHandler returns an HTTP handler for initial user setup (TOFU model).
// Only works when no users exist in the database.
func NewSetupHandler(graph *ent.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			count := graph.User.Query().CountX(req.Context())
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]interface{}{
				"needsSetup": count == 0,
			})
			return
		}

		if req.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Only allow setup when no users exist
		count := graph.User.Query().CountX(req.Context())
		if count > 0 {
			http.Error(w, ErrSetupNotAllowed.Error(), http.StatusForbidden)
			return
		}

		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(req.Body).Decode(&creds); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if creds.Username == "" || creds.Password == "" {
			http.Error(w, ErrMissingFields.Error(), http.StatusBadRequest)
			return
		}

		if len(creds.Password) < 8 {
			http.Error(w, ErrPasswordTooShort.Error(), http.StatusBadRequest)
			return
		}

		if len(creds.Username) < 3 || len(creds.Username) > 25 {
			http.Error(w, "username must be between 3 and 25 characters", http.StatusBadRequest)
			return
		}

		// Hash the password
		hash, err := HashPassword(creds.Password)
		if err != nil {
			slog.ErrorContext(req.Context(), "setup failed to hash password", "err", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		// Create the first user as admin + activated
		usr := graph.User.Create().
			SetName(creds.Username).
			SetPasswordHash(hash).
			SetPhotoURL("").
			SetIsAdmin(true).
			SetIsActivated(true).
			SaveX(req.Context())

		// Set session cookie
		http.SetCookie(w, &http.Cookie{
			Name:     SessionCookieName,
			Value:    usr.SessionToken,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().AddDate(0, 1, 0),
		})

		slog.InfoContext(req.Context(), "initial admin user created via setup", "user_id", usr.ID, "user_name", usr.Name)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
		})
	})
}
