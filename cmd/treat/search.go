package main

import (
    "log"
    "fmt"
    "strings"
    "github.com/ubccr/treat"
)


func Search(dbpath string, fields *SearchFields) {
    s, err := NewStorage(dbpath)
    if err != nil {
        log.Fatalf("%s", err)
    }

    err = s.Search(fields, func (key *treat.AlignmentKey, a *treat.Alignment) {
        alt := fmt.Sprintf("%d", a.AltEditing)
        if a.AltEditing != 0 {
            alt = fmt.Sprintf("A%d", a.AltEditing)
        }

        fmt.Println(strings.Join([]string{
            key.Gene,
            key.Sample,
            fmt.Sprintf("%.4f", RoundPlus(a.Norm, 4)),
            fmt.Sprintf("%d", a.ReadCount),
            alt,
            fmt.Sprintf("%d", a.HasMutation),
            fmt.Sprintf("%d", a.EditStop),
            fmt.Sprintf("%d", a.JuncEnd),
            fmt.Sprintf("%d", a.JuncLen),
            a.GrnaEditString(),
            a.JuncSeq}, "\t"))
    })

    if err != nil {
        log.Fatalf("%s", err)
    }
}
