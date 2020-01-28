package entities

type User struct {
	ID        uint
	Name      string
	Icon      string
	Profile   string
	Followers Followers
	Followees Followees
}

type Users = []*User

func NewUser(id uint, name string, icon string, profile string, followers Followers, followees Followees) *User {
	return &User{
		ID:        id,
		Name:      name,
		Icon:      icon,
		Profile:   profile,
		Followers: followers,
		Followees: followees,
	}
}

// ユーザが引数のuserIDにフォローされている場合はtrueを返す
func (u *User) IsFollowedBy(userID uint) bool {
	for _, f := range u.Followees {
		if f.ID == userID {
			return true
		}
	}
	return false
}

// TODO: 認証機能を実装したら削除する
const LoginUserID = 1
