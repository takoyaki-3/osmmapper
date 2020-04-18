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
XMLName     xml.Name  `xml:"osm"`
Xml         string    `xml:",innerxml"`
	Ways	[]Way `xml:"way"`
}

type Way struct {
  Id string `xml:"id,attr"`
  Nds []nd `xml:"nd"`
}

type nd struct {
    Ref string `xml:"ref,attr"`
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
//fmt.Println(xmlData);
    var data OSM
    if err := xml.Unmarshal(xmlData, &data); err != nil {
        fmt.Println("XML Unmarshal error:", err)
        return
    }

    file2,err2 := os.OpenFile("edge.csv", os.O_WRONLY|os.O_CREATE, 0600)
//    failOnError(err2)
    defer file2.Close()

//    err2 = file2.Truncate(0) // ファイルを空っぽにする(実行2回目以降用)
//    failOnError(err2)
    if err2 != nil {
        log.Fatal(err)
        return
    }
    writer := csv.NewWriter(file2)
    for itm := 0;itm<len(data.Ways);itm++ {
        for itmnd := 0;itmnd<len(data.Ways[itm].Nds)-1;itmnd++{
            writer.Write([]string{data.Ways[itm].Id, data.Ways[itm].Nds[itmnd].Ref, data.Ways[itm].Nds[itmnd+1].Ref})
        }
    }
    writer.Flush()

//fmt.Println(data.XMLName)
//    fmt.Println(data.Name)
//    fmt.Println(data.Nodes)
//    fmt.Println(data.Nodes[0].Lat)
 //   fmt.Println(data.osmdata.Nodes[1].lon)
}
