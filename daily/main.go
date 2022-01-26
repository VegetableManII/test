package main

import "study/daily/questions"

func main() {
	questions.GetNetInterfaces()
}

/* func main() {
	tmap := map[int]string{
		1: "1",
		2: "2",
		3: "3",
	}
	go func() {
		for i := 0; i < math.MaxInt32; i++ {
			tmap[i] = strconv.Itoa(i)
		}
	}()

	go func() {
		for i := 0; i < math.MaxInt32; i++ {
			log.Println(tmap[i])
		}
	}()
	time.Sleep(5 * time.Second)
} */
