package exercises

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Email struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	From    string `json:"from"`
	To      string `json:"to"`
}

type EmailResponse struct {
	Emails        []Email `json:"emails"`
	NextPageToken string  `json:"nextPageToken"`
}

func Exercise9() {
	r := mux.NewRouter()

	// Middleware to handle authentication
	r.Use(authMiddleware)

	r.HandleFunc("/emails", getEmailsHandler).Methods("GET")
	r.HandleFunc("/emails", createEmailHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if !validateToken(token) {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getEmailsHandler(w http.ResponseWriter, r *http.Request) {
	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10 // Default page size
	}

	pageToken := r.URL.Query().Get("pageToken")
	emails, nextPageToken, err := getEmails(pageSize, pageToken)
	if err != nil {
		log.Printf("Failed to get emails: %v\n", err)
		http.Error(w, "Failed to get emails", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := EmailResponse{
		Emails:        emails,
		NextPageToken: nextPageToken,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func createEmailHandler(w http.ResponseWriter, r *http.Request) {
	var email Email
	if err := json.NewDecoder(r.Body).Decode(&email); err != nil {
		log.Printf("Invalid input: %v\n", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := createEmail(email); err != nil {
		log.Printf("Failed to create email: %v\n", err)
		http.Error(w, "Failed to create email", http.StatusInternalServerError)
		return
	}

	log.Printf("Email created: %+v\n", email)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(email); err != nil {
		log.Printf("Failed to encode response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func getEmails(pageSize int, pageToken string) ([]Email, string, error) {
	mu.Lock()
	defer mu.Unlock()

	start := 0
	if pageToken != "" {
		for i, email := range mockEmails {
			if email.ID == pageToken {
				start = i + 1
				break
			}
		}
	}

	end := start + pageSize
	if end > len(mockEmails) {
		end = len(mockEmails)
	}

	nextPageToken := ""
	if end < len(mockEmails) {
		nextPageToken = mockEmails[end].ID
	}

	return mockEmails[start:end], nextPageToken, nil
}

func createEmail(email Email) error {
	mu.Lock()
	defer mu.Unlock()

	for _, existingEmail := range mockEmails {
		if existingEmail.ID == email.ID {
			return errors.New("email with this ID already exists")
		}
	}

	mockEmails = append(mockEmails, email)
	log.Printf("Email created: %+v\n", email)
	return nil
}

var mockEmails = []Email{
	{ID: "1", Subject: "Hello", Body: "Hello World!", From: "alice@example.com", To: "bob@example.com"},
	{ID: "2", Subject: "Hi", Body: "Hi Bob!", From: "carol@example.com", To: "bob@example.com"},
}

var jwtSecret = os.Getenv("JWT_SECRET")

func validateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		log.Println("Token validation error:", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["exp"].(float64))
		if exp < time.Now().Unix() {
			log.Println("Token is expired")
			return false
		}
		return true
	}

	return false
}

func generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
