# FastVault
FastHttp Vault via HTTP REST

### API
- Write to FastVault

Write value to FastVault. the response must return up to 3 keys 
with should be in specific headers.
```
POST /secret 
```
Request Body
```
HEADER:
- X-Application-Key: []key
BODY:
[your data]
```
Response
```json
{
  "key1": "qwertyuiop",
  "key2": "asdfghjkl;",
  "key3": "zxcvbnm,./"
}
```

