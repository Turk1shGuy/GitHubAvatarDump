package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const base_url string = "https://avatars.githubusercontent.com/u/"

func main() {
	for i := 1; i < 9999999999; i++ {
		go user_photo_dump(i)
		time.Sleep(time.Millisecond * 10)
	}

	println("Data Dump Complite")
}

func user_photo_dump(id int) {
	url := fmt.Sprintf("%s%d", base_url, id)

	println("Dumping: ", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to retrieve image. Status code: %d\n", resp.StatusCode)
		return
	}

	file, err := os.Create("./dump/" + strconv.Itoa(id) + ".png")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
