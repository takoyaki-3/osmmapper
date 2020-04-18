# OSM Mapper

Convert Open Street Map .pdb to .osm and slect node and edge as csv.
After that, draw as an image.

## Dependency

- OS : Windows pro or Ubuntu
- Docker
- Docker Compose
- g++

## How to use

### 1. Edit map URL

First, edit the map URL write in `start.sh`.
In this repo, it contains the URL of the Kanto region of Japan.

### 2. Build image and start container

On Windows, osmconverter64 was garbled in my computer. So, run converter in Centos docker container.

```
docker-compouse up -d
```

### 3. Select node, edge, tag.

If go is not installed, please [install go](https://golang.org/doc/install).

```
cd osm
go run node.go
go run edge.go
go run tag.go
```

### 4. combind node,edge,tag

```
g++ combind.cpp -std=c++11 -O3
./a.out
```

### 5. Draw as an image.

```
go run draw.go
```

![](https://gyazo.com/e182ed20dfe31140b3caa8a0d310851a.png)