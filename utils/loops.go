package utils

func ForLoopDemo() {
	loopThatTerminateBasedOnAConditionBreak()
	println("---------------------------------------")
	loopThatTerminateBasedOnAConditionContinue()
	println("---------------------------------------")
	loopWithPostClauses()
	println("---------------------------------------")
	infiniteLoop()
}

func loopThatTerminateBasedOnAConditionBreak() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}
}

func loopThatTerminateBasedOnAConditionContinue() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
		println("continuing...")
	}
}

func loopWithPostClauses() {
	// var i int
	// for ; i < 5; i++ {
	// 	println(i)
	// }

	for i := 0; i < 5; i++ {
		println(i)
	}
}

func infiniteLoop() {
	// // Ugly syntax
	// var i int
	// for ; ; {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	println(i)
	// 	i++
	// }

	var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}
