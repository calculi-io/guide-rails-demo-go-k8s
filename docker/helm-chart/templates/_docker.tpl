
{{/* Define a secret for the docker registry */}}
{{- define "imagePullSecret" }}
{{- printf "{\"insecure-registries\" : [\"%s\"], \"auths\":{\"%s\": {\"username\":\"%s\",\"password\":\"%s\",\"email\":\"%s\"}}}" .Values.imageCredentials.registry .Values.imageCredentials.registry .Values.imageCredentials.username .Values.imageCredentials.password .Values.imageCredentials.email | b64enc }}
{{- end }}