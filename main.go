package main

import (
	"fmt"
	"math"

	"github.com/batchatco/go-native-netcdf/netcdf"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	//-20.602992, -23.594518 //https://maps.app.goo.gl/TGBuqiGsmjKky2SE7
	//Depth example -347 m

	depth := -87.0
	lat := -20.602992
	long := -23.594518

	var idepth int
	depth = math.Abs(depth)
	if depth > 0 && depth <= 100 {
		//Convert to integer positive
		idepth = int(math.Abs(depth)) / 5
	} else {
		idepth = int(math.Abs(depth)) / 25
	}
	fmt.Printf("Depth array index: %d (%f)\n", idepth, depth)

	var ilat int
	if lat < 90.0 {
		ilat = int((90.0 - (math.Abs(math.Round(lat)))))
	} else {
		ilat = int(math.Round(lat))
	}
	fmt.Printf("Depth array lat: %d (%f)\n", ilat, lat)

	var ilong int
	if long < 180.0 {
		ilong = int(180.0-(math.Abs(math.Round(long))))
	} else {
		ilong = int(math.Round(long))
	}
	fmt.Printf("Depth array long: %d (%f)\n", ilong, long)

	// Open the file
	nc, err := netcdf.Open("woa23_decav_s01_01.nc")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// Read the NetCDF variable from the file
	vr, _ := nc.GetVariable("s_an")
	if vr == nil {
		panic("latitude variable not found")
	}

	fmt.Println("------ VARIABLES ---------")

	// List all variables
	vl := nc.ListVariables()
	for i, v := range vl {
		fmt.Println(i, v)
	}

	fmt.Println("------ SUBGROUPS ---------")

	gl := nc.ListSubgroups()
	for i, g := range gl {
		fmt.Println(i, g)
	}

	fmt.Println("------ VARGETTER ---------")

	vs, _ := nc.GetVarGetter("s_an")
	if vs == nil {
		panic("lat variable not found")
	}
	vsd := vs.Dimensions()
	for i, g := range vsd {
		fmt.Println(i, g)
	}
	/* Using GetSlice
	vss,_ := vs.GetSlice(0,3)
	vssd := vss.([][][][]float32)
	spew.Dump(vssd)
	*/

	vssdd, _ := vs.Values()
	vsddd := vssdd.([][][][]float32)
	for i := 0; i < 57; i++ {
		spew.Dump(vsddd[0][i][ilat][ilong])
	}
	//See this:
	//https://github.com/batchatco/go-native-netcdf/blob/5849c1f424b12bc9f6441723ba6297f2d484e5d2/netcdf/cdf/slicer_test.go

	fmt.Println("------ VALUES ---------")
	/*
		// Cast the data into a Go type we can use
		lats, has := vr.Values.([]float32)
		if !has {
			panic("latitude data not found")
		}
		for i, lat := range lats {
			fmt.Println(i, lat)
		}
	*/
}
