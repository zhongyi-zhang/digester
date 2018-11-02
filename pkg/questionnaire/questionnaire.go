package main

import (
    "github.com/kun-lun/digester/pkg/detector"
    "github.com/kun-lun/digester/pkg/vmgroupcalc"
    "bufio"
    "fmt"
    "os"
    "reflect"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Project path?")
    scanner.Scan()
    path := scanner.Text()
    d, err := detector.New(path)
    if err != nil {
        panic(err)
    }

    possiblePackageManagers := d.DetectPackageManager()
    fmt.Printf("What is the package manager for the project?")
    for i, pm := range possiblePackageManagers {
        if i == 0 {
            fmt.Printf(" %s \\", strings.ToUpper(string(pm)))
        } else {
            fmt.Printf(" %s \\", string(pm))
        }
    }
    if len(possiblePackageManagers) > 0 {
        fmt.Printf(" other?\n")
    }

    scanner.Scan()
    inputPackageManager := scanner.Text()
    if inputPackageManager == "" {
        inputPackageManager = string(possiblePackageManagers[0])
    }
    d.ConfirmPackageManager(inputPackageManager)

    possibleFrameworks := d.DetectFramework()
    fmt.Printf("What is the framework of the project?")
    for i, fw := range possibleFrameworks {
        if i == 0 {
            fmt.Printf(" %s \\", strings.ToUpper(string(fw)))
        } else {
            fmt.Printf(" %s \\", string(fw))
        }
    }
    if len(possibleFrameworks) > 0 {
        fmt.Printf(" other?\n")
    }

    scanner.Scan()
    inputFramework := scanner.Text()
    if inputFramework == "" {
        inputFramework = string(possibleFrameworks[0])
    }
    d.ConfirmFramework(inputFramework)

    d.DetectConfig()

    bp := d.ExposeKnownInfo()

    // Ask for the empty fields
    nisp := &bp.NonIaaSPart
    if nisp.ProgrammingLanguage == "" {
        fmt.Println("What's the programming language?")
        scanner.Scan()
        nisp.ProgrammingLanguage = scanner.Text()
    }
    if nisp.Framework == "" {
        fmt.Println("What's the framework? NONE?" )
        scanner.Scan()
        nisp.Framework = scanner.Text()
    }
    if len(nisp.Databases) > 0 {
        needExtraInfo := false
        fmt.Println("Here is the database(s) we found:" )
        for i, db := range nisp.Databases {
            fmt.Printf("No.%d: {\n", i+1)
            s := reflect.ValueOf(&db).Elem()
            for j := 0; j < s.NumField(); j++ {
                valField := s.Field(j)
                typeField := s.Type().Field(j)
                tag := typeField.Tag
                val := valField.Interface().(string)
                fmt.Printf("  %s: %s\n", tag.Get("name"), val)
                if (val == "") {
                    needExtraInfo = true
                }
            }
            fmt.Println("}")
        }
        if (needExtraInfo) {
            fmt.Println("Please help fill the blank fields.")
            for i, _ := range nisp.Databases {
                db := &nisp.Databases[i]
                s := reflect.ValueOf(db).Elem()
                for j := 0; j < s.NumField(); j++ {
                    valField := s.Field(j)
                    typeField := s.Type().Field(j)
                    tag := typeField.Tag
                    val := valField.Interface().(string)
                    if (val == "") {
                        fmt.Printf("For the database No.%d: %s\n", i+1, tag.Get("question"))
                        scanner.Scan()
                        input := scanner.Text()
                        valField.SetString(input)
                    }
                }
            }
            fmt.Println("Done.")
        }
    }

    fmt.Println("What's your expected number of concurrent users?")
    scanner.Scan()
    concurrentUserNumberStr := scanner.Text()
    var concurrentUserNumber int
    concurrentUserNumber, err = strconv.Atoi(concurrentUserNumberStr)
    if err != nil {
        panic(err)
    }
    bp.IaaSPart.VMGroup = vmgroupcalc.Calc(vmgroupcalc.Requirment{
        ConcurrentUserNumber: concurrentUserNumber,
    })

    fmt.Printf("%#v\n", bp)
}

func isTrueOrFalse(s string) bool {
    us := strings.ToUpper(s)
    return us == "TRUE" || us == "FALSE"
}
