package respository

import (
	"Quiz-3/structs"
	"database/sql"
	"time"
)

func GetAllCategories(db *sql.DB) (err error, results []structs.Category) {
	sql := "select * from category order by id asc"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var category = structs.Category{}

		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt)

		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "insert into category (name,created_at, updated_at) values ($1, $2,$3)"
	category.CreatedAt = time.Now()
	category.UpdatedAt = category.CreatedAt
	errs := db.QueryRow(sql,
		category.Name,
		category.CreatedAt,
		category.UpdatedAt)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "update category set name=$2, updated_at=$3 where id = $1 "
	category.UpdatedAt = time.Now()
	errs := db.QueryRow(sql,
		category.ID,
		category.Name,
		category.UpdatedAt)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "delete from category where id= $1"
	errs := db.QueryRow(sql, category.ID)

	return errs.Err()
}

func GetAllBookByCategoryId(db *sql.DB, category structs.Category) (err error, results []structs.Book) {
	sql := "select * from book where category_id = $1 order by id asc"

	rows, err := db.Query(sql, category.ID)
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
