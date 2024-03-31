package message

import (
	"fmt"
)

type LogoutMessage struct {
	Player string
}

func NewLogoutMessage(player string) LogoutMessage {
	return LogoutMessage{Player: player}
}

func (m LogoutMessage) IsEmpty() bool {
	return m.Player == ""
}

func (m LogoutMessage) GetMessage() string {
	return fmt.Sprintf("%s left the game", m.Player)
}

func (m LogoutMessage) GetTitle() string {
	return "Goodbye"
}

func (m LogoutMessage) GetWHColor() string {
	return "0xFF8C00"
}
