package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Print(<- c)

	// for i := 0; i < len(links) ; i++ {
	// 	fmt.Println(<- c)
	// }

	// for { // infinite loop till blocking function comes into flow
	// 	// fmt.Println(<- c)
	// 	go checkLink(<-c, c)
	// }

	for l:= range c { // alternative syntax for above one
		// go checkLink(l, c)
		go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		} (l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		// fmt.Println(link, "might be down!")
		// c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yes Its Up"
}
