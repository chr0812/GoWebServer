package model

type Token struct {
	ResultCode   string `json:"resultCode"`
	ResultMsg    string `json:"resultMsg"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	EmailId      string `json:"emailId"`
}
