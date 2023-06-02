package entity

type UserRole string

const (
	ADMIN UserRole = "ADMIN"
	USER  UserRole = "USER"
)

type User struct {
	Id       string   `bson:"id"`
	Username string   `bson:"username"`
	Password string   `bson:"password"`
	Role     UserRole `bson:"role"`
}
