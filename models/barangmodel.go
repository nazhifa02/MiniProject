package models

import (
	"database/sql"
	"fmt"

	"github.com/nazhifa02/MiniProject/config"
	"github.com/nazhifa02/MiniProject/entities"
)

type BarangModel struct {
	conn *sql.DB
}

func NewBarangModel() *BarangModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &BarangModel{
		conn: conn,
	}
}

func (p *BarangModel) FindAll() ([]entities.Barang, error) {

	rows, err := p.conn.Query("select * from barang")
	if err != nil {
		return []entities.Barang{}, err
	}
	defer rows.Close()

	var dataBarang []entities.Barang
	for rows.Next() {
		var barang entities.Barang
		rows.Scan(&barang.Id,
			&barang.Nama,
			&barang.Jenis,
			&barang.Harga,
			&barang.Stok)

		dataBarang = append(dataBarang, barang)
	}

	return dataBarang, nil

}

func (p *BarangModel) Create(barang entities.Barang) bool {

	result, err := p.conn.Exec("insert into barang (Nama, Jenis, Harga, Stok) values(?,?,?,?)",
		barang.Nama, barang.Jenis, barang.Harga, barang.Stok)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *BarangModel) Find(id int64, barang *entities.Barang) error {

	return p.conn.QueryRow("select * from barang where id = ?", id).Scan(
		&barang.Id,
		&barang.Nama,
		&barang.Jenis,
		&barang.Harga,
		&barang.Stok)
}

func (p *BarangModel) Update(barang entities.Barang) error {

	_, err := p.conn.Exec(
		"update barang set Nama = ?, Jenis= ?, Harga = ?, Stok= ? where id = ?",
		barang.Nama, barang.Jenis, barang.Harga, barang.Stok, barang.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *BarangModel) Delete(id int64) {
	p.conn.Exec("delete from barang where id = ?", id)
}
