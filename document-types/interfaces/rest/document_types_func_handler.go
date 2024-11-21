/*
 * File: document_types_func_usecase.go
 * Author: lady
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-12-07
 */

package rest

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	restCore "gitlab.smartcitiesperu.com/smartone/api-shared/api-core/interfaces/rest"
	httpResponse "gitlab.smartcitiesperu.com/smartone/api-shared/custom-http/interfaces/rest"
	paramsDomain "gitlab.smartcitiesperu.com/smartone/api-shared/params/domain"

	"gitlab.smartcitiesperu.com/smartone/api-core/document-types/domain"
)

// GetDocumentTypes is a method to get document types
// @Summary get document types
// @Description get document types
// @Tags DocumentTypes
// @Accept json
// @Produce json
// @Param number query string false "the number of the document type"
// @Param description query string false "the description of the document type"
// @Param abbreviated_description query string false "the abbreviated description of the document type"
// @Success 200 {object} documentTypeResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/document_types/ [get]
// @Security BearerAuth
func (h documentTypesHandler) GetDocumentTypes(c *gin.Context) {
	ctx := c.Request.Context()

	searchParams := domain.GetDocumentTypeParams{}
	searchParams.QueryParamsToStruct(c.Request, &searchParams)
	pagination := paramsDomain.NewPaginationParams(c.Request)

	documentTypes, paginationRes, err := h.documentTypesUseCase.GetDocumentTypes(ctx, searchParams, pagination)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := documentTypeResult{
		Data:       documentTypes,
		Pagination: *paginationRes,
		Status:     http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// CreateDocumentType is a method to create a document type
// @Summary Create a document type
// @Description Create a document type
// @Tags DocumentTypes
// @Accept json
// @Produce json
// @Param documentTypeId path string true "document type id"
// @Param createDocumentTypeBody body domain.CreateDocumentTypeBody true "Create document type body"
// @Success 201 {object} httpResponse.IdResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/document_types/create_document_types/{documentTypeId} [post]
// @Security BearerAuth
func (h documentTypesHandler) CreateDocumentType(c *gin.Context) {
	ctx := c.Request.Context()
	documentTypeId := c.Param("documentTypeId")
	var documentTypesValidate CreateDocumentTypeValidate
	if err := c.ShouldBindJSON(&documentTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("CreateDocumentType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("CreateDocumentType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	createDocumentTypesBody := domain.CreateDocumentTypeBody{
		Number:                 documentTypesValidate.Number,
		Description:            documentTypesValidate.Description,
		AbbreviatedDescription: documentTypesValidate.AbbreviatedDescription,
		Enable:                 documentTypesValidate.Enable,
	}
	id, err := h.documentTypesUseCase.CreateDocumentType(ctx, documentTypeId, createDocumentTypesBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.IdResult{
		Data:   *id,
		Status: http.StatusCreated,
	}
	restCore.Json(c, http.StatusCreated, res)
}

// UpdateDocumentType is a method to update a document type
// @Summary Update a document type
// @Description Update a document type
// @Tags DocumentTypes
// @Accept json
// @Produce json
// @Param documentTypeId path string true "document type id"
// @Param updateDocumentTypeBody body domain.UpdateDocumentTypeBody true "Update document types body"
// @Success 200 {object} httpResponse.StatusResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/document_types/update_document_types/{documentTypeId} [put]
// @Security BearerAuth
func (h documentTypesHandler) UpdateDocumentType(c *gin.Context) {
	ctx := c.Request.Context()
	documentTypeId := c.Param("documentTypeId")

	var documentTypesValidate UpdateDocumentTypeValidate
	if err := c.ShouldBindJSON(&documentTypesValidate); err != nil {
		validationErrs, errFind := err.(validator.ValidationErrors)
		if !errFind {
			err = h.err.Clone().SetFunction("UpdateDocumentType").SetRaw(errors.New("casting ValidationErrors"))
			restCore.ErrJson(c, err)
			return
		}

		messagesErr := make([]string, 0)
		for _, validationErr := range validationErrs {
			messagesErr = append(messagesErr, validationErr.Field()+" "+validationErr.Tag())
		}
		err = h.err.Clone().SetFunction("UpdateDocumentType").SetMessages(messagesErr)
		restCore.ErrJson(c, err)
		return
	}

	documentTypeBody := domain.UpdateDocumentTypeBody{
		Number:                 documentTypesValidate.Number,
		Description:            documentTypesValidate.Description,
		AbbreviatedDescription: documentTypesValidate.AbbreviatedDescription,
		Enable:                 documentTypesValidate.Enable,
	}
	err := h.documentTypesUseCase.UpdateDocumentType(ctx, documentTypeId, documentTypeBody)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}

	res := httpResponse.StatusResult{
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}

// DeleteDocumentType is a method to delete a document type
// @Summary Delete a document type
// @Description Delete a document type
// @Tags DocumentTypes
// @Accept json
// @Produce json
// @Param documentTypeId path string true "document type id"
// @Success 200 {object} deleteDocumentTypesResult "Success Request"
// @Failure 500 {object} errorDomain.SmartError "Bad Request"
// @Router /api/v1/core/document_types/delete_document_types/{documentTypeId} [delete]
// @Security BearerAuth
func (h documentTypesHandler) DeleteDocumentType(c *gin.Context) {
	ctx := c.Request.Context()
	documentTypeId := c.Param("documentTypeId")
	result, err := h.documentTypesUseCase.DeleteDocumentType(ctx, documentTypeId)
	if err != nil {
		restCore.ErrJson(c, err)
		return
	}
	res := deleteDocumentTypesResult{
		Data:   result,
		Status: http.StatusOK,
	}
	restCore.Json(c, http.StatusOK, res)
}
