package data

const (
	ModuleName = "identity"
)

type ModulePayload struct {
	RequestId string `json:"request_id"`
	UserId    string `json:"user_id"`
	Action    string `json:"action"`

	//other fields that are required for module
	Username *string `json:"username"`
	Phone    *string `json:"phone"`
}
