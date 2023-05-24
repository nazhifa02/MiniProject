package models

import (
	"database/sql"
	"fmt"

	"github.com/nazhifa02/MiniProject/config"
	"github.com/nazhifa02/MiniProject/entities"
)

type PenjualModel struct {
	conn *sql.DB
}

func NewPenjualModel() *PenjualModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PenjualModel{
		conn: conn,
	}
}

func (p *PenjualModel) FindAll() ([]entities.Penjual, error) {

	rows, err := p.conn.Query("select * from penjual")
	if err != nil {
		return []entities.Penjual{}, err
	}
	defer rows.Close()

	var dataPenjual []entities.Penjual
	for rows.Next() {
		var penjual entities.Penjual
		rows.Scan(&penjual.Id,
			&penjual.Nama,
			&penjual.NoTelepon,
			&penjual.Alamat)

		dataPenjual = append(dataPenjual, penjual)
	}

	return dataPenjual, nil

}

func (p *PenjualModel) Create(penjual entities.Penjual) bool {

	result, err := p.conn.Exec("insert into penjual(Nama, NoTelepon, Alamat) values(?,?,?,?)",
		penjual.Nama, penjual.NoTelepon, penjual.Alamat)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *PenjualModel) Find(id int64, penjual *entities.Penjual) error {

	return p.conn.QueryRow("select * from penjual where id = ?", id).Scan(
		&penjual.Id,
		&penjual.Nama,
		&penjual.NoTelepon,
		&penjual.Alamat)

}

func (p *PenjualModel) Update(penjual entities.Penjual) error {

	_, err := p.conn.Exec(
		"update penjual set Nama = ?, Notelepon= ?, Alamat= ? where id = ?",
		penjual.Nama, penjual.NoTelepon, penjual.Alamat, penjual.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PenjualModel) Delete(id int64) {
	p.conn.Exec("delete from penjual where id = ?", id)
}
