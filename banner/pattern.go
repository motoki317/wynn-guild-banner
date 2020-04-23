package banner

import (
	"image/color"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

var patterns map[string]func(b *Banner, col color.RGBA)

func init() {
	patterns = make(map[string]func(*Banner, color.RGBA))

	patterns["SQUARE_TOP_LEFT"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(0, 0, 90, 130)
		b.setColor(Mask)
		b.fillRect(0, 0, 90, 10)
		b.fillRect(0, 10, 10, 120)
		b.setColor(getMaskColor(col))
		b.fillRect(90, 0, 10, 130)
		b.fillRect(0, 130, 100, 10)
	}

	patterns["SQUARE_TOP_RIGHT"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(110, 0, 90, 130)
		b.setColor(Mask)
		b.fillRect(110, 0, 90, 10)
		b.fillRect(190, 10, 10, 120)
		b.setColor(getMaskColor(col))
		b.fillRect(100, 0, 10, 130)
		b.fillRect(100, 130, 100, 10)
	}

	patterns["SQUARE_BOTTOM_LEFT"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(0, 270, 90, 130)
		b.setColor(Mask)
		b.fillRect(0, 390, 90, 10)
		b.fillRect(0, 270, 10, 120)
		b.setColor(getMaskColor(col))
		b.fillRect(90, 270, 10, 130)
		b.fillRect(0, 260, 100, 10)
	}

	patterns["SQUARE_BOTTOM_RIGHT"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(110, 270, 90, 130)
		b.setColor(Mask)
		b.fillRect(110, 390, 90, 10)
		b.fillRect(190, 270, 10, 120)
		b.setColor(getMaskColor(col))
		b.fillRect(100, 270, 10, 130)
		b.fillRect(100, 260, 100, 10)
	}

	patterns["STRIPE_BOTTOM"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(0, 270, 200, 130)
		b.setColor(Mask)
		b.fillRect(0, 390, 200, 10)
		b.fillRect(0, 270, 10, 120)
		b.fillRect(190, 270, 10, 120)
		b.setColor(getMaskColor(col))
		b.fillRect(0, 260, 200, 10)
	}

	patterns["STRIPE_TOP"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(0, 0, 200, 130)
		b.setColor(Mask)
		b.fillRect(0, 0, 200, 10)
		b.fillRect(0, 10, 10, 120)
		b.fillRect(190, 10, 10, 120)
		b.setColor(getMaskColor(col))
		b.fillRect(0, 130, 200, 10)
	}

	patterns["STRIPE_LEFT"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(0, 0, 60, 400)
		b.setColor(Mask)
		b.fillRect(0, 0, 10, 400)
		b.fillRect(10, 0, 50, 10)
		b.fillRect(10, 390, 50, 10)
		b.setColor(getMaskColor(col))
		b.fillRect(60, 0, 10, 400)
	}

	patterns["STRIPE_RIGHT"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(140, 0, 60, 400)
		b.setColor(Mask)
		b.fillRect(190, 0, 10, 400)
		b.fillRect(140, 0, 50, 10)
		b.fillRect(140, 390, 50, 10)
		b.setColor(getMaskColor(col))
		b.fillRect(130, 0, 10, 400)
	}

	patterns["STRIPE_CENTER"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(80, 0, 40, 400)
		b.setColor(Mask)
		b.fillRect(80, 0, 40, 10)
		b.fillRect(80, 390, 40, 10)
		b.setColor(getMaskColor(col))
		b.fillRect(70, 0, 10, 400)
		b.fillRect(120, 0, 10, 400)
	}

	patterns["STRIPE_MIDDLE"] = func(b *Banner, col color.RGBA) {
		b.setColor(col)
		b.fillRect(0, 180, 200, 40)
		b.setColor(Mask)
		b.fillRect(0, 180, 10, 40)
		b.fillRect(190, 180, 10, 40)
		b.setColor(getMaskColor(col))
		b.fillRect(0, 170, 200, 10)
		b.fillRect(0, 220, 200, 10)
	}

	patterns["STRIPE_DOWNRIGHT"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 18; i++ {
			b.fillRect(10*i, 20*i, 30, 60)
		}
		b.setColor(Mask)
		b.fillRect(0, 0, 10, 60)
		b.fillRect(10, 0, 20, 10)
		b.fillRect(190, 340, 10, 50)
		b.fillRect(170, 390, 30, 10)
		b.setColor(getMaskColor(col))
		for i := 0; i < 17; i++ {
			b.fillRect(10*i, 60+20*i, 10, 20)
			b.fillRect(30+10*i, 20*i, 10, 20)
		}
	}

	patterns["STRIPE_DOWNLEFT"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 18; i++ {
			b.fillRect(170-10*i, 20*i, 30, 60)
		}
		b.setColor(Mask)
		b.fillRect(190, 0, 10, 60)
		b.fillRect(170, 0, 20, 10)
		b.fillRect(0, 340, 10, 50)
		b.fillRect(0, 390, 30, 10)
		b.setColor(getMaskColor(col))
		for i := 0; i < 17; i++ {
			b.fillRect(190-10*i, 60+20*i, 10, 20)
			b.fillRect(160-10*i, 20*i, 10, 20)
		}
	}

	patterns["STRIPE_SMALL"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 4; i++ {
			b.setColor(col)
			b.fillRect(20+50*i, 0, 10, 400)
			b.setColor(Mask)
			b.fillRect(20+50*i, 0, 10, 10)
			b.fillRect(20+50*i, 390, 10, 10)
			b.setColor(getMaskColor(col))
			b.fillRect(10+50*i, 0, 10, 400)
			b.fillRect(30+50*i, 0, 10, 400)
		}
	}

	patterns["CROSS"] = func(b *Banner, col color.RGBA) {
		j := 0
		for i := 0; i < 19; i++ {
			j = 180 - abs(180-20*i)
			b.fillRect(10*i, j, 20, 40)
			b.fillRect(10*i, 360-j, 20, 40)
		}
		b.fillRect(80, 180, 40, 60)
		b.setColor(getMaskColor(col))

		j = 0
		for i := 0; i < 16; i++ {
			j = 150 - abs(150-20*i)
			b.fillRect(20+10*i, j, 10, 20)
			b.fillRect(20+10*i, 380-j, 10, 20)
		}

		j = 0
		for i := 0; i < 16; i++ {
			j = 75 - abs(75-10*i)
			b.fillRect(j, 40+20*i, 10, 20)
			b.fillRect(190-j, 40+20*i, 10, 20)
		}

		b.setColor(Mask)

		b.fillRect(0, 0, 20, 10)
		b.fillRect(0, 10, 10, 30)
		b.fillRect(0, 360, 10, 30)
		b.fillRect(0, 390, 20, 10)
		b.fillRect(180, 0, 20, 10)
		b.fillRect(190, 10, 10, 30)
		b.fillRect(180, 390, 20, 10)
		b.fillRect(190, 360, 10, 30)
	}

	patterns["STRAIGHT_CROSS"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))
		b.fillRect(80, 0, 40, 400)
		b.fillRect(0, 180, 80, 40)
		b.fillRect(120, 180, 80, 40)

		b.setColor(col)
		b.fillRect(90, 0, 20, 400)
		b.fillRect(0, 190, 200, 20)

		b.setColor(Mask)
		b.fillRect(0, 190, 10, 20)
		b.fillRect(90, 0, 20, 10)
		b.fillRect(190, 190, 10, 20)
		b.fillRect(90, 390, 20, 10)
	}

	patterns["BORDER"] = func(b *Banner, col color.RGBA) {
		b.fillRect(0, 0, 200, 20)
		b.fillRect(0, 380, 200, 20)
		b.fillRect(0, 20, 20, 360)
		b.fillRect(180, 20, 20, 360)

		b.setColor(getMaskColor(col))
		b.fillRect(20, 20, 160, 10)
		b.fillRect(20, 370, 160, 10)
		b.fillRect(20, 30, 10, 340)
		b.fillRect(170, 30, 10, 340)

		b.setColor(Mask)
		b.fillRect(0, 0, 200, 10)
		b.fillRect(0, 390, 200, 10)
		b.fillRect(0, 10, 10, 380)
		b.fillRect(190, 10, 10, 380)
	}

	patterns["CURLY_BORDER"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))
		b.fillRect(0, 0, 200, 10)
		b.fillRect(0, 0, 10, 400)
		b.fillRect(190, 0, 10, 400)
		b.fillRect(0, 390, 200, 10)

		b.fillRect(0, 0, 20, 90)
		b.fillRect(0, 0, 40, 70)
		b.fillRect(0, 0, 50, 60)
		b.fillRect(50, 0, 10, 50)
		b.fillRect(0, 0, 70, 40)
		b.fillRect(0, 0, 90, 20)

		b.setColor(col)
		b.fillRect(0, 0, 10, 90)
		b.fillRect(0, 0, 20, 70)
		b.fillRect(0, 0, 40, 60)
		b.fillRect(0, 0, 60, 40)
		b.fillRect(0, 0, 70, 20)
		b.fillRect(0, 0, 90, 10)

		b.setColor(getMaskColor(col))
		b.fillRect(180, 0, 20, 90)
		b.fillRect(160, 0, 40, 70)
		b.fillRect(150, 0, 50, 60)
		b.fillRect(140, 0, 10, 50)
		b.fillRect(130, 0, 70, 40)
		b.fillRect(110, 0, 90, 20)

		b.setColor(col)
		b.fillRect(190, 0, 10, 90)
		b.fillRect(180, 0, 20, 70)
		b.fillRect(160, 0, 40, 60)
		b.fillRect(140, 0, 60, 40)
		b.fillRect(130, 0, 70, 20)
		b.fillRect(110, 0, 90, 10)

		b.setColor(getMaskColor(col))
		b.fillRect(0, 310, 20, 90)
		b.fillRect(0, 330, 40, 70)
		b.fillRect(0, 340, 50, 60)
		b.fillRect(50, 350, 10, 50)
		b.fillRect(0, 360, 70, 40)
		b.fillRect(0, 380, 90, 20)

		b.setColor(col)
		b.fillRect(0, 310, 10, 90)
		b.fillRect(0, 330, 20, 70)
		b.fillRect(0, 340, 40, 60)
		b.fillRect(0, 360, 60, 40)
		b.fillRect(0, 380, 70, 20)
		b.fillRect(0, 390, 90, 10)

		b.setColor(getMaskColor(col))
		b.fillRect(180, 310, 20, 90)
		b.fillRect(160, 330, 40, 70)
		b.fillRect(150, 340, 50, 60)
		b.fillRect(140, 350, 10, 50)
		b.fillRect(130, 360, 70, 40)
		b.fillRect(110, 380, 90, 20)

		b.setColor(col)
		b.fillRect(190, 310, 10, 90)
		b.fillRect(180, 330, 20, 70)
		b.fillRect(160, 340, 40, 60)
		b.fillRect(140, 360, 60, 40)
		b.fillRect(130, 380, 70, 20)
		b.fillRect(110, 390, 90, 10)

		b.setColor(getMaskColor(col))
		b.fillRect(0, 110, 20, 80)
		b.fillRect(0, 130, 40, 40)
		b.fillRect(0, 140, 50, 20)

		b.setColor(col)
		b.fillRect(0, 110, 10, 80)
		b.fillRect(0, 130, 20, 40)
		b.fillRect(0, 140, 40, 20)

		b.setColor(getMaskColor(col))
		b.fillRect(180, 110, 20, 80)
		b.fillRect(160, 130, 40, 40)
		b.fillRect(150, 140, 50, 20)

		b.setColor(col)
		b.fillRect(190, 110, 10, 80)
		b.fillRect(180, 130, 20, 40)
		b.fillRect(160, 140, 40, 20)

		b.setColor(getMaskColor(col))
		b.fillRect(0, 210, 20, 80)
		b.fillRect(0, 230, 40, 40)
		b.fillRect(0, 240, 50, 20)

		b.setColor(col)
		b.fillRect(0, 210, 10, 80)
		b.fillRect(0, 230, 20, 40)
		b.fillRect(0, 240, 40, 20)

		b.setColor(getMaskColor(col))
		b.fillRect(180, 210, 20, 80)
		b.fillRect(160, 230, 40, 40)
		b.fillRect(150, 240, 50, 20)

		b.setColor(col)
		b.fillRect(190, 210, 10, 80)
		b.fillRect(180, 230, 20, 40)
		b.fillRect(160, 240, 40, 20)

		b.setColor(Mask)
		b.fillRect(0, 0, 90, 10)
		b.fillRect(0, 10, 10, 80)
		b.fillRect(0, 110, 10, 80)

		b.fillRect(110, 0, 90, 10)
		b.fillRect(190, 10, 10, 80)
		b.fillRect(190, 110, 10, 80)

		b.fillRect(0, 390, 90, 10)
		b.fillRect(0, 310, 10, 80)
		b.fillRect(0, 210, 10, 80)

		b.fillRect(110, 390, 90, 10)
		b.fillRect(190, 310, 10, 80)
		b.fillRect(190, 210, 10, 80)
	}

	patterns["TRIANGLE_BOTTOM"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 10; i++ {
			b.setColor(col)
			b.fillRect(10*i, 390-20*i,
				20*(10-i), 10+20*i)
			b.setColor(getMaskColor(col))
			if i < 9 {
				b.fillRect(10*i, 370-20*i, 10, 20)
				b.fillRect(190-10*i, 370-20*i, 10, 20)
			} else {
				b.fillRect(10*i, 380-20*i, 10, 10)
				b.fillRect(190-10*i, 380-20*i, 10, 10)
			}
		}
		b.setColor(Mask)
		b.fillRect(0, 390, 200, 10)
	}

	patterns["TRIANGLE_TOP"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 10; i++ {
			b.setColor(col)
			b.fillRect(10*i, 0,
				20*(10-i), 10+20*i)
			b.setColor(getMaskColor(col))
			if i < 9 {
				b.fillRect(10*i, 10+20*i, 10, 20)
				b.fillRect(190-10*i, 10+20*i, 10, 20)
			} else {
				b.fillRect(10*i, 10+20*i, 10, 10)
				b.fillRect(190-10*i, 10+20*i, 10, 10)
			}
		}
		b.setColor(Mask)
		b.fillRect(0, 0, 200, 10)
	}

	patterns["TRIANGLES_BOTTOM"] = func(b *Banner, col color.RGBA) {
		b.fillRect(0, 390, 200, 10)
		for i := 0; i < 3; i++ {
			b.setColor(col)
			b.fillRect(70*i, 380, 60, 20)
			b.fillRect(10+70*i, 360, 40, 20)
			b.fillRect(20+70*i, 340, 20, 20)
			b.setColor(getMaskColor(col))
			b.fillRect(70*i, 360, 10, 20)
			b.fillRect(10+70*i, 340, 10, 20)
			b.fillRect(20+70*i, 330, 20, 10)
			b.fillRect(40+70*i, 340, 10, 20)
			b.fillRect(50+70*i, 360, 10, 20)
		}
		b.fillRect(60, 380, 10, 10)
		b.fillRect(130, 380, 10, 10)
		b.setColor(Mask)
		b.fillRect(0, 390, 200, 10)
	}

	patterns["TRIANGLES_TOP"] = func(b *Banner, col color.RGBA) {
		b.fillRect(0, 0, 200, 10)
		for i := 0; i < 3; i++ {
			b.setColor(col)
			b.fillRect(70*i, 0, 60, 20)
			b.fillRect(10+70*i, 20, 40, 20)
			b.fillRect(20+70*i, 40, 20, 20)
			b.setColor(getMaskColor(col))
			b.fillRect(70*i, 20, 10, 20)
			b.fillRect(10+70*i, 40, 10, 20)
			b.fillRect(20+70*i, 60, 20, 10)
			b.fillRect(40+70*i, 40, 10, 20)
			b.fillRect(50+70*i, 20, 10, 20)
		}
		b.fillRect(60, 10, 10, 10)
		b.fillRect(130, 10, 10, 10)
		b.setColor(Mask)
		b.fillRect(0, 0, 200, 10)
	}

	patterns["DIAGONAL_LEFT"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 19; i++ {
			b.fillRect(0, 0, 10+i*10, 380-i*20)
		}
		b.setColor(getMaskColor(col))
		for i := 0; i < 20; i++ {
			b.fillRect(0+i*10, 0+380-i*20, 10, 20)
		}
		b.setColor(Mask)
		b.fillRect(10, 0, 180, 10)
		b.fillRect(0, 0, 10, 380)
	}

	patterns["DIAGONAL_RIGHT"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 19; i++ {
			b.fillRect(0+190-i*10, 0+(i+1)*20, 10, 380-i*20)
		}
		b.setColor(getMaskColor(col))
		for i := 0; i < 20; i++ {
			b.fillRect(0+190-i*10, 0+i*20, 10, 20)
		}
		b.setColor(Mask)
		b.fillRect(10, 390, 180, 10)
		b.fillRect(190, 20, 10, 380)
	}

	// invert 0
	patterns["DIAGONAL_LEFT_MIRROR"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 19; i++ {
			b.fillRect(0+i*10, 0+(i+1)*20, 10, 380-i*20)
		}
		b.setColor(getMaskColor(col))
		for i := 0; i < 20; i++ {
			b.fillRect(0+i*10, 0+i*20, 10, 20)
		}
		b.setColor(Mask)
		b.fillRect(10, 390, 180, 10)
		b.fillRect(0, 20, 10, 380)
	}

	patterns["DIAGONAL_RIGHT_MIRROR"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 19; i++ {
			b.fillRect(0+190-i*10, 0, 10+i*10, 380-i*20)
		}
		b.setColor(getMaskColor(col))
		for i := 0; i < 20; i++ {
			b.fillRect(0+190-i*10, 0+380-i*20, 10, 20)
		}
		b.setColor(Mask)
		b.fillRect(10, 0, 180, 10)
		b.fillRect(190, 0, 10, 380)
	}

	patterns["CIRCLE_MIDDLE"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))
		b.fillRect(90, 140, 20, 120)
		b.fillRect(70, 150, 60, 100)
		b.fillRect(60, 160, 80, 80)
		b.fillRect(50, 170, 100, 60)
		b.fillRect(40, 190, 120, 20)
		b.setColor(col)
		b.fillRect(90, 150, 20, 100)
		b.fillRect(70, 160, 60, 80)
		b.fillRect(60, 170, 80, 60)
		b.fillRect(50, 190, 100, 20)
	}

	patterns["RHOMBUS_MIDDLE"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 7; i++ {
			b.setColor(getMaskColor(col))
			if i < 6 {
				b.fillRect(30+10*i, 190-20*i, 140-20*i, 20+40*i)
				b.setColor(col)
				b.fillRect(40+10*i, 190-20*i, 120-20*i, 20+40*i)
			} else {
				b.fillRect(90, 80, 20, 240)
			}
		}
	}

	patterns["HALF_VERTICAL"] = func(b *Banner, col color.RGBA) {
		b.fillRect(0, 0, 100, 400)
		b.setColor(getMaskColor(col))
		b.fillRect(100, 0, 10, 400)
		b.setColor(Mask)
		b.fillRect(10, 0, 90, 10)
		b.fillRect(10, 390, 90, 10)
		b.fillRect(0, 0, 10, 400)
	}

	// invert 0
	patterns["HALF_VERTICAL_MIRROR"] = func(b *Banner, col color.RGBA) {
		b.fillRect(100, 0, 100, 400)
		b.setColor(getMaskColor(col))
		b.fillRect(90, 0, 10, 400)
		b.setColor(Mask)
		b.fillRect(100, 0, 90, 10)
		b.fillRect(100, 390, 90, 10)
		b.fillRect(190, 0, 10, 400)
	}

	patterns["HALF_HORIZONTAL"] = func(b *Banner, col color.RGBA) {
		b.fillRect(0, 0, 200, 200)
		b.setColor(getMaskColor(col))
		b.fillRect(0, 200, 200, 10)
		b.setColor(Mask)
		b.fillRect(0, 0, 10, 200)
		b.fillRect(190, 0, 10, 200)
		b.fillRect(10, 0, 180, 10)
	}

	patterns["HALF_HORIZONTAL_MIRROR"] = func(b *Banner, col color.RGBA) {
		b.fillRect(0, 200, 200, 200)
		b.setColor(getMaskColor(col))
		b.fillRect(0, 190, 200, 10)
		b.setColor(Mask)
		b.fillRect(0, 200, 10, 200)
		b.fillRect(190, 200, 10, 200)
		b.fillRect(10, 390, 180, 10)
	}

	patterns["CREEPER"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))
		b.fillRect(10, 140, 60, 60)
		b.fillRect(130, 140, 60, 60)
		b.fillRect(40, 230, 30, 90)
		b.fillRect(70, 200, 60, 90)
		b.fillRect(130, 230, 30, 90)
		b.setColor(col)
		b.fillRect(20, 150, 40, 40)
		b.fillRect(140, 150, 40, 40)
		b.fillRect(50, 240, 10, 70)
		b.fillRect(80, 210, 40, 70)
		b.fillRect(140, 240, 10, 70)
		b.fillRect(50, 240, 100, 40)
	}

	patterns["BRICKS"] = func(b *Banner, col color.RGBA) {
		offset := 0
		for i := 0; i < 10; i++ {
			offset = i % 2 * 30
			for j := 0; j < 4; j++ {
				b.setColor(getMaskColor(col))
				b.fillRect(60*j-offset, i*40, 50, 30)
				b.setColor(col)
				b.fillRect(60*j-offset, i*40, 40, 20)
				b.setColor(getMaskColor(col))
				b.fillRect(40+60*j-offset, i*40,
					10, 10)
			}
		}
		b.setColor(Background)
		b.fillRect(0-30, 0, 30, 400)
		b.fillRect(200, 0, 30, 400)
	}

	patterns["GRADIENT"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 40; i++ {
			col.A = uint8((40 - i) / 40.0 * 255)
			b.setColor(col)
			b.fillRect(0, i*10, 200, 10)
		}
		b.setColor(getMaskColor(col))
		b.fillRect(0, 0, 200, 10)
	}

	patterns["GRADIENT_UP"] = func(b *Banner, col color.RGBA) {
		for i := 0; i < 40; i++ {
			col.A = uint8((40 - i) / 40.0 * 255)
			b.setColor(col)
			b.fillRect(0, 390-i*10, 200, 10)
		}
		b.setColor(getMaskColor(col))
		b.fillRect(0, 390, 200, 10)
	}

	patterns["SKULL"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))

		b.fillRect(30, 210, 20, 20)
		b.fillRect(50, 220, 20, 20)
		b.fillRect(70, 230, 20, 20)
		b.fillRect(30, 260, 30, 10)

		b.fillRect(90, 240, 20, 10)
		b.fillRect(60, 250, 80, 10)

		b.fillRect(150, 210, 20, 20)
		b.fillRect(130, 220, 20, 20)
		b.fillRect(110, 230, 20, 20)
		b.fillRect(140, 260, 30, 10)

		b.setColor(col)

		b.fillRect(60, 120, 80, 40)
		b.fillRect(60, 190, 80, 10)
		b.fillRect(60, 120, 10, 80)
		b.fillRect(130, 120, 10, 80)
		b.fillRect(70, 170, 20, 10)
		b.fillRect(90, 160, 20, 10)
		b.fillRect(110, 170, 20, 10)

		b.fillRect(20, 200, 20, 20)
		b.fillRect(20, 200, 10, 30)
		b.fillRect(40, 220, 20, 10)
		b.fillRect(60, 230, 20, 10)
		b.fillRect(80, 240, 10, 20)

		b.fillRect(90, 250, 20, 10)

		b.fillRect(160, 200, 20, 20)
		b.fillRect(170, 200, 10, 30)
		b.fillRect(140, 220, 20, 10)
		b.fillRect(120, 230, 20, 10)
		b.fillRect(110, 240, 10, 20)

		b.fillRect(20, 260, 10, 30)
		b.fillRect(30, 270, 10, 20)
		b.fillRect(40, 270, 20, 10)
		b.fillRect(60, 260, 20, 10)

		b.fillRect(170, 260, 10, 30)
		b.fillRect(160, 270, 10, 20)
		b.fillRect(140, 270, 20, 10)
		b.fillRect(120, 260, 20, 10)

		col.A = 90
		b.setColor(col)

		b.fillRect(70, 160, 20, 10)
		b.fillRect(90, 170, 20, 10)
		b.fillRect(110, 160, 20, 10)
		b.fillRect(70, 180, 60, 10)
	}

	patterns["FLOWER"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))

		b.fillRect(80, 170, 40, 60)
		b.fillRect(70, 180, 60, 40)

		b.fillRect(80, 120, 40, 20)
		b.fillRect(40, 140, 120, 20)
		b.fillRect(90, 110, 20, 10)

		b.fillRect(30, 130, 40, 10)
		b.fillRect(40, 120, 20, 10)
		b.fillRect(60, 160, 10, 10)

		b.fillRect(130, 130, 40, 10)
		b.fillRect(140, 120, 20, 10)
		b.fillRect(130, 160, 10, 10)

		b.fillRect(80, 260, 40, 20)
		b.fillRect(40, 240, 120, 20)
		b.fillRect(90, 280, 20, 10)

		b.fillRect(30, 260, 40, 10)
		b.fillRect(40, 270, 20, 10)
		b.fillRect(60, 230, 10, 10)

		b.fillRect(130, 260, 40, 10)
		b.fillRect(140, 270, 20, 10)
		b.fillRect(130, 230, 10, 10)

		b.fillRect(10, 160, 10, 20)
		b.fillRect(20, 150, 10, 10)
		b.fillRect(20, 160, 20, 30)

		b.fillRect(50, 170, 10, 60)

		b.fillRect(10, 220, 10, 20)
		b.fillRect(20, 240, 10, 10)
		b.fillRect(20, 210, 20, 30)

		b.fillRect(180, 160, 10, 20)
		b.fillRect(170, 150, 10, 10)
		b.fillRect(160, 160, 20, 30)

		b.fillRect(140, 170, 10, 60)

		b.fillRect(180, 220, 10, 20)
		b.fillRect(170, 240, 10, 10)
		b.fillRect(160, 210, 20, 30)

		b.setColor(col)

		b.fillRect(80, 180, 40, 40)

		b.fillRect(90, 120, 20, 20)
		b.fillRect(50, 140, 100, 10)

		b.fillRect(40, 130, 20, 10)
		b.fillRect(50, 150, 20, 10)
		b.fillRect(40, 150, 20, 20)

		b.fillRect(140, 130, 20, 10)
		b.fillRect(130, 150, 20, 10)
		b.fillRect(140, 150, 20, 20)

		b.fillRect(90, 260, 20, 20)
		b.fillRect(50, 250, 100, 10)

		b.fillRect(40, 260, 20, 10)
		b.fillRect(50, 240, 20, 10)
		b.fillRect(40, 230, 20, 20)

		b.fillRect(140, 260, 20, 10)
		b.fillRect(130, 240, 20, 10)
		b.fillRect(140, 230, 20, 20)

		b.fillRect(20, 160, 10, 20)
		b.fillRect(30, 170, 10, 10)

		b.fillRect(40, 170, 10, 60)

		b.fillRect(20, 220, 10, 20)
		b.fillRect(30, 220, 10, 10)

		b.fillRect(170, 160, 10, 20)
		b.fillRect(160, 170, 10, 10)

		b.fillRect(150, 170, 10, 60)

		b.fillRect(170, 220, 10, 20)
		b.fillRect(160, 220, 10, 10)
	}

	patterns["MOJANG"] = func(b *Banner, col color.RGBA) {
		b.setColor(getMaskColor(col))
		b.fillRect(20, 160, 40, 110)
		b.fillRect(30, 150, 10, 10)
		b.fillRect(50, 140, 40, 10)
		b.fillRect(90, 150, 10, 10)
		b.fillRect(110, 160, 10, 10)
		b.fillRect(120, 150, 10, 10)
		b.fillRect(130, 160, 10, 10)
		b.fillRect(140, 120, 10, 10)
		b.fillRect(140, 150, 20, 10)
		b.fillRect(70, 170, 40, 10)
		b.fillRect(30, 270, 10, 10)
		b.fillRect(50, 280, 110, 10)
		b.fillRect(160, 260, 20, 20)
		b.fillRect(60, 240, 10, 10)
		b.fillRect(70, 250, 10, 10)
		b.fillRect(80, 260, 20, 10)
		b.fillRect(130, 180, 10, 10)
		b.fillRect(140, 190, 10, 10)
		b.fillRect(150, 200, 10, 10)
		b.fillRect(160, 220, 10, 10)
		b.fillRect(170, 200, 10, 20)
		b.fillRect(160, 170, 10, 30)

		b.setColor(col)
		b.fillRect(30, 160, 20, 110)
		b.fillRect(40, 150, 20, 50)
		b.fillRect(40, 150, 30, 40)
		b.fillRect(40, 150, 50, 10)
		b.fillRect(40, 160, 70, 10)
		b.fillRect(110, 170, 50, 10)
		b.fillRect(120, 160, 10, 10)
		b.fillRect(140, 180, 20, 10)
		b.fillRect(150, 190, 10, 10)
		b.fillRect(160, 200, 10, 20)
		b.fillRect(140, 130, 20, 20)
		b.fillRect(40, 230, 20, 40)
		b.fillRect(60, 250, 10, 20)
		b.fillRect(40, 260, 40, 20)
		b.fillRect(80, 270, 90, 10)
	}
}
