/*
 * Diarization documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloud_diarization

// Object for exceptions/messages wrapping
type MessageDto struct {

	// Reason
	Reason string `json:"reason,omitempty"`

	// Message
	Message string `json:"message,omitempty"`
}