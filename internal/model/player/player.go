// internal/model/player/player.go
package player

//import "fmt"

type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) SetName(name string) {
	if name != "" {
		p.name = name
	}
}

func (p *Player) String() string {
	return p.name
}