package entities

type Follower struct {
	ID      uint
	Name    string
	Icon    string
	Profile string
}

type Followers = []*Follower

func NewFollower(id uint, name string, icon string, profile string) *Follower {
	return &Follower{
		ID:      id,
		Name:    name,
		Icon:    icon,
		Profile: profile,
	}
}
