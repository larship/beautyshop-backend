package models

type Client struct {
	Uuid             string `json:"uuid"`
	FullName         string `json:"fullName"`
	Phone            string `json:"phone"`
	SessionId        string `json:"sessionId"`
	SessionPrivateId string `json:"-"`
	Salt             string `json:"salt"`
}
