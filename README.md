# Api-Wrapper

This repository include api-wrapper source code and write by golang.

## Getting Started
These instructions will get you to know basically `api-wrapper` requirement in your docker.
1. Golang version 1.10.3
2. Tool `dep` should be included
3. All source must put it inside `gateway` (Include: main.go)
4. Dep should be used to control all third party and set ignore in the repo

### Usage
```
git clone --recurse-submodules https://github.com/TinkLabs/api-wrapper
docker exec -it {name} /bin/bash
go version
cd gateway
dep ensure
```

### Singleton Testing
```
docker run -itd --name web -e ENV=ldev -v /Users/$USER/{directory}/api-wrapper:/var/src -p 8080:80 bananabb/ubuntu-go
```

### Push to AWS
```
$(aws ecr get-login --no-include-email --region ap-southeast-1)
docker build -t api-wrapper:latest .
docker tag api-wrapper:latest 204328232493.dkr.ecr.ap-southeast-1.amazonaws.com/api-wrapper:latest
docker push 204328232493.dkr.ecr.ap-southeast-1.amazonaws.com/api-wrapper:latest
MANIFEST=$(aws ecr batch-get-image --repository-name api-wrapper --image-ids imageTag=latest --query images[].imageManifest --output text)
aws ecr put-image --repository-name api-wrapper --image-tag {Please Change Version 1.1.22} --image-manifest "$MANIFEST"
```

### Dep Usage
Please visit as below:
* [Dep](https://github.com/golang/dep)

## License
This project is licensed under the MIT License