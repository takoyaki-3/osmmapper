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
    Nodes   []Node `xml:"node"`
}

type Node struct {
    Id  string `xml:"id,attr"`
    Lat string `xml:"lat,attr"`
    Lon string `xml:"lon,attr"`
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

    file2,err2 := os.OpenFile("node.csv", os.O_WRONLY|os.O_CREATE, 0600)
    defer file2.Close()

    if err2 != nil {
        log.Fatal(err)
        return
    }
    writer := csv.NewWriter(file2)
    for itm := 0;itm<len(data.Nodes);itm++ {
        writer.Write([]string{data.Nodes[itm].Id, data.Nodes[itm].Lat,data.Nodes[itm].Lon})
    }
    writer.Flush()
}
