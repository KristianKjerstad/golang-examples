package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func cliexample() {
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()
	file, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() 


	bufReader := bufio.NewReader(file)
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		} 
	}	
}