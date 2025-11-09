package gfx

func (pg *PepperGraphics) SetMask(id, x, y int) {
	if _, ok := pg.sprites[id]; !ok {
		return // Or handle error: mask sprite not found
	}
	pg.Surface.PushGroup()
	pg.isMasking = true
	pg.maskSpriteID = id
	pg.maskX = x
	pg.maskY = y
}

func (pg *PepperGraphics) ResetMask() {
	if !pg.isMasking {
		return
	}

	contentPattern := pg.Surface.PopGroup()
	defer contentPattern.Destroy()

	maskSprite := pg.sprites[pg.maskSpriteID]

	pg.Surface.PushGroup()
	pg.Surface.SetSourceSurface(maskSprite.Surface, float64(pg.maskX), float64(pg.maskY))
	pg.Surface.Paint()
	maskPattern := pg.Surface.PopGroup()
	defer maskPattern.Destroy()

	pg.Surface.SetSource(contentPattern)
	pg.Surface.Mask(*maskPattern)

	pg.isMasking = false
	pg.maskSpriteID = -1
}
