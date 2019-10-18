# golang-examples/hello-aws

```sh
$ make build

# Show configuration
$ build/hello-aws
region: us-east-1
config: ~/.aws/credentials
profile: default

# Override configuration with options
$ build/hello-aws --region ap-northeast-1 \
    --config /path/to/.aws/credentials \
    --profile developers
region: ap-northeast-1
config: /path/to/.aws/credentials
profile: developers

# List all readable s3 buckets
$ build/hello-aws --region ap-northeast-1 s3 buckets
...

$ make clean
```
