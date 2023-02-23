D:\dev2023\go_notify>docker build . 
[+] Building 31.3s (12/14)
 => [internal] load build definition from Dockerfile                                                                                                                                                                                                        0.0s 
 => => transferring dockerfile: 349B                                                                                                                                                                                                                        0.0s 
 => [internal] load .dockerignore                                                                                                                                                                                                                           0.0s 
 => => transferring context: 2B                                                                                                                                                                                                                             0.0s 
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                                                                                                            3.4s 
 => [internal] load metadata for docker.io/library/golang:1.17                                                                                                                                                                                              3.5s 
 => [stage-1 1/5] FROM docker.io/library/alpine:latest@sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a                                                                                                                             18.7s 
 => => resolve docker.io/library/alpine:latest@sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a                                                                                                                                      0.0s 
 => => sha256:e2e16842c9b54d985bf1ef9242a313f36b856181f188de21313820e177002501 528B / 528B                                                                                                                                                                  0.0s 
 => => sha256:b2aa39c304c27b96c1fef0c06bee651ac9241d49c4fe34381cab8453f9a89c7d 1.47kB / 1.47kB                                                                                                                                                              0.0s 
 => => sha256:69665d02cb32192e52e07644d76bc6f25abeb5410edc1c7a81a10ba3f0efb90a 1.64kB / 1.64kB                                                                                                                                                              0.0s 
 => => sha256:63b65145d645c1250c391b2d16ebe53b3747c295ca8ba2fcb6b0cf064a4dc21c 3.37MB / 3.37MB                                                                                                                                                             18.5s 
 => => extracting sha256:63b65145d645c1250c391b2d16ebe53b3747c295ca8ba2fcb6b0cf064a4dc21c                                                                                                                                                                   0.1s 
 => [internal] load build context                                                                                                                                                                                                                           0.2s 
 => => transferring context: 33.03kB                                                                                                                                                                                                                        0.2s 
 => [build 1/4] FROM docker.io/library/golang:1.17@sha256:87262e4a4c7db56158a80a18fefdc4fee5accc41b59cde821e691d05541bbb18                                                                                                                                 26.9s 
 => => resolve docker.io/library/golang:1.17@sha256:87262e4a4c7db56158a80a18fefdc4fee5accc41b59cde821e691d05541bbb18                                                                                                                                        0.0s 
 => => sha256:87262e4a4c7db56158a80a18fefdc4fee5accc41b59cde821e691d05541bbb18 2.35kB / 2.35kB                                                                                                                                                              0.0s 
 => => sha256:742df529b073e7d1e213702a6cca40c32f3f5068125988de249416ba0abee517 7.12kB / 7.12kB                                                                                                                                                              0.0s 
 => => sha256:d836772a1c1f9c4b1f280fb2a98ace30a4c4c87370f89aa092b35dfd9556278a 55.00MB / 55.00MB                                                                                                                                                            5.5s 
 => => sha256:55636cf1983628109e569690596b85077f45aca810a77904e8afad48b49aa500 1.80kB / 1.80kB                                                                                                                                                              0.0s 
 => => sha256:66a9e63c657ad881997f5165c0826be395bfc064415876b9fbaae74bcb5dc721 5.16MB / 5.16MB                                                                                                                                                              3.9s 
 => => sha256:d1989b6e74cfdda1591b9dd23be47c5caeb002b7a151379361ec0c3f0e6d0e52 10.88MB / 10.88MB                                                                                                                                                            4.9s 
 => => sha256:c28818711e1ed38df107014a20127b41491b224d7aed8aa7066b55552d9600d2 54.58MB / 54.58MB                                                                                                                                                           12.7s 
 => => sha256:9d6246ba248cc80872dc2995f9080ef76305b540968dadb096b75f2e2146a38a 85.90MB / 85.90MB                                                                                                                                                           19.8s 
 => => sha256:21d43f0d73c2979514706af3d892f631b75d5c2d56aebfac0172e5a4e934b447 135.06MB / 135.06MB                                                                                                                                                         24.3s 
 => => extracting sha256:d836772a1c1f9c4b1f280fb2a98ace30a4c4c87370f89aa092b35dfd9556278a                                                                                                                                                                   1.1s 
 => => extracting sha256:66a9e63c657ad881997f5165c0826be395bfc064415876b9fbaae74bcb5dc721                                                                                                                                                                   0.1s 
 => => extracting sha256:d1989b6e74cfdda1591b9dd23be47c5caeb002b7a151379361ec0c3f0e6d0e52                                                                                                                                                                   0.2s 
 => => sha256:d8a1c5873f408d3f5a8d8d73c6b9a3d77818bab0b26142a493909ea8c4d0c020 154B / 154B                                                                                                                                                                 14.6s 
 => => extracting sha256:c28818711e1ed38df107014a20127b41491b224d7aed8aa7066b55552d9600d2                                                                                                                                                                   1.1s 
 => => extracting sha256:9d6246ba248cc80872dc2995f9080ef76305b540968dadb096b75f2e2146a38a                                                                                                                                                                   1.4s 
 => => extracting sha256:21d43f0d73c2979514706af3d892f631b75d5c2d56aebfac0172e5a4e934b447                                                                                                                                                                   2.3s 
 => => extracting sha256:d8a1c5873f408d3f5a8d8d73c6b9a3d77818bab0b26142a493909ea8c4d0c020                                                                                                                                                                   0.0s 
 => [stage-1 2/5] RUN apk --no-cache add ca-certificates                                                                                                                                                                                                    4.9s 
 => [stage-1 3/5] WORKDIR /app                                                                                                                                                                                                                              0.0s 
 => [build 2/4] WORKDIR /app                                                                                                                                                                                                                                0.5s 
 => [build 3/4] COPY . .                                                                                                                                                                                                                                    0.0s 
 => ERROR [build 4/4] RUN go mod download &&     CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GoNotify .                                                                                                                                      0.4s 
