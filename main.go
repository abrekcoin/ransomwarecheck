package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"ransomwarecheck/watch"
	"ransomwarecheck/scout"
	"ransomwarecheck/models"
)

func main() {

	//pwd
	path, err := os.Getwd()
	models.Pathes = path
	if err != nil {
		log.Println(err)
	}
	//list directory
	filepath.WalkDir(path, scout.Visit)
	fmt.Println("total", len(models.My_slice), " files created")
	fmt.Println(models.My_slice)
	a := 0
	for a < len(models.My_slice)-1 {
		go watch.Watch(models.My_slice[a])
		a++
	}
	watch.Watch(models.My_slice[len(models.My_slice)-1])
}




