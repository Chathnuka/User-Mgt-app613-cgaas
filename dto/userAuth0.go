package dto


type UserAuth0 struct {
	Email         string                 `json:"email" validate:"required"`
    Password      string                 `json:"password,omitempty" validate:"required"`
	Connection    string                 `json:"connection"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
	Name          string                 `json:"name"`
	EmailVerified bool                   `json:"email_verified"`
}