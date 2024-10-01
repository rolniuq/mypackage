package google_oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OAuth2Data struct {
	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type AuthService struct {
	googleOauth2URL string
	clientId        string
	clientSecret    string
	redirectURI     string
	googleAdsScope  string
	state           string
	googleTokenURL  string
}

func NewAuthService(
	googleOauth2URL,
	clientId,
	clientSecret,
	redirectURI,
	googleAdsScope,
	state,
	googleTokenURL string,
) *AuthService {
	return &AuthService{
		googleOauth2URL: googleOauth2URL,
		clientId:        clientId,
		clientSecret:    clientSecret,
		redirectURI:     redirectURI,
		googleAdsScope:  googleAdsScope,
		state:           state,
		googleTokenURL:  googleTokenURL,
	}
}

func (a *AuthService) BuildOAuthUrl() string {
	return fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&access_type=offline&response_type=code&scope=%s&state=%s",
		a.googleOauth2URL,
		a.clientId,
		a.redirectURI,
		a.googleAdsScope,
		a.state,
	)
}

func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, a.BuildOAuthUrl(), http.StatusTemporaryRedirect)
}

func (a *AuthService) GetToken(code string) (*OAuth2Data, error) {
	postBody, _ := json.Marshal(
		map[string]any{
			"code":          code,
			"client_id":     a.clientId,
			"client_secret": a.clientSecret,
			"redirect_uri":  a.redirectURI,
			"grant_type":    "authorization_code",
		},
	)

	body := bytes.NewBuffer(postBody)
	resp, err := http.Post(a.googleTokenURL, "application/json", body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get access token with return code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token OAuth2Data
	err = json.Unmarshal(respData, &token)
	if err != nil {
		return nil, err
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("failed to get access token")
	}
	return &token, nil
}

func (a *AuthService) Callback(w http.ResponseWriter, r *http.Request) (*OAuth2Data, error) {
	code := r.URL.Query().Get("code")

	token, err := a.GetToken(code)
	if err != nil {
		return nil, err
	}

	return token, nil
}
