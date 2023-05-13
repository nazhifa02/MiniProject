package main

import (
	"net/http"

	"github.com/jeypc/go-crud/controllers/barangcontroller/"
	"github.com/jeypc/go-crud/controllers/penjualcontroller/"
)

func main() {

	http.HandleFunc("/", penjualcontroller.Index)
	http.HandleFunc("/penjual", penjualcontroller.Index)
	http.HandleFunc("/penjual/index", penjualcontroller.Index)
	http.HandleFunc("/penjual/add", penjualcontroller.Add)
	http.HandleFunc("/penjual/edit", penjualcontroller.Edit)
	http.HandleFunc("/penjual/delete", penjualcontroller.Delete)

	http.HandleFunc("/", barangcontroller.Index)
	http.HandleFunc("/barang", barangcontroller.Index)
	http.HandleFunc("/barang/index", barangcontroller.Index)
	http.HandleFunc("/barang/add", barangcontroller.Add)
	http.HandleFunc("/barang/edit", barangcontroller.Edit)
	http.HandleFunc("/barang/delete", barangcontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
