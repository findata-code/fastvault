package controllers

type SecretCreatedResponse struct {
	Token string `json:"token"`
}

type ApplicationErrorResponse struct {
	Error string `json:"error"`
}