package user

type GetUserProfileResponse struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
