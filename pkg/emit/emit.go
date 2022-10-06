package emit

import (
	"errors"
	"fmt"
	"github.com/micmonay/keybd_event"
	"github.com/tedli/shaman/pkg/common"
	"io"
	"os"
	"time"
)

var (
	characterSetKeyMapping = map[rune]*keybd_event.KeyBonding{
		'0':  newKeyBonding(keybd_event.VK_0),
		'1':  newKeyBonding(keybd_event.VK_1),
		'2':  newKeyBonding(keybd_event.VK_2),
		'3':  newKeyBonding(keybd_event.VK_3),
		'4':  newKeyBonding(keybd_event.VK_4),
		'5':  newKeyBonding(keybd_event.VK_5),
		'6':  newKeyBonding(keybd_event.VK_6),
		'7':  newKeyBonding(keybd_event.VK_7),
		'8':  newKeyBonding(keybd_event.VK_8),
		'9':  newKeyBonding(keybd_event.VK_9),
		'a':  newKeyBonding(keybd_event.VK_A),
		'b':  newKeyBonding(keybd_event.VK_B),
		'c':  newKeyBonding(keybd_event.VK_C),
		'd':  newKeyBonding(keybd_event.VK_D),
		'e':  newKeyBonding(keybd_event.VK_E),
		'f':  newKeyBonding(keybd_event.VK_F),
		'g':  newKeyBonding(keybd_event.VK_G),
		'h':  newKeyBonding(keybd_event.VK_H),
		'i':  newKeyBonding(keybd_event.VK_I),
		'j':  newKeyBonding(keybd_event.VK_J),
		'k':  newKeyBonding(keybd_event.VK_K),
		'l':  newKeyBonding(keybd_event.VK_L),
		'm':  newKeyBonding(keybd_event.VK_M),
		'n':  newKeyBonding(keybd_event.VK_N),
		'o':  newKeyBonding(keybd_event.VK_O),
		'p':  newKeyBonding(keybd_event.VK_P),
		'q':  newKeyBonding(keybd_event.VK_Q),
		'r':  newKeyBonding(keybd_event.VK_R),
		's':  newKeyBonding(keybd_event.VK_S),
		't':  newKeyBonding(keybd_event.VK_T),
		'u':  newKeyBonding(keybd_event.VK_U),
		'v':  newKeyBonding(keybd_event.VK_V),
		'w':  newKeyBonding(keybd_event.VK_W),
		'x':  newKeyBonding(keybd_event.VK_X),
		'y':  newKeyBonding(keybd_event.VK_Y),
		'z':  newKeyBonding(keybd_event.VK_Z),
		'A':  newKeyBonding(keybd_event.VK_A, true),
		'B':  newKeyBonding(keybd_event.VK_B, true),
		'C':  newKeyBonding(keybd_event.VK_C, true),
		'D':  newKeyBonding(keybd_event.VK_D, true),
		'E':  newKeyBonding(keybd_event.VK_E, true),
		'F':  newKeyBonding(keybd_event.VK_F, true),
		'G':  newKeyBonding(keybd_event.VK_G, true),
		'H':  newKeyBonding(keybd_event.VK_H, true),
		'I':  newKeyBonding(keybd_event.VK_I, true),
		'J':  newKeyBonding(keybd_event.VK_J, true),
		'K':  newKeyBonding(keybd_event.VK_K, true),
		'L':  newKeyBonding(keybd_event.VK_L, true),
		'M':  newKeyBonding(keybd_event.VK_M, true),
		'N':  newKeyBonding(keybd_event.VK_N, true),
		'O':  newKeyBonding(keybd_event.VK_O, true),
		'P':  newKeyBonding(keybd_event.VK_P, true),
		'Q':  newKeyBonding(keybd_event.VK_Q, true),
		'R':  newKeyBonding(keybd_event.VK_R, true),
		'S':  newKeyBonding(keybd_event.VK_S, true),
		'T':  newKeyBonding(keybd_event.VK_T, true),
		'U':  newKeyBonding(keybd_event.VK_U, true),
		'V':  newKeyBonding(keybd_event.VK_V, true),
		'W':  newKeyBonding(keybd_event.VK_W, true),
		'X':  newKeyBonding(keybd_event.VK_X, true),
		'Y':  newKeyBonding(keybd_event.VK_Y, true),
		'Z':  newKeyBonding(keybd_event.VK_Z, true),
		'!':  newKeyBonding(keybd_event.VK_1, true), // TODO: keymap related
		'#':  newKeyBonding(keybd_event.VK_3, true),
		'$':  newKeyBonding(keybd_event.VK_4, true),
		'%':  newKeyBonding(keybd_event.VK_5, true),
		'&':  newKeyBonding(keybd_event.VK_7, true),
		'(':  newKeyBonding(keybd_event.VK_9, true),
		')':  newKeyBonding(keybd_event.VK_0, true),
		'*':  newKeyBonding(keybd_event.VK_8, true),
		'+':  newKeyBonding(keybd_event.VK_EQUAL, true),
		'-':  newKeyBonding(keybd_event.VK_MINUS),
		';':  newKeyBonding(keybd_event.VK_SEMICOLON),
		'<':  newKeyBonding(keybd_event.VK_COMMA, true),
		'=':  newKeyBonding(keybd_event.VK_EQUAL),
		'>':  newKeyBonding(keybd_event.VK_DOT, true),
		'?':  newKeyBonding(keybd_event.VK_SLASH, true),
		'@':  newKeyBonding(keybd_event.VK_2, true),
		'^':  newKeyBonding(keybd_event.VK_6, true),
		'_':  newKeyBonding(keybd_event.VK_MINUS, true),
		'`':  newKeyBonding(keybd_event.VK_GRAVE),
		'{':  newKeyBonding(keybd_event.VK_LEFTBRACE, true),
		'|':  newKeyBonding(keybd_event.VK_BACKSLASH, true),
		'}':  newKeyBonding(keybd_event.VK_RIGHTBRACE, true),
		'~':  newKeyBonding(keybd_event.VK_GRAVE, true),
		'"':  newKeyBonding(keybd_event.VK_APOSTROPHE, true),
		'\'': newKeyBonding(keybd_event.VK_APOSTROPHE),
		',':  newKeyBonding(keybd_event.VK_COMMA),
		'.':  newKeyBonding(keybd_event.VK_DOT),
		'/':  newKeyBonding(keybd_event.VK_SLASH),
		':':  newKeyBonding(keybd_event.VK_SEMICOLON, true),
		'[':  newKeyBonding(keybd_event.VK_LEFTBRACE),
		'\\': newKeyBonding(keybd_event.VK_BACKSLASH),
		']':  newKeyBonding(keybd_event.VK_RIGHTBRACE),
		' ':  newKeyBonding(keybd_event.VK_SPACE),
	}

	enterKeyBonding = newKeyBonding(keybd_event.VK_ENTER)

	ErrUnsupportedCharacterExist = errors.New("unsupported character exist")
)

func newKeyBonding(key int, shift ...bool) *keybd_event.KeyBonding {
	bonding, _ := keybd_event.NewKeyBonding()
	bonding.SetKeys(key)
	for _, s := range shift {
		bonding.HasSHIFT(s)
		break
	}
	return &bonding
}

func emit(filePath, folder string, index, after int, out io.Writer) (err error) {
	var file string
	if file, err = common.GetFile(filePath, folder, index); err != nil {
		return
	}
	var content []byte
	if content, err = os.ReadFile(file); err != nil {
		return
	}
	<-time.After(time.Second * time.Duration(after))
	var lastRNIndex int
	for i, c := range content {
		key, supported := characterSetKeyMapping[rune(c)]
		if !supported {
			if c == 10 || c == 13 {
				if i == lastRNIndex+1 {
					lastRNIndex = 0
					continue
				} else {
					lastRNIndex = i
					key = enterKeyBonding
				}
			} else {
				err = ErrUnsupportedCharacterExist
				_, _ = fmt.Fprintf(out, "Unsupported: [%d]\n", int(c))
				return
			}
		}
		if err = key.Press(); err != nil {
			return
		}
		<-time.After(5 * time.Millisecond)
		if err = key.Release(); err != nil {
			return
		}
	}
	return
}
