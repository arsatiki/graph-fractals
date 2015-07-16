package main

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Id int64

type OSM struct {
	Bounds OSMBounds `xml:"bounds"`

	Nodes     []OSMNode     `xml:"node"`
	Ways      []OSMWay      `xml:"way"`
	Relations []OSMRelation `xml:"relation"`
}

type OSMBounds struct {
	MinLat float64 `xml:"minlat,attr"`
	MinLon float64 `xml:"minlon,attr"`
	MaxLat float64 `xml:"maxlat,attr"`
	MaxLon float64 `xml:"maxlon,attr"`
}

type OSMTag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

type OSMNode struct {
	Id   Id       `xml:"id,attr"`
	Tags []OSMTag `xml:"tag"`

	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
}

type OSMWay struct {
	Id   Id       `xml:"id,attr"`
	Tags []OSMTag `xml:"tag"`

	Refs []OSMRef `xml:"nd"`
}

type OSMRef struct {
	Ref Id `xml:"ref,attr"`
}

type OSMRelation struct {
	Id   Id       `xml:"id,attr"`
	Tags []OSMTag `xml:"tag"`
}

func main() {
	var doc OSM

	r, err := gzip.NewReader(os.Stdin)
	if err != nil {
		log.Fatal("Could not open stdin??", err)
	}

	decoder := xml.NewDecoder(r)
	err = decoder.Decode(&doc)
	if err != nil {
		log.Fatal("Error parsing XML", err)
	}

	b := doc.Bounds
	fmt.Printf("Found doc with bounds %f, %f to %f, %f\n",
		b.MinLat, b.MinLon, b.MaxLat, b.MaxLon)
	fmt.Println("Diameter is",
		GCDist(b.MinLat, b.MinLon, b.MaxLat, b.MaxLon))
	fmt.Printf("%d nodes, %d ways and %d relations\n",
		len(doc.Nodes), len(doc.Ways), len(doc.Relations))
	fmt.Printf("Node 2 looks like this: %v\n", doc.Nodes[2])
	fmt.Printf("Way 0 looks like this: %v\n", doc.Ways[0])
	fmt.Printf("Relation 0 looks like this: %v\n", doc.Relations[0])
}
