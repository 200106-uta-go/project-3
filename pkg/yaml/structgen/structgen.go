package structgen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

//Generated holds the data structure generated from a structured file
type Generated map[string]interface{}

// FromFile generates an anonymous struct holding all values present in the given file
// This function supports json and yaml files. Returns nil if struct could not be created
func FromFile(file *os.File) Generated {
	ext := getFileExtension(file)

	switch ext {
	case "json":
		a := fromJSON(file)
		return a
	case "yaml":
		b := fromYAML(file)
		return b
	default:
		invalidFile(file)
		return nil
	}
}

//Copy creates a new file by unmarshaling the old into a map and then marshaling the map into new ... pointless
//This can eventaully be modified to copy contents and injecct updates
func Copy(oldFile *os.File, newFile *os.File) {
	ext := getFileExtension(oldFile)

	switch ext {
	case "json":
		a := fromJSON(oldFile)
		toJSON(a, newFile)
	case "yaml":
		b := fromYAML(oldFile)
		toYAML(b, newFile)
	default:
		invalidFile(oldFile)
	}
}

//HasKey checks if the generated struct has a key value pair
//This also iterates through nested maps if there are any
func (g Generated) HasKey(key string) []Generated {
	var matches []Generated

	//range over generated struct
	for i, v := range g {
		if g.IsMap(v) {
			//create a new generated and a temporary map to iterate over
			newGen := Generated{}
			tempMap := reflect.ValueOf(v)

			//for each key value in tempMap, copy into newGen
			for innerIndex := 0; innerIndex < tempMap.Len(); innerIndex++ {
				newGen[strconv.Itoa(innerIndex)] = tempMap.Index(innerIndex)
			}

			//recursively append matches into matches slice
			matches = append(matches, newGen.HasKey(key)...)
		} else {
			//put match into new Generated to append to matches
			if i == key {
				matches = append(matches, Generated{
					i: v,
				})
			}
		}
	}
	return matches
}

//IsMap ...
func (g Generated) IsMap(value interface{}) bool {
	return "structgen.Generated" == reflect.TypeOf(value).String()
}

//Print prints the contents of a generated struct
func (g Generated) Print() {
	fmt.Println(g)
}

func getFileExtension(file *os.File) string {
	var ext string
	if strings.Contains(file.Name(), ".") {
		filename := strings.Split(file.Name(), ".")
		ext = filename[len(filename)-1]
	} else {
		fmt.Println("File has no extension")
		ext = file.Name()
	}
	return ext
}

func fromJSON(file *os.File) Generated {
	//read file's contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	//parse json into anonymous struct
	m := make(Generated)

	err = json.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

func fromYAML(file *os.File) Generated {
	//read file's contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	//parse yaml into anonymous struct
	m := make(Generated)

	err = yaml.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return m
}

func toJSON(myJSON Generated, file *os.File) {
	bytes, err := json.Marshal(myJSON)
	if err != nil {
		log.Fatalln(err)
	}

	file.Write(bytes)
}

func toYAML(myYAML Generated, file *os.File) {
	bytes, err := yaml.Marshal(myYAML)
	if err != nil {
		log.Fatalln(err)
	}

	file.Write(bytes)
}

func createFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	return file
}

func invalidFile(file *os.File) {
	fmt.Println("The file given is not supported by structgen")
}
