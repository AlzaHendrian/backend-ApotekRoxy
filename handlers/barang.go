package handlers

import (
	barangdto "backend/dto/barang"
	dto "backend/dto/result"
	"backend/models"
	"backend/repositories"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type handlerBarang struct {
	BarangRepository repositories.BarangRepository
}

func HandlerBarang(BarangRepository repositories.BarangRepository) *handlerBarang {
	return &handlerBarang{BarangRepository}
}

func (h *handlerBarang) AddBarang(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	harga, _ := strconv.Atoi(c.FormValue("harga"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))

	request := barangdto.CreateBarangRequest{
		ID:    id,
		Nama:  c.FormValue("nama"),
		Harga: harga,
		Qty:   qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	barang := models.Barang{
		ID:    id,
		Nama:  request.Nama,
		Harga: harga,
		Qty:   qty,
	}

	barang, err = h.BarangRepository.AddBarang(barang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Failed to add barang"})
	}

	barang, _ = h.BarangRepository.GetBarang(barang.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: barang})
}

func (h *handlerBarang) GetBarang(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	barang, err := h.BarangRepository.GetBarang(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: barang})
}

func (h *handlerBarang) GetAllBarang(c echo.Context) error {
	searchName := c.QueryParam("search_name")

	var barangs []models.Barang

	var err error
	if searchName != "" {
		barangs, err = h.BarangRepository.GetBarangByName(searchName)
	} else {
		barangs, err = h.BarangRepository.GetAllBarang()
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: barangs})
}

func (h *handlerBarang) DeleteBarang(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	barang, err := h.BarangRepository.GetBarang(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.BarangRepository.DeleteBarang(barang, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerBarang) UpdateBarang(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	harga, _ := strconv.Atoi(c.FormValue("harga"))
	qty, _ := strconv.Atoi(c.FormValue("qty"))
	request := barangdto.CreateBarangRequest{
		ID:    id,
		Nama:  c.FormValue("nama"),
		Harga: harga,
		Qty:   qty,
	}

	fmt.Println("start update barang")

	fmt.Println("ini request :", request)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	barang, err := h.BarangRepository.GetBarang(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	barang.ID = request.ID
	barang.Nama = request.Nama
	barang.Harga = request.Harga
	barang.Qty = request.Qty

	data, err := h.BarangRepository.UpdateBarang(barang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
