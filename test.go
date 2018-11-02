package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
)

func main() {
    envPath := fmt.Sprintf(
        "samples\\php\\laravel5-example%s.env",
        string(os.PathSeparator),
    )
    envBytes, err := ioutil.ReadFile(envPath)
    if err != nil {
        fmt.Print(err)
    }
    envStr := string(envBytes)
    var rex = regexp.MustCompile("([a-zA-Z_][a-zA-Z0-9_]*)=(.*)")
    envMatch := rex.FindAllStringSubmatch(envStr, -1)
    //envMap := make(map[string]string)
	for _, kv := range envMatch {
        fmt.Println(kv[1])
        fmt.Println(kv[2])
	}

}
