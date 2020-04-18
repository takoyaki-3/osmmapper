package main

import (
    "encoding/csv"
    "encoding/xml"
    "fmt"
    "os"
    "log"
    "io/ioutil"
)

type OSM struct {
    XMLName xml.Name  `xml:"osm"`
    Xml     string    `xml:",innerxml"`
    Ways    []Way     `xml:"way"`
}

type Way struct {
  Id    string  `xml:"id,attr"`
  Tags   []tag    `xml:"tag"`
}

type tag struct {
    K string `xml:"k,attr"`
    V string `xml:"v,attr"`
}

func main() {
    xmlFile, err := os.Open("kanto.osm")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer xmlFile.Close()
    xmlData, err := ioutil.ReadAll(xmlFile)
    if err != nil {
        log.Fatal(err)
        return
    }
    var data OSM
    if err := xml.Unmarshal(xmlData, &data); err != nil {
        fmt.Println("XML Unmarshal error:", err)
        return
    }

    file2,err2 := os.OpenFile("tag.csv", os.O_WRONLY|os.O_CREATE, 0600)
    defer file2.Close()

    if err2 != nil {
        log.Fatal(err)
        return
    }
    writer := csv.NewWriter(file2)
    for itm := 0;itm<len(data.Ways);itm++ {
        for itmnd := 0;itmnd<len(data.Ways[itm].Tags);itmnd++{
            writer.Write([]string{data.Ways[itm].Id, data.Ways[itm].Tags[itmnd].K, data.Ways[itm].Tags[itmnd].V})
        }
    }
    writer.Flush()
}