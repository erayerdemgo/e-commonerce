package model

type UserSignup struct {
	Name     string `form:"name" validate:"required,min=3,max=20"`
	Surname  string `form:"surname" validate:"required,min=2,max=20"`
	Email    string `form:"email" validate:"email,checkdb"`    //  is exist email
	Password string `form:"password"  validate:"min=3,max=15"` // todo regex here
}

type UserSignin struct {
	Email    string `form:"email" validate:"email"`
	Password string `form:"password"  validate:"min=3,max=15"` // todo regex here
}
