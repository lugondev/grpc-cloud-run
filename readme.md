```shell
gcloud api-gateway api-configs create grpc-config-1 \
--api=grpc-test --project=bamdev-apt \
--grpc-files=descriptor.pb,api_config.yaml   
#WARNING: Proto descriptor's source protos [ping.proto, service.proto] were not found on the file system and will not be included in the submitted GRPC service definition. If you meant to include these files, ensure the proto compiler was invoked in the same directory where the proto descriptor [temp.pb] now resides.
#Waiting for API Config [grpc-config] to be created for API [grpc-test]...done.
```

```shell
gcloud api-gateway gateways create waas-grpc \
  --api=grpc-test --api-config=grpc-config-1 \
  --location=us-central1 --project=bamdev-apt
```
