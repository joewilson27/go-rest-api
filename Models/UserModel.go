//Models/UserModel.go

package Models

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

/*
ref = https://stackoverflow.com/questions/45641999/parameter-before-the-function-name-in-go

It means that TableName() is a method on *User.
*/
func (b *User) TableName() string {
	return "users"
}
