package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GET /auth/login
func MicrosoftLogin(c *gin.Context) {
	clientID := os.Getenv("MICROSOFT_CLIENT_ID")
	redirectURI := os.Getenv("MICROSOFT_REDIRECT_URI")
	tenantID := os.Getenv("MICROSOFT_TENANT_ID")
	scope := url.QueryEscape("User.Read openid email profile")
	authURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/authorize?client_id=%s&response_type=code&redirect_uri=%s&response_mode=query&scope=%s", tenantID, clientID, url.QueryEscape(redirectURI), scope)
	c.Redirect(http.StatusFound, authURL)
}

// GET /auth/callback
func MicrosoftCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code in callback"})
		return
	}
	clientID := os.Getenv("MICROSOFT_CLIENT_ID")
	clientSecret := os.Getenv("MICROSOFT_CLIENT_SECRET")
	redirectURI := os.Getenv("MICROSOFT_REDIRECT_URI")
	tenantID := os.Getenv("MICROSOFT_TENANT_ID")

	tokenURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenantID)
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("scope", "User.Read openid email profile")
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)
	data.Set("grant_type", "authorization_code")
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token request"})
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token from Microsoft"})
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	json.Unmarshal(body, &tokenResp)
	if tokenResp.AccessToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No access token received"})
		return
	}

	// Get user info from Microsoft Graph
	userReq, _ := http.NewRequest("GET", "https://graph.microsoft.com/v1.0/me", nil)
	userReq.Header.Set("Authorization", "Bearer "+tokenResp.AccessToken)
	userResp, err := http.DefaultClient.Do(userReq)
	if err != nil || userResp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info from Microsoft"})
		return
	}
	defer userResp.Body.Close()
	userBody, _ := ioutil.ReadAll(userResp.Body)
	var msUser struct {
		ID                string `json:"id"`
		Mail              string `json:"mail"`
		UserPrincipalName string `json:"userPrincipalName"`
		DisplayName       string `json:"displayName"`
	}
	json.Unmarshal(userBody, &msUser)
	userEmail := msUser.Mail
	if userEmail == "" {
		userEmail = msUser.UserPrincipalName
	}
	if userEmail == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No email found in Microsoft account"})
		return
	}

	// Create or update user in DB
	db := models.DB
	var user models.User
	db.Where("email = ?", userEmail).First(&user)
	if user.ID == 0 {
		user = models.User{
			Email: userEmail,
			Name:  msUser.DisplayName,
		}
		db.Create(&user)
	} else {
		user.Name = msUser.DisplayName
		db.Save(&user)
	}

	// Issue JWT
	secret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign JWT"})
		return
	}

	// Redirect to frontend with token and user info in query string
	frontendURL := os.Getenv("APP_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	redirectURL := fmt.Sprintf("%s/login?token=%s&name=%s&id=%d", frontendURL, url.QueryEscape(tokenString), url.QueryEscape(user.Name), user.ID)
	c.Redirect(http.StatusFound, redirectURL)
}

// GET /auth/logout
func Logout(c *gin.Context) {
	// For JWT, logout is handled on the client by deleting the token
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
