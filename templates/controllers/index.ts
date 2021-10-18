{{range .Actions -}}
export { {{.}} } from "./{{.}}";
{{end}}