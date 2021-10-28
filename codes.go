package main

const routeUpdate = `
	// /%[1]s routes
	%[1]sRouter := r.PathPrefix("/%[1]s").Subrouter()
	%[1]sRouter.HandleFunc("", handlers.Handler(cfg, &services.Create%[2]sReq{})).Methods(http.MethodPost)
	%[1]sRouter.HandleFunc("", handlers.Handler(cfg, &services.Get%[2]sReq{})).Methods(http.MethodGet)
	%[1]sRouter.HandleFunc("", handlers.Handler(cfg, &services.Delete%[2]sReq{})).Methods(http.MethodDelete)
	%[1]sRouter.HandleFunc("", handlers.Handler(cfg, &services.Update%[2]sReq{})).Methods(http.MethodPut)
	// ------> codegen_line_tracker ------->`

const getRouteContent = `
package services

import (
	"context"
	"net/http"

	"coinprofile/waas/src/config"
	"coinprofile/waas/src/models"
	"coinprofile/waas/src/schema"
	"coinprofile/waas/src/utils"
	"coinprofile/waas/src/validation"
)

type Get%[1]sReq struct {
	schema.%[1]s
}

type Get%[1]sRes struct {
	schema.%[1]s
}

func (d *Get%[1]sReq) Validate() (isValid bool, errs []error) {
	v := validation.NewValidator()
	errs = v.
		Exec()
	return len(errs) == 0, errs
}

// Controller returns the result of the logic
func (d *Get%[1]sReq) Controller(ctx context.Context, cfg *config.Config) (status int, msg string, data interface{}, err error) {

	%[2]s, err := models.Get%[1]s(d.ID)

	if err != nil {
		return http.StatusInternalServerError, "Get %[1]s failed", data, err
	}

	if %[2]s == nil {
		return http.StatusNotFound, "%[1]s not found", Get%[1]sRes{}, err
	}

	response := &Get%[1]sRes{
		%[1]s: *%[2]s,
	}

	return http.StatusOK, "Request completed", response, err
}

func (d *Get%[1]sReq) GetParamsMap() utils.QueryMap {
	return utils.QueryMap{}
}

func (d *Get%[1]sReq) New() utils.RequestData {
	instance := Get%[1]sReq{}
	return &instance
}
`

const deleteRouteContent = `
package services

import (
	"context"
	"net/http"

	"coinprofile/waas/src/config"
	"coinprofile/waas/src/models"
	"coinprofile/waas/src/schema"
	"coinprofile/waas/src/utils"
	"coinprofile/waas/src/validation"
)

type Delete%[1]sReq struct {
	schema.%[1]s
}

type Delete%[1]sRes struct {}

func (d *Delete%[1]sReq) Validate() (isValid bool, errs []error) {
	v := validation.NewValidator()
	errs = v.
		Exec()
	return len(errs) == 0, errs
}

// Controller returns the result of the logic
func (d *Delete%[1]sReq) Controller(ctx context.Context, cfg *config.Config) (status int, msg string, data interface{}, err error) {

	err = models.Delete%[1]s(d.ID)

	if err != nil {
		return http.StatusInternalServerError, "Delete %[1]s failed", data, err
	}

	return http.StatusOK, "Request completed", data, err
}

func (d *Delete%[1]sReq) GetParamsMap() utils.QueryMap {
	return utils.QueryMap{}
}

func (d *Delete%[1]sReq) New() utils.RequestData {
	instance := Delete%[1]sReq{}
	return &instance
}
`

const postRouteContent = `
package services

import (
	"context"
	"net/http"

	"coinprofile/waas/src/config"
	"coinprofile/waas/src/models"
	"coinprofile/waas/src/schema"
	"coinprofile/waas/src/utils"
	"coinprofile/waas/src/validation"
)

type Create%[1]sReq struct {
	schema.%[1]s
}

type Create%[1]sRes struct {}

func (d *Create%[1]sReq) Validate() (isValid bool, errs []error) {
	v := validation.NewValidator()
	errs = v.
		Exec()
	return len(errs) == 0, errs
}

// Controller returns the result of the logic
func (d *Create%[1]sReq) Controller(ctx context.Context, cfg *config.Config) (status int, msg string, data interface{}, err error) {

	err = models.Create%[1]s(&d.%[1]s)

	if err != nil {
		return http.StatusInternalServerError, "Create %[1]s failed", data, err
	}

	return http.StatusOK, "Request completed", data, err
}

func (d *Create%[1]sReq) GetParamsMap() utils.QueryMap {
	return nil
}

func (d *Create%[1]sReq) New() utils.RequestData {
	instance := Create%[1]sReq{}
	return &instance
}
`

