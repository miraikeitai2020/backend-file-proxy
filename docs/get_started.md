# Get started

## How to setup backend-file-proxy API
### Launch MinIO
Please prepare the Minio API in your PC. If you can access MinIO from an external network, skip this chapter.  

**Create new access keys**  

In order for you to launch MinIO, you need to use `Access Key` and `Secret Key`. You need two types of keys. In this project, we provided default keys, but we do not recommend using these keys if you are considering long-term use of MinIO.  

|key|value|
|:-:|:-:|
|Access Key|Cg2g6f63KGvzm2a623UEGdiPKYTe66Nb|
|Secret Key|Mf67pN3LsJRabd8j97pk7nxGLq7B3mD8|

If you want to use a new key, you can use the `make` or `go` command to create a new key.

```
$ make
or
$ go run scripts/keys/keyGen.go
```

**Launch MinIO container**  
Launch MinIO container using the `make` or command

```
$ make minio-run PUB_KEY="[ Access Key ]" SEC_KEY="[ Secret Key ]"
```

Once the MinIO container is up, check if you can access MinIO from your web browser.  

Access URL:  
- [http://localhost:9000](http://localhost:9000)

**Port forwarding local MinIO container**  
Use [ngrok](https://ngrok.com/) to make the MinIO container running on your PC accessible from the global network. If you use this service entirely within your local environment, skip this step.

```
$ ngrok http 9000
```
Please keep the URL returned by the above command.
### Launch backend-file-proxy
Invoke the backend-file-proxy API using the `make` or` docker` command. `make` provides a command to build the container, so it is recommended to use it.  

```
$ make docker
```

### Setup backend-file-proxy
backend-file-proxy can dynamically change the target MinIO. In order to access MinIO from backend-file-proxy, it is necessary to update the settings of backend-file-proxy from [Postman](https://www.postman.com/) etc.

|pram|value|
|:-:|:-:|
|Path|/minio/config/update|
|HTTP Method|PUT|

**JSON body example**  
```json
{
    "url": "hogefugafoo.ngrok.io",
    "publicKey": "Access Key",
    "secretKey": "Secret Key"
}
```
※`url`: MinIO URL excluding the scheme

Make sure you get the **same JSON** as the value you entered.  

### Provisioning MinIO
provisioning is done so that backend-file-proxy can use MinIO. This can be achieved by sending the following request to backend-file-proxy.

|pram|value|
|:-:|:-:|
|Path|/minio/init|
|HTTP Method|PUT|
|HTTP Header|key:Access-Key<br>key:Secret-Key|

Check [config.json](../config/bucket.json) to see which bucket will be created.

### [Appendix] POST object from PC
You can use the `make` command to push JPG format objects from your PC.  

```
$ make test
```

To achieve this, you need an environment of `golang v1.15.3`.  

### Access MinIO objects
You can access MinIO objects from your browser.

|pram|value|
|:-:|:-:|
|Path|/image/detour/read/`objectID`|
|HTTP Method|GET|
