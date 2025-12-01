package petproj

import (
	"log"
	"net/http"

	"html/template"

	"github.com/spf13/cobra"
)

const docTemplateStandalone string = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <meta name="description" content="SwaggerUI" />
    <title>SwaggerUI</title>
    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
  </head>
  <body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
  <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-standalone-preset.js" crossorigin></script>
  <script>
    window.onload = () => {
      window.ui = SwaggerUIBundle({
        url: 'swagger.json',
        dom_id: '#swagger-ui',
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        layout: "StandaloneLayout",
      });
    };
  </script>
  </body>
</html>
`

var RunServer = &cobra.Command{
	Use:   "petproj",
	Short: "Run petproj",
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()

		t := template.Must(template.New("docs").Parse(docTemplateStandalone))

		mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			t.Execute(w, nil)
		})
		mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./petproj/gen/pkg/v1/api.swagger.json")
		})
		s := &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	},
}
