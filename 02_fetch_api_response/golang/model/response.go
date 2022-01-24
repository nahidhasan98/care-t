package model

type UserResponse struct {
	Meta interface{} `json:"meta"` // taking interfaces, [ignoring meta information right now, only focusing user data]
	Data []User      `json:"data"` // taking only user list
}

type DynamicResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"` // data field may be array or an object
}
