package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

func SlugifyFiles(root string) {
	fmt.Println("slugify files.", root)
	stack := make([] string,0)
	rootStat , err := os.Stat(root)
	if err != nil {
		log.Fatal(err)
		return
	}
	if rootStat.IsDir() {
		stack = append(stack, root)
		replaceRe := regexp.MustCompile("[^a-zA-Z0-9]")
		for len(stack) > 0 {
			currentLen := len(stack)
			current := stack[currentLen-1]
			stack = stack[:currentLen-1]
			contents, _ := ioutil.ReadDir(current)
			for _, v := range contents {
				oldname := path.Join(current, v.Name())
				slug := string(replaceRe.ReplaceAll([]byte(v.Name()), []byte("_")))
				newname := path.Join(current, slug)
				if !v.IsDir() {
					extension := path.Ext(v.Name())
					basename := strings.Replace(v.Name(), extension, "", 1)
					slug = string(replaceRe.ReplaceAll([]byte(basename), []byte("_")))
					newname = path.Join(current, slug + extension)
				} 
				
				if oldname != newname {
					newname = strings.ToLower(newname)
					fmt.Println(oldname, newname)
					os.Rename(oldname, newname)
				}
				if v.IsDir() {
					stack = append(stack, newname)
				}
			}
		}
	}
}