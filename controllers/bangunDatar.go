package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func SegitigaSamaSisi(c *gin.Context) {
	var (
		result gin.H
	)
	alas, err := strconv.Atoi(c.Query("alas"))
	if err != nil {
		panic(err)
	}
	tinggi, err := strconv.Atoi(c.Query("tinggi"))
	if err != nil {
		panic(err)
	}
	hitung := c.Query("hitung")

	if hitung == "luas" {
		result = gin.H{
			"luas": (alas * tinggi) / 2,
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 3 * alas,
		}
	}

	c.JSON(http.StatusOK, result)
}

func Persegi(c *gin.Context) {
	var (
		result gin.H
	)
	sisi, err := strconv.Atoi(c.Query("sisi"))
	if err != nil {
		panic(err)
	}
	hitung := c.Query("hitung")

	if hitung == "luas" {
		result = gin.H{
			"luas": sisi * sisi,
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 4 * sisi,
		}
	}

	c.JSON(http.StatusOK, result)
}

func PersegiPanjang(c *gin.Context) {
	var (
		result gin.H
	)
	panjang, err := strconv.Atoi(c.Query("panjang"))
	if err != nil {
		panic(err)
	}
	lebar, err := strconv.Atoi(c.Query("lebar"))
	if err != nil {
		panic(err)
	}
	hitung := c.Query("hitung")

	if hitung == "luas" {
		result = gin.H{
			"luas": panjang * lebar,
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 2 * (panjang + lebar),
		}
	}

	c.JSON(http.StatusOK, result)
}

func Lingkaran(c *gin.Context) {
	var (
		result gin.H
	)
	radius, err := strconv.Atoi(c.Query("radius"))
	if err != nil {
		panic(err)
	}
	hitung := c.Query("hitung")

	if hitung == "luas" {
		result = gin.H{
			"luas": math.Pi * float64(radius*radius),
		}
	} else if hitung == "keliling" {
		result = gin.H{
			"keliling": 2 * math.Pi * float64(radius),
		}
	}

	c.JSON(http.StatusOK, result)
}
