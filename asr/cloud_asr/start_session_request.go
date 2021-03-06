/*
 * ASR documentation
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloud_asr

type StartSessionRequest struct {

	// User login
	Username string `json:"username"`

	// User password
	Password string `json:"password"`

	// User domain identifier
	DomainId int64 `json:"domain_id"`
}