const updateRouteContent = `
package services

import (
	"context"
	"net/http"

	"coinprofile/waas/src/config"
	"coinprofile/waas/src/models"
	"coinprofile/waas/src/schema"
	"coinprofile/waas/src/utils"
	"coinprofile/waas/src/validation"
)

type Update%[1]sReq struct {
	schema.%[1]s
}

type Update%[1]sRes struct {}

func (d *Update%[1]sReq) Validate() (isValid bool, errs []error) {
	v := validation.NewValidator()
	errs = v.
		Exec()
	return len(errs) == 0, errs
}

// Controller returns the result of the logic
func (d *Update%[1]sReq) Controller(ctx context.Context, cfg *config.Config) (status int, msg string, data interface{}, err error) {

	err = models.Update%[1]s(d.%[1]s.ID, &d.%[1]s)

	if err != nil {
		return http.StatusInternalServerError, "Update %[1]s failed", data, err
	}

	return http.StatusOK, "Request completed", data, err
}

func (d *Update%[1]sReq) GetParamsMap() utils.QueryMap {
	return nil
}

func (d *Update%[1]sReq) New() utils.RequestData {
	instance := Update%[1]sReq{}
	return &instance
}
`

const modelContent = `
package models

import (
	"coinprofile/waas/src/db"
	"coinprofile/waas/src/schema"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const %[2]sCollection = "%[2]ss"

func Create%[1]s(%[2]s *schema.%[1]s) error {
	c := db.Store.Database.Collection(%[2]sCollection)
	result, err := c.InsertOne(context.TODO(), %[2]s)
	if err != nil {
		return err
	}
	%[2]s.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return err
}

func Get%[1]s(name string) (*schema.%[1]s, error) {
	c := db.Store.Database.Collection(%[2]sCollection)
	query := bson.M{}
	doc := schema.%[1]s{}
	err := c.FindOne(context.TODO(), query).Decode(&doc)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return &doc, err
	}
	return &doc, err
}

func find%[1]s(query primitive.M) ([]*schema.%[1]s, error) {
	c := db.Store.Database.Collection(%[2]sCollection)
	ctx := context.TODO()
	cursor, err := c.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	docs := []*schema.%[1]s{}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var doc *schema.%[1]s
		err = cursor.Decode(doc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func Find%[1]sByID(ID string) (doc *schema.%[1]s, err error) {
	_ID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": _ID}
	docs, err := find%[1]s(query)
	if err != nil {
		return nil, err
	}
	if len(docs) > 0 {
		return docs[0], nil
	}
	return doc, nil
}

func Delete%[1]s(name string) (err error) {
	c := db.Store.Database.Collection(%[2]sCollection)
	query := bson.M{}
	_, err = c.DeleteOne(context.TODO(), query)
	return err
}

func Update%[1]s(name string, %[2]s *schema.%[1]s) (err error) {
	c := db.Store.Database.Collection(%[2]sCollection)
	query := bson.M{}
	update := bson.M{
		"$set": bson.M{

		},
	}
	err = c.FindOneAndUpdate(context.TODO(), query, update).Decode(%[2]s)
	return err
}

func Upsert%[1]s(%[2]s *schema.%[1]s) error {
	c := db.Store.Database.Collection(%[2]sCollection)
	filter := bson.M{}
	upsert := true
	opts := options.FindOneAndUpdateOptions{
		Upsert: &upsert,
	}
	b, err := bson.Marshal(%[2]s)
	if err != nil {
		return err
	}
	var update bson.M
	err = bson.Unmarshal(b, &update)
	if err != nil {
		return err
	}
	query := bson.D{{Key: "$set", Value: update}}
	err = c.FindOneAndUpdate(context.TODO(), filter, query, &opts).Decode(%[2]s)
	return err
}
`
