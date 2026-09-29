package main

import "fmt"

func main () {

	namaorang := map[string]string{
		 //cara membuat map namavariable := map[tipe_data_key]tipe_data_value { }
		"namadepan": "Bintang", //key itu namadepan, value itu Bintang
		"namabelakang": "Akasyah",//key itu namabelakang, value itu Akasyah
	}

	namaorang["Gelar"] = "Dev Ops" //menambah data baru pada map

    fmt.Println(namaorang ["namadepan"], namaorang["namabelakang"], namaorang["Gelar"])
	// fmt.Println(namaorang["namadepan"])
	// fmt.Println(namaorang["namabelakang"])
	// fmt.Println(namaorang["Gelar"])

	minuman := map[string]string{
		"minuman1": "tehbotol",
		"minuman2": "kopi",
		"minuman3": "airputih",
		"minuman4": "mieayam",
	}

	delete (minuman, "minuman4") 
	//delete itu langsung aja caranya delete (nama_map, "key_map_yang_ingin_dihapus")
	fmt.Println(minuman)
}