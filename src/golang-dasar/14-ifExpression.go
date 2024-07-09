package main

func ifExpression() {
	// if => if else => else

	// if
	a := 10
	b := 20
	if a > b {
		println("a lebih besar dari b")
	}

	// if else
	if a > b {
		println("a lebih besar dari b")
	} else {
		println("a lebih kecil dari b")
	}

	// else if
	if a > b {
		println("a lebih besar dari b")
	} else if a == b {
		println("a sama dengan b")
	} else {
		println("a lebih kecil dari b")
	}
}
