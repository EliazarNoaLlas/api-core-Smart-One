/*
 * File: error_log_message_entity.go
 * Author: Jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Here are all the custom error messages.
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"errors"
	"net/http"
)

type CustomError struct {
	Description string
	Type        TypeError
	HTTPErrCode int
}

type TypeError string

const Info TypeError = "info"
const Warning TypeError = "warning"
const Error TypeError = "error"
const Fatal TypeError = "fatal"

type SCPError string
type SCPMapTags string

func (e *SCPError) String() string {
	return string(*e)
}

func (e *SCPError) Error() error {
	return errors.New(e.String())
}

var ErrSCPUnknown SCPError = "SCP-1000"
var ErrSCP1001 SCPError = "SCP-1001"
var ErrSCP2009 SCPError = "SCP-2009"
var ErrSCP2900 SCPError = "SCP-2900"
var ErrSCP2001 SCPError = "SCP-2001"
var ErrSCP2002 SCPError = "SCP-2002"
var ErrSCP2003 SCPError = "SCP-2003"
var ErrSCP2004 SCPError = "SCP-2004"
var ErrSCP2005 SCPError = "SCP-2005"
var ErrSCP2006 SCPError = "SCP-2006"
var ErrSCP2007 SCPError = "SCP-2007"
var ErrSCP2008 SCPError = "SCP-2008"
var ErrSCP2010 SCPError = "SCP-2010"
var ErrSCP2011 SCPError = "SCP-2011"
var ErrSCP2012 SCPError = "SCP-2012"
var ErrSCP2013 SCPError = "SCP-2013"
var ErrSCP2014 SCPError = "SCP-2014"
var ErrSCP2015 SCPError = "SCP-2015"
var ErrSCP2016 SCPError = "SCP-2016"
var ErrSCP2017 SCPError = "SCP-2017"
var ErrSCP2018 SCPError = "SCP-2018"
var ErrSCP2019 SCPError = "SCP-2019"
var ErrSCP2020 SCPError = "SCP-2020"
var ErrSCP2021 SCPError = "SCP-2021"
var ErrSCP2022 SCPError = "SCP-2022"
var ErrSCP2023 SCPError = "SCP-2023"
var ErrSCP2024 SCPError = "SCP-2024"
var ErrSCP2025 SCPError = "SCP-2025"
var ErrSCP2026 SCPError = "SCP-2026"
var ErrSCP2027 SCPError = "SCP-2027"
var ErrSCP2028 SCPError = "SCP-2028"
var ErrSCP2029 SCPError = "SCP-2029"
var ErrSCP2030 SCPError = "SCP-2030"
var ErrSCP2031 SCPError = "SCP-2031"
var ErrSCP3000 SCPError = "SCP-3000"
var ErrSCP1002 SCPError = "SCP-1002"

// ErrorMap stores the error codes and descriptions.
var ErrorMap = map[SCPError]CustomError{
	ErrSCP1001: {Description: "INTERNAL-SERVER-ERROR", Type: Error, HTTPErrCode: http.StatusInternalServerError},
	ErrSCP1002: {Description: "INTERNAL-SERVER-ERROR", Type: Error, HTTPErrCode: http.StatusInternalServerError},
	ErrSCP2009: {Description: "NOT FOUND", Type: Error, HTTPErrCode: http.StatusInternalServerError},
	ErrSCP2001: {Description: "INPUT-INVALID", Type: Error, HTTPErrCode: http.StatusInternalServerError},
	ErrSCP3000: {Description: "INPUT-REQUIRED", Type: Error, HTTPErrCode: http.StatusInternalServerError},
	"SCP-2002": {Description: "AUTH-FAILED"},
	"SCP-2003": {Description: "AUTH-INVALID"},
	"SCP-2004": {Description: "AUTH-EXPIRED"},
	"SCP-2005": {Description: "AUTH-UNAUTHORIZED"},
	"SCP-2006": {Description: "AUTH-UNAUTHENTICATED"},
	ErrSCP2900: {Description: "UNMARSHALL ERROR", Type: Warning, HTTPErrCode: http.StatusBadRequest},
	ErrSCP2004: {
		Description: "THERE IS ALREADY A SCHEDULE WITH THAT DATA",
		Type:        Warning, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2007: {
		Description: "THE QUOTA HAS ALREADY BEEN TAKEN",
		Type:        Warning, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2008: {
		Description: "FAILED TO REGISTER THE MEDICAL APPOINTMENT\n",
		Type:        Warning, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2010: {
		Description: "THE CLIENT HAS ALREADY BEEN REGISTERED\n",
		Type:        Warning, HTTPErrCode: http.StatusUnprocessableEntity,
	},
	ErrSCP2005: {
		Description: "NOT FOUND",
		Type:        Error, HTTPErrCode: http.StatusNotFound,
	},
	ErrSCP2006: {
		Description: "SCHEDULING IS SUBJECT TO PRIOR APPOINTMENT. IT WILL NOT BE POSSIBLE TO DISABLE IT",
		Type:        Error, HTTPErrCode: http.StatusUnprocessableEntity,
	},
	ErrSCP2011: {
		Description: "AUTHENTICATION TOKEN IS MISSING",
		Type:        Error, HTTPErrCode: http.StatusNotFound,
	},
	ErrSCP2012: {
		Description: "",
		Type:        Error, HTTPErrCode: http.StatusNotFound,
	},
	ErrSCP2013: {
		Description: "YOU DON'T HAVE AN OPEN CASH DESK",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2014: {
		Description: "EXIST A OPEN CASH DESK WITH THIS SALES POINT",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2015: {
		Description: "EXIST A OPEN CASH DESK WITH SERIES IN USE",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2016: {
		Description: "EXISTS CASHIER ASSOCIATED IN THIS SALES POINT",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2017: {
		Description: "THIS CASH DESK HAS NO INCOMES OR EXPENSES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2018: {
		Description: "EXIST A OPEN CASH DESK WITH THIS SERIES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2019: {
		Description: "EXIST AN INCOMES WITH THIS SERIES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2020: {
		Description: "EXIST AN INVOICES WITH THIS SERIES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2021: {
		Description: "EXIST AN CREDIT NOTES WITH THIS SERIES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2022: {
		Description: "AT LEAST ONE SALES POINT IS NECESSARY",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2023: {
		Description: "EXIST A OPEN CASH DESK WITH THIS CASHIER IN USE",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2024: {
		Description: "EXIST A OPEN CASH DESK IN SALES POINT WITH THIS CASHIER IN USE",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2025: {
		Description: "EXIST A SALES POINT WITH THIS SERIES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2026: {
		Description: "EXIST A CASHIER WITH THIS SERIES",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2027: {
		Description: "EXIST A APPOINTMENT FOR THIS PATIENT WITH THE SAME SPECIALITY AND DATE",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2028: {
		Description: "EXIST A SALES POINTS WITH THE SAME NAME",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2029: {
		Description: "EXIST A PATIENT WITH THE SAME DOCUMENT",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2030: {
		Description: "THIS PATIENT HAS A REQUEST FOR HOSPITALIZATION",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	ErrSCP2031: {
		Description: "THIS PATIENT IS HOSPITALIZED",
		Type:        Error, HTTPErrCode: http.StatusBadRequest,
	},
	// Add more error codes and descriptions as needed
}

func (e *SCPMapTags) CustomError() CustomError {
	err, found := ErrorMapTags[*e]
	if !found {
		return CustomError{}
	}
	err2, found2 := ErrorMap[err]
	if !found2 {
		return CustomError{}
	}
	return err2
}

func (e *SCPMapTags) Error() SCPError {
	err, found := ErrorMapTags[*e]
	if !found {
		return ErrSCPUnknown
	}
	return err
}

var ErrTAGRequired SCPMapTags = "required"

var ErrorMapTags = map[SCPMapTags]SCPError{
	ErrTAGRequired: ErrSCP3000,
}
