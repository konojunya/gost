package model

type Result struct {
	GistURL     string `frame:"Gist's URL"`
	Description string
	Public      bool   `frame:"Is the published gist?"`
	FilePath    string `frame:"Local file path"`
}

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AuthURL      string `json:"authorize_url"`
	TokenURL     string `json:"token_url"`
	RedirectURL  string `json:"redirect_url"`
}
