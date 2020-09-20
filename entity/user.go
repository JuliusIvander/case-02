package entity

type (
	//User entity
	User struct {
		ID        string `json:"id,omitempty"`
		Nama      string `json:"nama"`
		Email     string `json:"email"`
		NomorHp   string `json:"nomorhp"`
		Pekerjaan string `json:"pekerjaan"`
	}
)
