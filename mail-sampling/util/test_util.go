package util

import (
	"embed"
	"fmt"
	"text/template"
)

//go:embed template/data_connector_execute_fail_mail.html
var tmplFS embed.FS

func CallUtil(templatePath string) {
	fmt.Println("CallUtil", templatePath)

	tmplData, err := tmplFS.ReadFile(templatePath)
	if err != nil {
		fmt.Println("embed.FS: ", err)
	}

	fmt.Println("tmplData: ", tmplData, string(tmplData))

	tmpl, err := template.New("html_template").Parse(string(tmplData))
	if err != nil {
		fmt.Println("template.Parse: ", err)
	}

	fmt.Println("tmpl: ", tmpl)

}
