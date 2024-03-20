package {{.PkgName}}

import (
	"net/http"

	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	http2 "github.com/scutrobotlab/goctl-template/src/http"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			http2.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			{{if .HasResp}}http2.JsonBaseResponseCtx(r.Context(), w, resp){{else}}http2.JsonBaseResponseCtx(r.Context(), w, nil){{end}}
		}
	}
}
