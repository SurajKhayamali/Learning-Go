package utils

func ForLoopDemo() {
	loopThatTerminateBasedOnAConditionBreak()
	println("---------------------------------------")
	loopThatTerminateBasedOnAConditionContinue()
	println("---------------------------------------")
	loopWithPostClauses()
	println("---------------------------------------")
	infiniteLoop()
	println("---------------------------------------")
	loopWithSlice()
	println("---------------------------------------")
	loopWithRange()
	println("---------------------------------------")
	loopWithMap()
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

func loopWithSlice() {
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}
}

func loopWithRange() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		println(i, v)
	}
}

func loopWithMap() {
	wellKnownPorts := map[string]int{"http": 80, "https": 433}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

	// Special form of multiple returns where the compiler has special consideration for looping over collection
	// So, We can completely ignore that second return value,
	// normally you cannot do that when you get two return values comming back from a function
	for k := range wellKnownPorts {
		println(k)
	}

	// using write only variable '_'
	for _, v := range wellKnownPorts {
		println(v)
	}
}
