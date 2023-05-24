package penjualcontroller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/nazhifa02/MiniProject/libraries"

	"github.com/nazhifa02/MiniProject/models"

	"github.com/nazhifa02/MiniProject/entities"
)

var validation = libraries.NewValidation()
var penjualModel = models.NewPenjualModel()

func Index(response http.ResponseWriter, request *http.Request) {

	penjual, _ := penjualModel.FindAll()

	data := map[string]interface{}{
		"penjual": penjual,
	}

	temp, err := template.ParseFiles("views/penjual/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, data)
}

func Add(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/penjual/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var penjual entities.Penjual
		penjual.Nama = request.Form.Get("Nama")
		NoTeleponStr := request.Form.Get("Stok")
		NoTelepon, _ := strconv.ParseInt(NoTeleponStr, 10, 64)
		penjual.NoTelepon = int64(NoTelepon)
		penjual.Alamat = request.Form.Get("Alamat")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(penjual)

		if vErrors != nil {
			data["penjual"] = penjual
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data penjual berhasil disimpan"
			penjualModel.Create(penjual)
		}

		temp, _ := template.ParseFiles("views/penjual/add.html")
		temp.Execute(response, data)
	}

}

func Edit(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {

		queryString := request.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var penjual entities.Penjual
		penjualModel.Find(id, &penjual)

		data := map[string]interface{}{
			"penjual": penjual,
		}

		temp, err := template.ParseFiles("views/penjual/edit.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, data)

	} else if request.Method == http.MethodPost {

		request.ParseForm()

		var penjual entities.Penjual
		penjual.Id, _ = strconv.ParseInt(request.Form.Get("id"), 10, 64)
		penjual.Nama = request.Form.Get("Nama")
		NoTeleponStr := request.Form.Get("Stok")
		NoTelepon, _ := strconv.ParseInt(NoTeleponStr, 10, 64)
		penjual.NoTelepon = int64(NoTelepon)
		penjual.Alamat = request.Form.Get("Alamat")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(penjual)

		if vErrors != nil {
			data["penjual"] = penjual
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data penjual berhasil diperbarui"
			penjualModel.Update(penjual)
		}

		temp, _ := template.ParseFiles("views/penjual/edit.html")
		temp.Execute(response, data)
	}

}

func Delete(response http.ResponseWriter, request *http.Request) {

	queryString := request.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	penjualModel.Delete(id)

	http.Redirect(response, request, "/penjual", http.StatusSeeOther)
}
