/*
 * ASR documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloud_asr

type StreamRequestDto struct {

	// Recognize with package
	PackageId string `json:"package_id,omitempty"`

	// Audio file mime type
	Mime string `json:"mime"`
}
