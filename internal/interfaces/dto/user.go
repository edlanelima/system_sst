package dto

type User struct {
	FullName    string `json:"fullname" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	IsSstRole   bool   `json:"is_sst_role" validate:"required_without=IsAdminRole"`
	IsAdminRole bool   `json:"is_admin_role" validate:"required_without=IsSstRole"`
}
