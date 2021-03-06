package generator

import (
	"strings"
	"text/template"
)

var functionFuncs template.FuncMap = template.FuncMap{
	"ToLower":  strings.ToLower,
	"UnExport": unexport,
	"Replace":  strings.Replace,
}

const functionTemplate string = `// Code generated by counterfeiter. DO NOT EDIT.
package {{.DestinationPackage}}

import (
	{{- range .Imports}}
	{{.Alias}} "{{.Path}}"
	{{- end}}
)

type {{.Name}} struct {
	Stub func({{.Function.Params.AsArgs}}) {{.Function.Returns.AsReturnSignature}}
	mutex sync.RWMutex
	argsForCall []struct{
		{{- range .Function.Params}}
		{{.Name}} {{if .IsVariadic}}{{Replace .Type "..." "[]" -1}}{{else}}{{.Type}}{{end}}
		{{- end}}
	}
	{{- if .Function.Returns.HasLength}}
	returns struct{
		{{- range .Function.Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}
	returnsOnCall map[int]struct{
		{{- range .Function.Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}
	{{- end}}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *{{.Function.FakeName}}) Spy({{.Function.Params.AsNamedArgsWithTypes}}) {{.Function.Returns.AsReturnSignature}} {
	{{- range .Function.Params.Slices}}
	var {{UnExport .Name}}Copy {{.Type}}
	if {{UnExport .Name}} != nil {
		{{UnExport .Name}}Copy = make({{.Type}}, len({{UnExport .Name}}))
		copy({{UnExport .Name}}Copy, {{UnExport .Name}})
	}
	{{- end}}
	fake.mutex.Lock()
	{{if .Function.Returns.HasLength}}ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	{{end}}fake.argsForCall = append(fake.argsForCall, struct{
		{{- range .Function.Params}}
		{{.Name}} {{if .IsVariadic}}{{Replace .Type "..." "[]" -1}}{{else}}{{.Type}}{{end}}
		{{- end}}
	}{ {{- .Function.Params.AsNamedArgs -}} })
	fake.recordInvocation("{{.TargetName}}", []interface{}{ {{- if .Function.Params.HasLength}}{{.Function.Params.AsNamedArgs}}{{end -}} })
	fake.mutex.Unlock()
	if fake.Stub != nil {
		{{if .Function.Returns.HasLength}}return fake.Stub({{.Function.Params.AsNamedArgsForInvocation}}){{else}}fake.Stub({{.Function.Params.AsNamedArgsForInvocation}}){{end}}
	}
	{{- if .Function.Returns.HasLength}}
	if specificReturn {
		return {{.Function.Returns.WithPrefix "ret."}}
	}
	return {{.Function.Returns.WithPrefix "fake.returns."}}
	{{- end}}
}

func (fake *{{.Function.FakeName}}) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *{{.Function.FakeName}}) Calls(stub func({{.Function.Params.AsArgs}}) {{.Function.Returns.AsReturnSignature}}) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

{{if .Function.Params.HasLength -}}
func (fake *{{.Function.FakeName}}) ArgsForCall(i int) {{.Function.Params.AsReturnSignature}} {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return {{.Function.Params.WithPrefix "fake.argsForCall[i]."}}
}
{{- end}}

{{if .Function.Returns.HasLength -}}
func (fake *{{.Function.FakeName}}) Returns({{.Function.Returns.AsNamedArgsWithTypes}}) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		{{- range .Function.Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}{ {{- .Function.Returns.AsNamedArgs -}} }
}

func (fake *{{.Function.FakeName}}) ReturnsOnCall(i int, {{.Function.Returns.AsNamedArgsWithTypes}}) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			{{- range .Function.Returns}}
			{{UnExport .Name}} {{.Type}}
			{{- end}}
		})
	}
	fake.returnsOnCall[i] = struct {
		{{- range .Function.Returns}}
		{{UnExport .Name}} {{.Type}}
		{{- end}}
	}{ {{- .Function.Returns.AsNamedArgs -}} }
}
{{- end}}

func (fake *{{.Function.FakeName}}) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *{{.Name}}) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ {{.TargetAlias}}.{{.TargetName}} = new({{.Name}}).Spy
`
