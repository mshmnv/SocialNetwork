package datastruct

type User struct {
	FirstName  string `db:"first_name" json:"first_name"`
	SecondName string `db:"second_name" json:"second_name"`
	Age        int64  `db:"age" json:"age"`
	BirthDate  string `db:"birthdate" json:"birthdate"`
	Biography  string `db:"biography" json:"biography"`
	City       string `db:"city" json:"city"`
	Password   string `db:"password"`
}

type LoginData struct {
	ID       uint64 `db:"id"`
	Password string `db:"password"`
}

func (u *User) GetUsersDBRecord() map[string]interface{} {
	return map[string]any{
		"first_name":  u.FirstName,
		"second_name": u.SecondName,
		"age":         u.Age,
		"birthdate":   u.BirthDate,
		"biography":   u.Biography,
		"city":        u.City,
		"password":    u.Password,
	}
}
