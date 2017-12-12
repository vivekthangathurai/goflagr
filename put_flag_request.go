/*
 * Flagr
 *
 * Flagr is a feature flagging, A/B testing and dynamic configuration microservice
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goflagr

type PutFlagRequest struct {

	Description string `json:"description"`

	// enabled data records will get data logging in the metrics pipeline, for example, kafka.
	DataRecordsEnabled bool `json:"dataRecordsEnabled,omitempty"`
}