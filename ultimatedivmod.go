package piscine

func UltimateDivMod(a *int, b *int) {
	resDiv := *a / *b
	resMod := *a % *b
	*a = resDiv
	*b = resMod
}
