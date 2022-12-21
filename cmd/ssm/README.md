### How to pull SSM var

There are times when you want to pull all SSM variable to run the code locally. You may have to export all of them OR
set it in some IDE like IntelliJ. To get this you can use this utility whihc can dump the SSM params as export or
key/value.

You can use this utility to pull SSM parameters. We dump these in 2 format:

1. Dump which can be used as export

```shell
>> go run cmd/ssm/dump_ssm_params.go --path=<Any SSM path> --print=export
export MY_PASSWORD=something
```

2. Dump which can be used as normal key/value pare

```shell
>> go run cmd/ssm/dump_ssm_params.go --path=<Any SSM path> --print=na
MY_PASSWORD=something
```