package tokenuser

type TokenRequest struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
	Expires int64 `json:"expires"`
}

// func (t *TokenRequest) Validate(tokenId string) *errors.RestErr {
// 	usersRest
// }