/*
 * ASR documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloud_asr

type AudioFileDto struct {

	// Binary audio file (Wav format) data as Base64 string
	Data []string `json:"data"`

	// Audio file mime type
	Mime string `json:"mime,omitempty"`
}
