/*
 * File: swagger_json.go
 * Author: edward
 * Copyright: 2024, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the json template documentation for the microservice.
 *
 * Last Modified: 2024-04-09
 */

package docs

import (
	_ "embed"
)

//go:embed swagger3.json
var DocTemplateJson string
