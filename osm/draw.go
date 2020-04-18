package main

import (
    "os"
//  "math"
    "image"
    "image/jpeg"
    "image/color"
    "strconv"
    "encoding/csv"
//  "fmt"
)

const scale = 100000
const pie = 3.14159265359

const area_minlat,area_minlon,area_maxlat,area_maxlon=35.55,139.6,35.8,139.9

var minlat,minlon,maxlat,maxlon int
var datamap,datamap2 [][]int
var coun,num_thread int

func main() {

    minlat,minlon,maxlat,maxlon = area_minlat*scale,area_minlon*scale,area_maxlat*scale,area_maxlon*scale

    // メモリ確保
    datamap = make([][]int,int(maxlat-minlat))
    for i := range datamap{
        datamap[i]=make([]int,int(maxlon-minlon))
    }
    datamap2 = make([][]int,int(maxlat-minlat))
    for i := range datamap2{
        datamap2[i]=make([]int,int(maxlon-minlon))
    }

    loaddatafile()

    // 画像処理用
    x := 0
    y := 0
    width := maxlon-minlon
    height := maxlat-minlat

    // RectからRGBAを作る(ゼロ値なので黒なはず)
    img := image.NewRGBA(image.Rect(x, y, width, height))

    imagefile, _ := os.Create("map.jpg")
    defer imagefile.Close()

    for i:=0;i<height;i++{
        for j:=0;j<width;j++{
            lonx:=j
            laty:=height-i
            col := datamap2[i][j]
//          col := math.Log10(float64(datamap2[i][j]+1))*100

            if col > 0{
                col = 255
            }
            img.Set(lonx, laty, color.RGBA{uint8(col), uint8(col), uint8(col), 1})
        }
    }

    // JPEGで出力(100%品質)
    if err := jpeg.Encode(imagefile, img, &jpeg.Options{100}); err != nil {
        panic(err)
    }
}

func numthreadfunc(){
    num_thread--
}

func loaddatafile(){
    file, err := os.Open("combind.csv")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    var line []string

    for {
        line, err = reader.Read()
        if err != nil {
            break
        }
        var aflat,aflon,bflat,bflon float64
        aflat, _ = strconv.ParseFloat(line[2],64)
        aflon, _ = strconv.ParseFloat(line[3],64)
        bflat, _ = strconv.ParseFloat(line[5],64)
        bflon, _ = strconv.ParseFloat(line[6],64)

        if aflat < area_minlat || area_maxlat < aflat || aflon < area_minlon || area_maxlon < aflon{
            continue
        }
        if bflat < area_minlat || area_maxlat < bflat || bflon < area_minlon || area_maxlon < bflon{
            continue
        }

        aintlat := int(aflat*scale)
        aintlon := int(aflon*scale)
        bintlat := int(bflat*scale)
        bintlon := int(bflon*scale)

        if aintlat >= maxlat || aintlon >= maxlon{
            continue
        }
        if aintlat < minlat || aintlon < minlon{
            continue
        }
        if bintlat >= maxlat || bintlon >= maxlon{
            continue
        }
        if bintlat < minlat || bintlon < minlon{
            continue
        }

        dlat,dlon := aintlat-bintlat,aintlon-bintlon;
        if dlat < 0 {
            dlat = -dlat
        }
        if dlon < 0 {
            dlon = -dlon
        }
        ddd:=dlat
        if dlon > dlat{
            ddd = dlon
        }
        if ddd==0{
            continue
        }
        for ilat:=0;ilat<=ddd;ilat++{
            ilon:=ilat
            tlat := int(ilat*aintlat/ddd+(ddd-ilat)*bintlat/ddd)
            tlon := int(ilon*aintlon/ddd+(ddd-ilon)*bintlon/ddd)

/*          fmt.Println("aa")
            fmt.Println(aintlat)
            fmt.Println(bintlat)
            fmt.Println("bb")*/

            if tlat-minlat < 0 || tlon-minlon < 0 {
                continue
            }
            if maxlat-minlat <= tlat-minlat || maxlat-minlat <= tlon-minlon{
                continue
            }

            datamap[tlat-minlat][tlon-minlon]++
            datamap2[tlat-minlat][tlon-minlon]++
        }

//      datamap[aintlat-minlat][aintlon-minlon]++
//      datamap2[aintlat-minlat][aintlon-minlon]++
    }
    num_thread--
}