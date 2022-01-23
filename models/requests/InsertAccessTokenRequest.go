package requests

type InsertAccessTokenRequest struct {
	UserName string `json:"user_name"`
	Token    *string `json:"token"`
}
