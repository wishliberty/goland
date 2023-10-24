package main

func Deferclosureloopv1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}
func Deferclosureloopv2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func Deferclousureloopv3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
