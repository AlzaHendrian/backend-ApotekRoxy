package repositories

import (
	"backend/models"
	"fmt"

	"gorm.io/gorm"
)

type BarangRepository interface {
	AddBarang(barang models.Barang) (models.Barang, error)
	GetBarang(ID int) (models.Barang, error)
	GetBarangByName(name string) ([]models.Barang, error)
	GetAllBarang() ([]models.Barang, error)
	DeleteBarang(barang models.Barang, ID int) (models.Barang, error)
	UpdateBarang(barang models.Barang) (models.Barang, error)
}

func RepositoryBarang(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddBarang(barang models.Barang) (models.Barang, error) {
	err := r.db.Create(&barang).Error
	if err != nil {
		fmt.Println("Error creating barang:", err)
	}
	return barang, err
}

func (r *repository) GetBarang(ID int) (models.Barang, error) {
	var barang models.Barang
	err := r.db.First(&barang, ID).Error
	if err != nil {
		fmt.Println("Error fetching barang:", err)
	}
	return barang, err
}

func (r *repository) GetAllBarang() ([]models.Barang, error) {
	var barangs []models.Barang
	err := r.db.Find(&barangs).Error

	return barangs, err
}

func (r *repository) GetBarangByName(nama string) ([]models.Barang, error) {
	var barangs []models.Barang
	err := r.db.Where("nama ILIKE ?", "%"+nama+"%").Find(&barangs).Error

	return barangs, err
}

func (r *repository) DeleteBarang(barang models.Barang, ID int) (models.Barang, error) {
	err := r.db.Delete(&barang).Error

	return barang, err
}

func (r *repository) UpdateBarang(barang models.Barang) (models.Barang, error) {
	fmt.Println(barang)
	err := r.db.Save(&barang).Error

	return barang, err
}
