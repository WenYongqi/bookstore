package dao

import (
	"bookstore/model"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func GetBooks() (*[]model.Book, error) {
	var books []model.Book
	result := db.Find(&books)
	err := result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return nil, err
	}
	return &books, nil
}

func AddBook(book *model.Book) error {
	result := db.Create(book)
	err := result.Error
	log.Println("Rows affected:", result.RowsAffected)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteBook(bookID int) error {
	result := db.Delete(&model.Book{}, bookID)
	err := result.Error
	log.Println("Rows affected:", result.RowsAffected)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetBookByID(bookID int) (*model.Book, error) {
	book := model.Book{}
	result := db.First(&book, bookID)
	err := result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return nil, err
	}
	return &book, nil
}

func UpdateBook(book *model.Book) error {
	result := db.Save(book)
	log.Println("Rows affected:", result.RowsAffected)
	err := result.Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetPageBooks(pageNo int64) (*model.Page, error) {
	var totalRecord int64
	db.Model(&model.Book{}).Count(&totalRecord)
	pageSize:= int64(4)
	var totalPageNo int64
	totalPageNo = totalRecord / pageSize
	if totalRecord % pageSize != 0 {
		totalPageNo += 1
	}
	if pageNo < 1 {
		log.Println("pageNo " + strconv.FormatInt(pageNo, 10) + " not legal")
		pageNo = 1
	} else if pageNo > totalPageNo {
		log.Println("pageNo " + strconv.FormatInt(pageNo, 10) + " not legal")
		pageNo = totalPageNo
	}
	var books []model.Book
	err := db.Model(&model.Book{}).Limit(int(pageSize)).Offset(int((pageNo - 1) * pageSize)).Find(&books).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	page := model.Page{
		Books: &books,
		PageNo: pageNo,
		PageSize: pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return &page, nil
}

func GetPageBooksByPrice(pageNo int64, minPrice, maxPrice float64) (*model.Page, error) {
	var totalRecord int64
	db.Model(&model.Book{}).Where("price between ? and ?", minPrice, maxPrice).Count(&totalRecord)
	pageSize:= int64(4)
	var totalPageNo int64
	totalPageNo = totalRecord / pageSize
	if totalRecord % pageSize != 0 {
		totalPageNo += 1
	}
	if pageNo < 1 {
		log.Println("pageNo " + strconv.FormatInt(pageNo, 10) + " not legal")
		pageNo = 1
	} else if pageNo > totalPageNo {
		log.Println("pageNo " + strconv.FormatInt(pageNo, 10) + " not legal")
		pageNo = totalPageNo
	}
	var books []model.Book
	err := db.Where("price between ? and ?", minPrice, maxPrice).Limit(int(pageSize)).Offset(int((pageNo - 1) * pageSize)).Find(&books).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	page := model.Page{
		Books: &books,
		PageNo: pageNo,
		PageSize: pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
		MinPrice: strconv.FormatFloat(minPrice, byte('g'), 5, 64),
		MaxPrice: strconv.FormatFloat(maxPrice, byte('g'), 5, 64),
	}
	return &page, nil
}