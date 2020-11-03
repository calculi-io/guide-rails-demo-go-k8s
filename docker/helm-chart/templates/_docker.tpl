
{{/* Define a secret for the docker registry */}}
{{- define "imagePullSecret" }}
{{- printf "{\"insecure-registries\" : [\"%s:%d\"], \"auths\":{\"%s:%d\": {\"username\":\"%s\",\"password\":\"%s\",\"email\":\"%s\"}}}" .Values.imageCredentials.registry .Values.imageCredentials.registryPort .Values.imageCredentials.registry .Values.imageCredentials.registryPort .Values.imageCredentials.username .Values.imageCredentials.password .Values.imageCredentials.email | b64enc }}
{{- end }}