package txtcolor

import (
	"fmt"
)

//文字颜色
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

//Black 输出黑色字
func Black(str string) string {
	return textColor(TextBlack, str)
}

//Red 输出红色字
func Red(str string) string {
	return textColor(TextRed, str)
}

//Green 输出绿色字
func Green(str string) string {
	return textColor(TextGreen, str)
}

//Yellow 输出黄色字
func Yellow(str string) string {
	return textColor(TextYellow, str)
}

//Blue 输出蓝色字
func Blue(str string) string {
	return textColor(TextBlue, str)
}

//Magenta 输出紫色字
func Magenta(str string) string {
	return textColor(TextMagenta, str)
}

//Cyan 输出青色字
func Cyan(str string) string {
	return textColor(TextCyan, str)
}

//White 输出白色字
func White(str string) string {
	return textColor(TextWhite, str)
}

func textColor(color int, str string) string {
	// if IsWindows() {
	// 	return str
	// }

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

// func IsWindows() bool {
// 	if runtime.GOOS == "windows" {
// 		return true
// 	} else {
// 		return false
// 	}
// }
