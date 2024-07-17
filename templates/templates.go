package templates

var ProjectTemplate = `Id:             {{ .Id }}
Name:           {{ .Name }}
Abbreviation:   {{ .Abbreviation }}
Description:    {{ .Description }}
State:          {{ .State }}
Visibility:     {{ .Visibility }}
LastUpdateTime: {{ .LastUpdateTime }}
Revision:       {{ .Revision }}
Url:            {{ .Url }}
DefaultTeam:
  Id:   {{ .DefaultTeam.Id }}
  Name: {{ .DefaultTeam.Name }}
  Url:  {{ .DefaultTeam.Url }}{{ if .Capabilities }}
Capabilities:
  processTemplate:
    templateTypeId: {{ .Capabilities.processTemplate.templateTypeId }}
    templateName:   {{ .Capabilities.processTemplate.templateName }}
  versioncontrol:
    gitEnabled:        {{ .Capabilities.versioncontrol.gitEnabled }}
    sourceControlType: {{ .Capabilities.versioncontrol.sourceControlType }}
    tfvcEnabled:       {{ .Capabilities.versioncontrol.tfvcEnabled }}{{ else }}
Capabilities: <not requested>{{ end }}
Links:
  collection: {{ .Links.collection.href }}
  self:       {{ .Links.self.href }}
  web:        {{ .Links.web.href }}
`
