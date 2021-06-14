package request

type SignUp struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (su SignUp) Validate() error {
	return nil
}

func (si SignIn) Validate() error {
	return nil
}
