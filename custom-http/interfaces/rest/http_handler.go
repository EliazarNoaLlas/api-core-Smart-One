/*
 * File: http_handler.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the http handler for the application.
 *
 * Last Modified: 2023-11-10
 */

package httpHandler

import (
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	errorLogDomain "gitlab.smartcitiesperu.com/smartone/api-core/error-log/domain"
)

type ErrorResult struct {
	Error  ErrorResponse `json:"error"`
	Status int           `json:"status"`
}

type StatusResult struct {
	Status int `json:"status" binding:"required" example:"200"`
}

type StatusIdResult struct {
	Data   string `json:"data" binding:"required" example:"201"`
	Status int    `json:"status" binding:"required"`
}

type IntIdResponse struct {
	Id int `json:"id" binding:"required"`
}

type StringIdResponse struct {
	Id string `json:"id" binding:"required"`
}

type DataResultStringID struct {
	Data   StringIdResponse `json:"data" binding:"required"`
	Status int              `json:"status" binding:"required"`
}

type DataResultIntId struct {
	Data   IntIdResponse `json:"data" binding:"required"`
	Status int           `json:"status" binding:"required"`
}

type ErrorResponse struct {
	// code error
	Code string `json:"code" binding:"required" example:"SCP-1001"`
	// message error
	Message *string `json:"message" example:"This is a example error"`
	// field errors array
	FieldErrors []FieldError `json:"fieldErrors" binding:"required"`
}

type FieldError struct {
	// field error
	Field string `json:"field" binding:"required" example:"attribute field example"`
	// error code
	ErrorCode string `json:"errorCode" binding:"required" example:"attribute error code example"`
}

func Error(
	errCode errorLogDomain.SCPError,
	fieldError []FieldError,
	ctx context.Context,
) ErrorResult {
	scpErr := errorLogDomain.ErrorMap[errCode]
	errResponse := ErrorResponse{
		Code:        string(errCode),
		Message:     &scpErr.Description,
		FieldErrors: fieldError,
	}

	res := ErrorResult{
		Status: scpErr.HTTPErrCode,
		Error:  errResponse,
	}

	RegisterFieldsLog(ctx, scpErr.Type, *errResponse.Message)

	return res
}

func HandleResponseV2(c *gin.Context, status int, data interface{}) {
	RegisterFieldsLog(c, "info", "Request success")
}

func RegisterFieldsLog(c context.Context, logType errorLogDomain.TypeError, message string) {
	fields := log.Fields{
		"project": os.Getenv("PROJECT"),
	}
	requestID, hasRequestID := c.Value("request_id").(string)
	if hasRequestID {
		fields["request_id"] = requestID
	}
	requestEndAt := time.Now()
	requestStartAt, hasRequestStartAt := c.Value("request_start_at").(time.Time)
	if hasRequestStartAt {
		requestDuration := requestEndAt.Sub(requestStartAt)
		durationMilliseconds := requestDuration.Milliseconds()
		fields["request_start_at"] = requestStartAt
		fields["request_duration"] = durationMilliseconds
	}
	fields["request_end_at"] = requestEndAt
	if logType == errorLogDomain.Info {
		log.WithFields(fields).Info(message)
	}
	if logType == errorLogDomain.Error {
		fields["error"] = message
		log.WithFields(fields).Error(message)
	}
	if logType == errorLogDomain.Warning {
		fields["error"] = message
		log.WithFields(fields).Warning(message)
	}
	if logType == errorLogDomain.Fatal {
		fields["error"] = message
		log.WithFields(fields).Fatal(message)
	}
}

func RegisterFieldsRepositoryLog(ctx context.Context, logType errorLogDomain.TypeError, message string) {
	fields := log.Fields{
		"project": os.Getenv("PROJECT"),
	}
	requestID, hasRequestID := ctx.Value("request_id").(string)
	if hasRequestID {
		fields["request_id"] = requestID
	}
	requestEndAt := time.Now()
	requestStartAt, hasRequestStartAt := ctx.Value("request_start_at").(time.Time)
	if hasRequestStartAt {
		requestDuration := requestEndAt.Sub(requestStartAt)
		durationMilliseconds := requestDuration.Milliseconds()
		fields["request_start_at"] = requestStartAt
		fields["request_duration"] = durationMilliseconds
	}
	fields["request_end_at"] = requestEndAt
	if logType == "info" {
		log.WithFields(fields).Info(message)
	}
	if logType == "error" {
		fields["error"] = message
		log.WithFields(fields).Error(message)
	}
	if logType == "warning" {
		fields["error"] = message
		log.WithFields(fields).Warning(message)
	}
}
