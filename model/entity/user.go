package entity

type UserRole string

const (
	ADMIN UserRole = "ADMIN"
	USER  UserRole = "USER"
)

type User struct {
	ProfilePicture string   `bson:"profile_picture"`
	Username       string   `bson:"username"`
	Password       string   `bson:"password"`
	Role           UserRole `bson:"role"`
}
