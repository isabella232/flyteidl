/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

type AdminSlackNotification struct {
	// Currently, Slack notifications leverage email to trigger a notification.
	RecipientsEmail []string `json:"recipients_email,omitempty"`
}
