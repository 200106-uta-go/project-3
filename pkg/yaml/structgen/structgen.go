package structgen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

//Generated holds the data structure Generated from a structured file
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

//GetKey returns a slice of type Generated containing all matches
//for key in a Generated struct
func (g Generated) GetKey(key string) []Generated {
	var matches []Generated

	//range over Generated struct
	for i, v := range g {
		if g.IsMap(v) {
			//create a new Generated and a temporary map to iterate over
			newGen := Generated{}
			tempMap := reflect.ValueOf(v)

			//iterates through each value in tempMap and adds the key value to newGen
			iterate := tempMap.MapRange()
			for iterate.Next() {
				newGen[iterate.Key().String()] = iterate.Value().Interface()
			}

			//if the key you're searching for is a map, add it
			if i == key {
				matches = append(matches, newGen)
			}

			//recursively append matches into matches slice
			matches = append(matches, newGen.GetKey(key)...)
		} else if g.IsGenSlice(v) {
			//iterate over slice to check values of internal Generated structs for key
			tempSlice := reflect.ValueOf(v)

			// iterates through each Generated in tempSlice and adds the key value to newGen
			for i := 0; i < tempSlice.Len(); i++ {
				// create a new Generated and a temporary map to iterate over
				newGen := Generated{}
				tempMap := tempSlice.Index(i)

				//iterates through each value in tempMap and adds the key value to newGen
				iterate := tempMap.Elem().MapRange()
				for iterate.Next() {
					// fmt.Println(iterate.Key().String())
					newGen[iterate.Key().String()] = iterate.Value().Interface()
				}

				//recursively append matches into matches slice
				matches = append(matches, newGen.GetKey(key)...)
			}
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

//FilterValues returns a string slice of all --- it's not working yet
// func (g Generated) FilterValues(key string) []string {
// 	res := g.GetKey(key)
// 	var filtered []string
// 	for i := range res {
// 		// fmt.Println(res[i])
// 		filtered = append(filtered, fmt.Sprint(res[i][key]))
// 	}
// 	return filtered
// }

//IsMap checks if the value if of type structgen.Generated
func (g Generated) IsMap(value interface{}) bool {
	if value != nil {
		return "structgen.Generated" == reflect.TypeOf(value).String()
	}
	return false
}

//IsGenSlice checks if the value if of type []interface {}
func (g Generated) IsGenSlice(value interface{}) bool {
	if value != nil {
		//make sure type of value is []interface{}
		if "[]interface {}" == reflect.TypeOf(value).String() {
			if reflect.ValueOf(value).Len() > 0 {
				//ugly way to see if the values inside the slice are maps
				sliceType := fmt.Sprint(reflect.ValueOf(value).Index(0))
				if strings.HasPrefix(sliceType, "map") {
					return true
				}
			}
		}
	}
	return false
}

//Print prints the contents of a Generated struct
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
