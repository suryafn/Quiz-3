# Quiz-3
## _Surya Nainggolan_

## Bangun Datar

| Jenis | Path | Parameter | Url|
| ------ | ------ | ------ | ------ |
| Segitiga | /segitiga-sama-sisi | alas, tinggi, hitung | localhost:8080/bangun-datar/segitiga-sama-sisi?hitung=luas&alas=7&tinggi=10|
| Persegi | /persegi | sisi, hitung | localhost:8080/bangun-datar/persegi?hitung=luas&sisi=6 |
| Persegi Panjang | /persegi-panjang | panjang, lebar, hitung | localhost:8080/bangun-datar/persegi-panjang?hitung=luas&panjang=10&lebar=3 |
| Lingkaran | /lingkaran | radius, hitung | localhost:8080/bangun-datar/lingkaran?hitung=luas&radius=10 |

## Category

- Get All Categories
```sh
curl --location --request GET 'http://localhost:8080/categories'
```
- Create Category
```sh
curl --location --request POST 'http://localhost:8080/categories' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "bat"
}'
```
- Update Category
```sh
curl --location --request PUT 'http://localhost:8080/categories/5' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "seram kali diubah"
}'
```

- Delete Category
```sh
curl --location --request DELETE 'http://localhost:8080/categories/2' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "seram kali diubah"
}'
```

- Get Books by Category ID
```sh
curl --location --request GET 'http://localhost:8080/categories/1/books'
```

## Book

- Get All Books
```sh
curl --location --request GET 'http://localhost:8080/books'
```
- Create Book
```sh
curl --location --request POST 'http://localhost:8080/books' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "ini buku",
    "description": "bukulah",
    "image_url": "http://www.gmail.com",
    "release_year": 2000,
    "price": 30000,
    "total_page": 300,
    "category_id": 2
}'
```
- Update Book
```sh
curl --location --request PUT 'http://localhost:8080/books/3' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "ini buku udah diupdate",
    "description": "bukulah",
    "image_url": "http://www.gmail.com",
    "release_year": 2020,
    "price": 30000,
    "total_page": 300,
    "category_id": 1
}'
```

- Delete Book
```sh
curl --location --request DELETE 'http://localhost:8080/books/2' \
--header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "seram kali diubah"
}'
```