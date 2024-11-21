set GOARCH=amd64
set GOOS=linux
set CGO_ENABLED=0

set REGISTRY_URL=microk8s-vm.mshome.net:32000
set DB_HOST=192.168.71.200
set NAMESPACE=smartone-local
set IMAGE=core.smartone:1.0.0
set LABEL=core-smartone

go build -o app .
docker build -t "%REGISTRY_URL%/%IMAGE%" -f "%PROJECT_PATH%Dockerfile" .
docker push "%REGISTRY_URL%/%IMAGE%"
del "app"

powershell -NoProfile -Command "$content = Get-Content -Path cicd\k8s\local\deployment-api.yaml -Raw; $matches = [regex]::Matches($content, '\$\{([^}]+)\}'); foreach ($match in $matches) { $varName = $match.Groups[1].Value; $replacement = (Get-Item -LiteralPath \"Env:$varName\").Value; $content = $content.Replace($match.Value, $replacement); } Set-Content -Path deployment-api.yaml -Value $content;"

set "kubeconfig=%USERPROFILE%\.kube\config-local"
kubectl --kubeconfig="%kubeconfig%" apply -f deployment-api.yaml
del deployment-api.yaml

powershell -NoProfile -Command "$content = Get-Content -Path cicd\k8s\local\service-api.yaml -Raw; $matches = [regex]::Matches($content, '\$\{([^}]+)\}'); foreach ($match in $matches) { $varName = $match.Groups[1].Value; $replacement = (Get-Item -LiteralPath \"Env:$varName\").Value; $content = $content.Replace($match.Value, $replacement); } Set-Content -Path service-api.yaml -Value $content;"

kubectl --kubeconfig="%kubeconfig%" apply -f service-api.yaml
del service-api.yaml

kubectl --kubeconfig "%kubeconfig%" delete pods --selector app="%LABEL%" --namespace "%NAMESPACE%"
