package respository

import (
	"Quiz-3/structs"
	"database/sql"
	"errors"
	"regexp"
	"time"
)

func GetAllBooks(db *sql.DB) (err error, results []structs.Book) {
	sql := "select * from book order by id asc"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var book = structs.Book{}

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.CategoryID)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	sql := "insert into book (title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id) values ($1, $2,$3,$4,$5,$6,$7,$8,$9,$10)"
	book.CreatedAt = time.Now()
	book.UpdatedAt = book.CreatedAt
	if !(book.ReleaseYear >= 1980 && book.ReleaseYear <= 2021) {
		return errors.New("Tahun harus lebih besar sama dengan 1980 dan lebih kecil atau sama dengan 2021")
	}
	match, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", book.ImageUrl)
	if !match {
		return errors.New("url tidak valid")
	}
	if book.TotalPage <= 100 && book.TotalPage > 0 {
		book.Thickness = "tipis"
	} else if book.TotalPage > 100 && book.TotalPage <= 200 {
		book.Thickness = "sedang"
	} else if book.TotalPage > 200 {
		book.Thickness = "tebal"
	} else {
		return errors.New("Total halaman buku tidak valid")
	}
	errs := db.QueryRow(sql,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CreatedAt,
		book.UpdatedAt,
		book.CategoryID)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "update book set title=$2, description=$3, image_url=$4, release_year=$5, price=$6, total_page=$7, thickness=$8, updated_at=$9, category_id=$10 where id = $1 "
	book.UpdatedAt = time.Now()
	if !(book.ReleaseYear >= 1980 && book.ReleaseYear <= 2021) {
		return errors.New("Tahun harus lebih besar sama dengan 1980 dan lebih kecil atau sama dengan 2021")
	}
	match, _ := regexp.MatchString("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)", book.ImageUrl)
	if !match {
		return errors.New("url tidak valid")
	}
	if book.TotalPage <= 100 && book.TotalPage > 0 {
		book.Thickness = "tipis"
	} else if book.TotalPage > 100 && book.TotalPage <= 200 {
		book.Thickness = "sedang"
	} else if book.TotalPage > 200 {
		book.Thickness = "tebal"
	} else {
		return errors.New("Total halaman buku tidak valid")
	}
	errs := db.QueryRow(sql,
		book.ID,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.UpdatedAt,
		book.CategoryID)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "delete from book where id= $1"
	errs := db.QueryRow(sql, book.ID)

	return errs.Err()
}
