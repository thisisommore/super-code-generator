import { Router } from "express";
import { {{range .Routes -}}{{.}}, {{end}}} from "../controllers/{{.Name}}-controllers";
{{if .Validate}}
import { body } from "express-validator";
{{end}}
const {{.Name}}Router = Router();
{{$name := .Name -}}
{{range .Routes -}}
{{$name}}Router.post("/{{.}}", {{.}});
{{end}}
export default {{.Name}};