------
 > [build 4/4] RUN go mod download &&     CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GoNotify .:
#12 0.363 go mod download: no modules specified (see 'go help mod download')
------
executor failed running [/bin/sh -c go mod download &&     CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o GoNotify .]: exit code: 1

D:\dev2023\go_notify>go mod init
go: cannot determine module path for source directory D:\dev2023\go_notify (outside GOPATH, module path must be specified)

Example usage:
        'go mod init example.com/m' to initialize a v0 or v1 module
        'go mod init example.com/m/v2' to initialize a v2 module

Run 'go help mod init' for more information.

D:\dev2023\go_notify>go mod init .
go: malformed module path ".": is a local import path

D:\dev2023\go_notify>go mod tidy
go: go.mod file not found in current directory or any parent directory; see 'go help modules'

D:\dev2023\go_notify>git init   
Initialized empty Git repository in D:/dev2023/go_notify/.git/

D:\dev2023\go_notify>git add .

D:\dev2023\go_notify>git commit -m "first commit"
[master (root-commit) cbb5160] first commit
 20 files changed, 1082 insertions(+)
 create mode 100644 Dockerfile
 create mode 100644 README.md
 create mode 100644 config/config.yaml
 create mode 100644 main.go
 create mode 100644 platform/platform.go
 create mode 100644 propmt/step1.md
 create mode 100644 propmt/step10.md
 create mode 100644 propmt/step11.md
 create mode 100644 propmt/step2.md
 create mode 100644 propmt/step3.md
 create mode 100644 propmt/step4.md
 create mode 100644 propmt/step5.md
 create mode 100644 propmt/step6.md
 create mode 100644 propmt/step7.md
 create mode 100644 propmt/step8.md
 create mode 100644 propmt/step9.md
 create mode 100644 schedule/schedule.go
 create mode 100644 service/service.go
 create mode 100644 utils/email.go
 create mode 100644 utils/sms.go

D:\dev2023\go_notify>git branch -M main

D:\dev2023\go_notify>git remote add origin git@github.com:aaronpeng2020/gonotify.git

D:\dev2023\go_notify>git push -u origin main
Enumerating objects: 28, done.
Counting objects: 100% (28/28), done.
Delta compression using up to 16 threads
Compressing objects: 100% (24/24), done.
Writing objects: 100% (28/28), 13.77 KiB | 1.53 MiB/s, done.
Total 28 (delta 6), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (6/6), done.
To github.com:aaronpeng2020/gonotify.git
 * [new branch]      main -> main
Branch 'main' set up to track remote branch 'main' from 'origin'.

D:\dev2023\go_notify>go mod init .
go: malformed module path ".": is a local import path

D:\dev2023\go_notify>go mod init github.com/aaronpeng2020/gonotify
go: creating new go.mod: module github.com/aaronpeng2020/gonotify
go: to add module requirements and sums:
        go mod tidy

D:\dev2023\go_notify>go mod tidy
go: finding module for package github.com/spf13/viper
go: finding module for package github.com/aaronpeng2020/gonotify/util
go: finding module for package github.com/robfig/cron/v3
go: downloading github.com/robfig/cron v1.2.0
go: downloading github.com/spf13/viper v1.15.0
go: downloading github.com/robfig/cron/v3 v3.0.1
go: found github.com/robfig/cron/v3 in github.com/robfig/cron/v3 v3.0.1
go: found github.com/spf13/viper in github.com/spf13/viper v1.15.0
go: downloading github.com/spf13/afero v1.9.3
go: downloading github.com/spf13/cast v1.5.0
go: downloading github.com/spf13/jwalterweatherman v1.1.0
go: downloading github.com/fsnotify/fsnotify v1.6.0
go: downloading github.com/mitchellh/mapstructure v1.5.0
go: downloading gopkg.in/ini.v1 v1.67.0
go: downloading github.com/stretchr/testify v1.8.1
go: downloading github.com/hashicorp/hcl v1.0.0
go: downloading github.com/magiconair/properties v1.8.7
go: downloading github.com/subosito/gotenv v1.4.2
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/pelletier/go-toml/v2 v2.0.6
go: downloading golang.org/x/sys v0.3.0
go: downloading golang.org/x/text v0.5.0
go: downloading github.com/frankban/quicktest v1.14.3
go: downloading gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
go: downloading github.com/rogpeppe/go-internal v1.6.1
go: finding module for package github.com/aaronpeng2020/gonotify/util
github.com/aaronpeng2020/gonotify imports
        github.com/aaronpeng2020/gonotify/util: no matching versions for query "latest"

D:\dev2023\go_notify>