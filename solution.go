package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {

	emojiMessage := emoji.Sprint("Hello :world_map:!")
	return emojiMessage
}
