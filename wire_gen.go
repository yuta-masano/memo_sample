// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"memo_sample/adapter/db"
	"memo_sample/adapter/logger"
	"memo_sample/adapter/memory"
	"memo_sample/adapter/view/render"
	"memo_sample/interface/api"
	"memo_sample/usecase"
)

// Injectors from injector.go:

func InjectAPIServer() api.API {
	jsonRender := view.NewJSONRender()
	logger := loggersub.NewLogger()
	presenter := api.NewPresenter(jsonRender, logger)
	transactionRepository := db.NewTransactionRepository()
	memoRepository := db.NewMemoRepository()
	tagRepository := db.NewTagRepository()
	memo := usecase.NewMemo(transactionRepository, memoRepository, tagRepository)
	interactor := usecase.NewInteractor(presenter, memo)
	apiAPI := api.NewAPI(interactor, logger)
	return apiAPI
}

// injector.go:

// WireInjectAPI inject api using wire
var WireInjectAPI = wire.NewSet(
	WireInjectUsecaseIterator, api.NewAPI,
)

// WireInjectPresenter inject presenter using wire
var WireInjectPresenter = wire.NewSet(
	WireInjectRender,
	WireInjectLog, api.NewPresenter,
)

// WireInjectMemoUsecase inject memo usecase using wire
var WireInjectMemoUsecase = wire.NewSet(
	WireInjectDBRepository, usecase.NewMemo,
)

// WireInjectUsecaseIterator inject usecase itetator using wire
var WireInjectUsecaseIterator = wire.NewSet(
	WireInjectPresenter,
	WireInjectMemoUsecase, usecase.NewInteractor,
)

// WireInjectInMemoryRepository inject repository using wire
var WireInjectInMemoryRepository = wire.NewSet(memory.NewTransactionRepository, memory.NewMemoRepository, memory.NewTagRepository)

// WireInjectDBRepository inject repository using wire
var WireInjectDBRepository = wire.NewSet(db.NewTransactionRepository, db.NewMemoRepository, db.NewTagRepository)

// WireInjectLog inject log using wire
var WireInjectLog = wire.NewSet(loggersub.NewLogger)

// WireInjectRender inject render using wire
var WireInjectRender = wire.NewSet(view.NewJSONRender)
