package model

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type GetRequest struct {
	Name string `json:"name"`
	//
	// add your data
	//
}

type AddRequest struct {
	User string `json:"user" binding:"required"`
	//
	// add your data
	//
}

type UpdateRequest struct {
	ID   int
	User string `json:"user" binding:"required"`
	//
	// add your data
	//
}

type DeleteRequest struct {
	ID   int
	User string `json:"user" binding:"required"`
	//
	// add your data
	//
}
