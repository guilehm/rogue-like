package helpers

func HasCollision(
	startXp1, startYp1, startXp2, startYp2, width, height, widthP2, heightP2, offset int,
) (bool, bool) {
	collisionX := false
	collisionY := false

	startXp1 -= offset
	startYp1 -= offset
	startXp2 -= offset
	startYp2 -= offset

	endXp1 := startXp1 + width + offset
	endYp1 := startYp1 + height + offset
	endXp2 := startXp2 + widthP2 + offset
	endYp2 := startYp2 + heightP2 + offset

	if startXp2 < endXp1 && endXp2 > startXp1 {
		collisionY = true
	}
	if startYp2 < endYp1 && endYp2 > startYp1 {
		collisionX = true
	}
	return collisionX, collisionY
}
