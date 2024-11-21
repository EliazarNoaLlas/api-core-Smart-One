/*
 * File: log_error_panic.go
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
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	log "github.com/sirupsen/logrus"
)

func PanicRecovery(ctx *context.Context, err *error) {
	r := recover()
	if r == nil {
		return
	}
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	stackTrace := string(buf[:n])
	if err != nil {
		*err = fmt.Errorf("%v", r)
	}
	fields := log.Fields{
		"error":   stackTrace,
		"project": os.Getenv("PROJECT"),
	}
	if ctx != nil {
		requestID, hasRequestID := (*ctx).Value("request_id").(string)
		if hasRequestID {
			fields["request_id"] = requestID
		}
	}
	log.WithFields(fields).Panic("panic")
}

func PanicThreadRecovery(ctx *context.Context, err *error, wg *sync.WaitGroup) {
	r := recover()
	if r == nil {
		return
	}
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	stackTrace := string(buf[:n])
	if err != nil {
		*err = fmt.Errorf("%v", r)
	}
	if wg != nil {
		wg.Done()
	}
	fields := log.Fields{
		"error":   stackTrace,
		"project": os.Getenv("PROJECT"),
	}
	if ctx != nil {
		requestID, hasRequestID := (*ctx).Value("request_id").(string)
		if hasRequestID {
			fields["request_id"] = requestID
		}
	}
	log.WithFields(fields).Panic("panic in thread")
}
