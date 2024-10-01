package google_oauth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ClientId        = "ClientId"
	ClientSecret    = "ClientSecret"
	GoogleAdsScope  = "https://www.googleapis.com/auth/business.manage"
	GoogleOauth2URL = "https://accounts.google.com/o/oauth2/v2/auth"
	GoogleTokenURL  = "https://oauth2.googleapis.com/token"
	RedirectURI     = "http://localhost:3000/callback"
	State           = "state"
)

var instance = NewAuthService(
	GoogleOauth2URL,
	ClientId,
	ClientSecret,
	RedirectURI,
	GoogleAdsScope,
	State,
	GoogleTokenURL,
)

func TestAuthService_BuildOAuthUrl(t *testing.T) {
	got := instance.BuildOAuthUrl()
	require.NotNil(t, got)
	require.Equal(t, got, "https://accounts.google.com/o/oauth2/v2/auth?client_id=ClientId&redirect_uri=http://localhost:3000/callback&access_type=offline&response_type=code&scope=https://www.googleapis.com/auth/business.manage&state=state")
}
