## Install sonarqube
```
cd ~/app
wget https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-4.7.0.2747-linux.zip  
unzip sonar-scanner-cli-4.7.0.2747-linux.zip
rm sonar-scanner-cli-4.7.0.2747-linux.zip
```
## Setup environment variable
```
nano ~/.profile

#add at the end this environment variable
 
export PATH="$HOME/app/sonar-scanner-4.7.0.2747-linux/bin:$PATH"
```
## Update environment variable

```
source ~/.profile
```

## Run sonar-scanner

```
sonar-scanner
```

## Run sonar-scanner for a specific folder

```
sonar-scanner -Dsonar.sources=./src
```

## Config to sonar-project.properties

    ```
    sonar-scanner \
    -Dsonar.projectKey=mac-salud-app-2 \
    -Dsonar.sources=. \
    -Dsonar.host.url=http://192.168.71.200:9006 \
    -Dsonar.login=sqp_cfd506762d6464a20ccd31c25845dbcd4d9367be
    ```
