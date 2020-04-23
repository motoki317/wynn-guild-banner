package banner

import "image/color"

var Mask = color.RGBA{R: 0, G: 0, B: 0, A: 20}
var Background = color.RGBA{R: 150, G: 150, B: 100, A: 0xff}

var (
	Black     = color.RGBA{R: 0x19, G: 0x19, B: 0x19, A: 0xff}
	Red       = color.RGBA{R: 0x99, G: 0x33, B: 0x33, A: 0xff}
	Green     = color.RGBA{R: 0x66, G: 0x7F, B: 0x33, A: 0xff}
	Brown     = color.RGBA{R: 0x66, G: 0x4C, B: 0x33, A: 0xff}
	Blue      = color.RGBA{R: 0x33, G: 0x4C, B: 0xB2, A: 0xff}
	Purple    = color.RGBA{R: 0x7F, G: 0x3F, B: 0xB2, A: 0xff}
	Cyan      = color.RGBA{R: 0x4C, G: 0x7F, B: 0x99, A: 0xff}
	LightGray = color.RGBA{R: 0x99, G: 0x99, B: 0x99, A: 0xff}
	Gray      = color.RGBA{R: 0x4C, G: 0x4C, B: 0x4C, A: 0xff}
	Pink      = color.RGBA{R: 0xF2, G: 0x7F, B: 0xA5, A: 0xff}
	Lime      = color.RGBA{R: 0x7F, G: 0xCC, B: 0x19, A: 0xff}
	Yellow    = color.RGBA{R: 0xE5, G: 0xE5, B: 0x33, A: 0xff}
	LightBlue = color.RGBA{R: 0x66, G: 0x99, B: 0xD8, A: 0xff}
	Magenta   = color.RGBA{R: 0xB2, G: 0x4C, B: 0xD8, A: 0xff}
	Orange    = color.RGBA{R: 0xD8, G: 0x7F, B: 0x33, A: 0xff}
	White     = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xff}
	Silver    = color.RGBA{R: 0x99, G: 0x99, B: 0x99, A: 0xff}
)

var colorMap map[string]color.RGBA

func init() {
	colorMap = make(map[string]color.RGBA)
	colorMap["BLACK"] = Black
	colorMap["RED"] = Red
	colorMap["GREEN"] = Green
	colorMap["BROWN"] = Brown
	colorMap["BLUE"] = Blue
	colorMap["PURPLE"] = Purple
	colorMap["CYAN"] = Cyan
	colorMap["LIGHT_GRAY"] = LightGray
	colorMap["GRAY"] = Gray
	colorMap["PINK"] = Pink
	colorMap["LIME"] = Lime
	colorMap["YELLOW"] = Yellow
	colorMap["LIGHT_BLUE"] = LightBlue
	colorMap["MAGENTA"] = Magenta
	colorMap["ORANGE"] = Orange
	colorMap["WHITE"] = White
	colorMap["SILVER"] = Silver
}

func GetColor(col string) (c color.RGBA, ok bool) {
	c, ok = colorMap[col]
	return
}

func getMaskColor(col color.RGBA) color.RGBA {
	col.A = 180
	return col
}
