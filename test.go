package main

import (
	"fmt"
	"log"
)

func main() {
	/*x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	fmt.Println(x + y)*/
	/*
		reader := bufio.NewReader(os.Stdin)
		var i int
		_, err := fmt.Scanf("%d", &i)
		var a int
		_, err := fmt.Scanf("%d", &a)
		fmt.println*/
	var width, height int
	if _, err := fmt.Scan(&width, &height); err != nil {
		log.Print("  Scan for i, j & k failed, due to ", err)
		return
	}
	if height == 0 || width == 0 {
		fmt.Println(0, 0)
	}
	whole := 0
	half := 0
	wmod := false
	hmod := false
	if height%2 == 0 {
		hmod = true
	}
	if width%2 == 0 {
		wmod = true
	}
	if wmod {
		if hmod {
			whole = (width/2)*height/2 + (width/2-1)*height/2
			half = height
		} else {
			whole = (width/2)*(height/2+1) + (width/2-1)*(height/2+1)
			half = height - 1
		}
	} else {
		if hmod {
			whole = (width-1)*(height/2) + (width-1)*(height/2)
			half = height + height/2
		} else {
			whole = (width-1)/2*(height+1)/2 + ((width-1)/2)*((height-1)/2)
			half = (height-1)/2 + height
		}

	}
	fmt.Println(whole, half)
}
