package requests

type GetTransactionRequest struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
	Page int    `json:"page,omitempty"`
}
