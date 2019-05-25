# FastVault
FastHttp Vault via HTTP REST
### Configuration
```
export ENV_CONFIG_LOCATIO=/path/to/your/configuration
```
Configuration parameter
```json
{
  "port": ":8080",
  "key": "0123456789abcdef",
  "secret_storage": "./data"
}
```


### API
##### Write to FastVault

Write value to FastVault. the response will return token to access the value.
```
POST /secret 
```
Request Body:
```
[your data]
```
Response:
```json
{
  "token": "this-is-token"
}
```

---
##### Read from FastVault

Read value from FastVault. By given a token via header.
```
GET /secret
```
Header:
```
X-Application-Token: [token]
```
Response:
```
[your data]
```