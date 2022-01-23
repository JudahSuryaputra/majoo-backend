package main

import (
	"errors"

	"github.com/google/martian/log"
	"gorm.io/gorm"
)

/* yang harus diperbaiki

ada beberapa line of code yang repetitif dan menurut saya tidak perlu diketik secara berulang
e.g. :

	Var area int -> var dengan V huruf kecil
	area = 0 -> sebenarnya tidak perlu bisa langsung assign seperti contoh di bawah

	var area := param1 * param2
	ar.AreaValue = area			-> bisa langsung ar.AreaValue = param1 * param 2 agar lebih hemat memory


	dan lainnya dapat dilihat di code seperti di bawah
*/

type Area struct {
	ID        int    `gorm:"column:id;primaryKey;"`
	AreaValue int    `gorm:"column:area_value"`
	AreaType  string `gorm:"column:type"`
	DB        *gorm.DB
}

func (ar *Area) InsertArea(param1, param2 int, shape string) (err error) {
	ar.AreaType = shape
	switch shape {
	case "persegi panjang":
		ar.AreaValue = param1 * param2
	case "persegi":
		ar.AreaValue = param1 * param2
	case "segitiga":
		ar.AreaValue = (param1 * param2) / 2
		ar.AreaType = "segitiga"
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefined data"
	}

	err = ar.DB.Create(&ar).Error
	if err != nil {
		return err
	}

	return nil
}

// dipanggil dengan cara berikut

func callFunction() error {
	var area Area
	err := area.InsertArea(10, 10, "persegi")
	if err != nil {
		log.Error().Msg(err.Error())
		err = errors.New(en.ERROR_DATABASE)
		return err
	}
}
