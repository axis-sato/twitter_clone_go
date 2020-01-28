package entities

type Followee struct {
	ID      uint
	Name    string
	Icon    string
	Profile string
}

type Followees = []*Followee

func NewFollowee(id uint, name string, icon string, profile string) *Followee {
	return &Followee{
		ID:      id,
		Name:    name,
		Icon:    icon,
		Profile: profile,
	}
}
