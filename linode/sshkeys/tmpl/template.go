package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/linode/acceptance"
)

type TemplateData struct {
	Label  string
	SSHKey string
}

func DataBasic(t *testing.T, label, ssh_key string) string {
	return acceptance.ExecuteTemplate(t,
		"ssh_keys_basic", TemplateData{
			Label:  label,
			SSHKey: ssh_key,
		})
}

func DataFilter(t *testing.T, label, ssh_key string) string {
	return acceptance.ExecuteTemplate(t,
		"ssh_keys_filter", TemplateData{
			Label:  label,
			SSHKey: ssh_key,
		})
}

func DataFilterEmpty(t *testing.T, label, ssh_key string) string {
	return acceptance.ExecuteTemplate(t,
		"ssh_keys_filter_empty", TemplateData{
			Label:  label,
			SSHKey: ssh_key,
		})
}

func DataAll(t *testing.T, label, ssh_key string) string {
	return acceptance.ExecuteTemplate(t,
		"ssh_keys_all", TemplateData{
			Label:  label,
			SSHKey: ssh_key,
		})
}
