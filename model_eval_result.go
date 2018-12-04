/*
 * Flagr
 *
 * Flagr is a feature flagging, A/B testing and dynamic configuration microservice. The base path for all the APIs is \"/api/v1\". 
 *
 * API version: 1.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goflagr

type EvalResult struct {
	FlagID int64 `json:"flagID"`
	FlagKey string `json:"flagKey"`
	FlagSnapshotID int64 `json:"flagSnapshotID,omitempty"`
	SegmentID int64 `json:"segmentID"`
	VariantID int64 `json:"variantID"`
	VariantKey string `json:"variantKey"`
	VariantAttachment *interface{} `json:"variantAttachment"`
	EvalContext *EvalContext `json:"evalContext"`
	Timestamp string `json:"timestamp"`
	EvalDebugLog *EvalDebugLog `json:"evalDebugLog,omitempty"`
}