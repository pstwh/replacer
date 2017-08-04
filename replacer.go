package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"io/ioutil"
)


func main() {	
	var map_text string

	if(len(os.Args) < 3) {
		fmt.Println("Você deve selecionar o mapa e o diretório!")
		return
	}

	map_file, err := os.Open(string(os.Args[1]))
	if err != nil {
		return
	}
	defer map_file.Close()

	stat, err := map_file.Stat()
	if err != nil {
  	return
  }

   map_bytes := make([]byte, stat.Size())

	_, err = map_file.Read(map_bytes)
	if err != nil {
		return
  	}

  	map_text = string(map_bytes)

  	const map_instance_regex = `[0-9a-z_]*=>[0-9a-z_]*`

  	map_regex, err := regexp.Compile(map_instance_regex)
	if err != nil {
		panic(err)
		fmt.Print(err)
	}

	maps := map_regex.FindAllStringSubmatch(map_text, -1)

	dir := string(os.Args[2])

	for _, m := range maps {
		temp := strings.Split(m[0], "=>")
		fmt.Println(temp)
	
		filepath.Walk(string(dir), func(path string, fi os.FileInfo, err error) error {
			
			if err != nil {
				return err
			}

			if !!fi.IsDir() {
				return nil //
			}

			arq, err := filepath.Match("*", fi.Name())

			if err != nil {
				panic(err)
				return err
			}

			if arq {
				read, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}

				fmt.Println(path)

				write := strings.Replace(string(read), temp[0], temp[1], -1)

				err = ioutil.WriteFile(path, []byte(write), 0)
				if err != nil {
					panic(err)
				}

			}

			return nil		
		})
	}
}

