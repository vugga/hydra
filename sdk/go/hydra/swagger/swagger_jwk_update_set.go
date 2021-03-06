/*
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type SwaggerJwkUpdateSet struct {
	Body JsonWebKeySet `json:"Body,omitempty"`

	// The set in: path
	Set string `json:"set"`
}
