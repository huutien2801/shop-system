package models

import (
	"github.com/huutien2801/shop-system/entities"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type BookModel struct {
	Db         *mgo.Database
	Collection string
}

func (bookModel BookModel) FindAll() (books []entities.Book, err error) {
	err = bookModel.Db.C(bookModel.Collection).Find(bson.M{}).All(&books)
	return
}

func (bookModel BookModel) FindById(id string) (books entities.Book, err error) {
	err = bookModel.Db.C(bookModel.Collection).FindId(bson.ObjectIdHex(id)).One(&books)
	return
}

func (bookModel BookModel) Create(book *entities.Book) error {
	err := bookModel.Db.C(bookModel.Collection).Insert(&book)
	return err
}

func (bookModel BookModel) Update(book *entities.Book) error {
	err := bookModel.Db.C(bookModel.Collection).UpdateId(book.Id, &book)
	return err
}

func (bookModel BookModel) Delete(book entities.Book) error {
	err := bookModel.Db.C(bookModel.Collection).Remove(book)
	return err
}
