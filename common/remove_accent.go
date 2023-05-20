package common

import (
	"strings"
)

var vietnameseToAlphabetSensitive = map[rune]rune{
	'Á': 'A',
	'á': 'a',
	'À': 'A',
	'à': 'a',
	'Ả': 'A',
	'ả': 'a',
	'Ã': 'A',
	'ã': 'a',
	'Ạ': 'A',
	'ạ': 'a',

	'Ă': 'A',
	'ă': 'a',
	'Ắ': 'A',
	'ắ': 'a',
	'Ằ': 'A',
	'ằ': 'a',
	'Ẳ': 'A',
	'ẳ': 'a',
	'Ẵ': 'A',
	'ẵ': 'a',
	'Ặ': 'A',
	'ặ': 'a',

	'Â': 'A',
	'â': 'a',
	'Ấ': 'A',
	'ấ': 'a',
	'Ầ': 'A',
	'ầ': 'a',
	'Ẩ': 'A',
	'ẩ': 'a',
	'Ẫ': 'A',
	'ẫ': 'a',
	'Ậ': 'A',
	'ậ': 'a',

	'Đ': 'D',
	'đ': 'd',

	'É': 'E',
	'é': 'e',
	'È': 'E',
	'è': 'e',
	'Ẻ': 'E',
	'ẻ': 'e',
	'Ẽ': 'E',
	'ẽ': 'e',
	'Ẹ': 'E',
	'ẹ': 'e',

	'Ê': 'E',
	'ê': 'e',
	'Ế': 'E',
	'ế': 'e',
	'Ề': 'E',
	'ề': 'e',
	'Ể': 'E',
	'ể': 'e',
	'Ễ': 'E',
	'ễ': 'e',
	'Ệ': 'E',
	'ệ': 'e',

	'Í': 'I',
	'í': 'i',
	'Ì': 'I',
	'ì': 'i',
	'Ỉ': 'I',
	'ỉ': 'i',
	'Ĩ': 'I',
	'ĩ': 'i',
	'Ị': 'I',
	'ị': 'i',

	'Ó': 'O',
	'ó': 'o',
	'Ò': 'O',
	'ò': 'o',
	'Ỏ': 'O',
	'ỏ': 'o',
	'Õ': 'O',
	'õ': 'o',
	'Ọ': 'O',
	'ọ': 'o',

	'Ô': 'O',
	'ô': 'o',
	'Ố': 'O',
	'ố': 'o',
	'Ồ': 'O',
	'ồ': 'o',
	'Ổ': 'O',
	'ổ': 'o',
	'Ỗ': 'O',
	'ỗ': 'o',
	'Ộ': 'O',
	'ộ': 'o',

	'Ơ': 'O',
	'ơ': 'o',
	'Ớ': 'O',
	'ớ': 'o',
	'Ờ': 'O',
	'ờ': 'o',
	'Ở': 'O',
	'ở': 'o',
	'Ỡ': 'O',
	'ỡ': 'o',
	'Ợ': 'O',
	'ợ': 'o',

	'Ú': 'U',
	'ú': 'u',
	'Ù': 'U',
	'ù': 'u',
	'Ủ': 'U',
	'ủ': 'u',
	'Ũ': 'U',
	'ũ': 'u',
	'Ụ': 'U',
	'ụ': 'u',

	'Ư': 'U',
	'ư': 'u',
	'Ứ': 'U',
	'ứ': 'u',
	'Ừ': 'U',
	'ừ': 'u',
	'Ử': 'U',
	'ử': 'u',
	'Ữ': 'U',
	'ữ': 'u',
	'Ự': 'U',
	'ự': 'u',

	'Ý': 'Y',
	'ý': 'y',
	'Ỳ': 'Y',
	'ỳ': 'y',
	'Ỷ': 'Y',
	'ỷ': 'y',
	'Ỹ': 'Y',
	'ỹ': 'y',
	'Ỵ': 'Y',
	'ỵ': 'y',
}

func RemoveVietnameseAccent(text string) string {
	runeSlice := []rune(strings.TrimSpace(text))
	var sb strings.Builder
	sb.Grow(len(runeSlice))

	for _, character := range runeSlice {
		if alphabet, ok := vietnameseToAlphabetSensitive[character]; ok {
			sb.Write([]byte{string(alphabet)[0]})
		} else {
			sb.Write([]byte{string(character)[0]})
		}
	}

	return strings.ToLower(sb.String())
}
