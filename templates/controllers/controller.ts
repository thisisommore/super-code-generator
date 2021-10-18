import { RequestHandler } from "express";

{{if .Body -}}
interface {{.NameTitle}}ReqBody {

}
{{end}}
export const {{.Name}}:RequestHandler=async (req,res)=>{
{{if .Body -}}
    const body = req.body as {{.NameTitle}}ReqBody
{{end -}}
}