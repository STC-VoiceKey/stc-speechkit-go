/*
 * TTS documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 3.1.73
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloud_tts

type SynthesizeText struct {

	// Type of content
	Mime string `json:"mime"`

	// Text for synthesize
	Value string `json:"value"`
}