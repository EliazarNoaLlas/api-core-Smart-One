/*
 * File: swagger_json.go
 * Author: bengie
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the json template documentation for the microservice.
 *
 * Last Modified: 2024-03-28
 */

package docs

import (
	_ "embed"
)

//go:embed swagger3.json
var DocTemplateJson string
