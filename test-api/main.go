package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	//canal <-15
	//valor := <-14
	start := time.Now()
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://ap.github.com",
		"https://outlook.office.com",
		"https://api/somewhereintheinternet.com",
		"https://graph.microsoft.com",
	}
	canal := make(chan string)
	for _, api := range apis {
		go checkAPI(api, canal)
	}
	for i := 0; i < len(apis); i++ {
		fmt.Println(<-canal)
	}

	//	time.Sleep(5 * time.Second)
	alapsed := time.Since(start)
	fmt.Printf("¡Listo %v, Tomo segundos!\n", alapsed.Seconds())
}

func checkAPI(api string, canal chan string) {
	if _, err := http.Get(api); err != nil {
		canal <- fmt.Sprintf("ERROR:¡%s esta caido!\n", api)
		return
	}
	canal <- fmt.Sprintf("SUCCESS:¡%s esta Ok!\n", api)

}
