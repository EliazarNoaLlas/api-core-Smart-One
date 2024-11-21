/*
 * File: views_handler_helper_validation.go
 * Author: Melendez
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * Entities helper to handler for views.
 *
 * Last Modified: 2023-11-23
 */

package rest

type deleteViewsResult struct {
	Data   bool `json:"data" binding:"required"`
	Status int  `json:"status" binding:"required"`
}
