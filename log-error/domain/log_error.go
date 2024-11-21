/*
 * File: log_error.go
 * Author: jesus
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-11-10
 */

package domain

import (
	"errors"
)

var (
	ErrPanic = errors.New("err_panic")
)
