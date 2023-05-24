package barangcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/nazhifa02/MiniProject/libraries"

	"github.com/nazhifa02/MiniProject/models"

	"github.com/nazhifa02/MiniProject/entities"
)

var validation = libraries.NewValidation()
var barangModel = models.NewBarangModel()

func Index(response http.ResponseWriter, request *http.Request) {

	barang, _ := barangModel.FindAll()

	data := map[string]interface{}{
		"barang": barang,
	}

	temp, err := template.ParseFiles("views/barang/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/barang/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var barang entities.Barang
		barang.Nama = request.Form.Get("Nama")
		barang.Jenis = request.Form.Get("Jenis")
		barang.Harga = request.Form.Get("Harga")
		stokStr := request.Form.Get("Stok")
		stok, _ := strconv.ParseInt(stokStr, 10, 64)
		barang.Stok = int64(stok)

		var data = make(map[string]interface{})

		vErrors := validation.Struct(barang)

		if vErrors != nil {
			data["barang"] = barang
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data barang berhasil disimpan"
			barangModel.Create(barang)
		}

		temp, _ := template.ParseFiles("views/barang/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var barang entities.Barang
		barangModel.Find(id, &barang)

		data := map[string]interface{}{
			"barang": barang,
		}

		temp, err := template.ParseFiles("views/barang/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var barang entities.Barang
		barang.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		barang.Nama = request.Form.Get("Nama")
		barang.Jenis = request.Form.Get("Jenis")
		barang.Harga = request.Form.Get("Harga")
		stokStr := request.Form.Get("Stok")
		stok, _ := strconv.ParseInt(stokStr, 10, 64)
		barang.Stok = int64(stok)

		var data = make(map[string]interface{})

		vErrors := validation.Struct(barang)

		if vErrors != nil {
			data["barang"] = barang
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data barang berhasil diperbarui"
			barangModel.Update(barang)
		}

		temp, _ := template.ParseFiles("views/barang/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	barangModel.Delete(id)

	http.Redirect(response, request, "/barang", http.StatusSeeOther)
}
