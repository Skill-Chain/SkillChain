#### Auth

##### Registr
* registerRequest

```http request
POST {host}/auth/register
```

```json
{
    "firstname": "Marie",
    "lastname": "Elman",
    "login": "Cocojamba12",
    "email": "MarieEl09@mail.ru",
    "date": "10-02-1999",
    "phone": "8(961)938-30-20",
    "password": "MeLdhgorfjj221...Gd"
}
```
* registerResponse
```http request
200, OK
```

```json
{
  "id":"fceacd7f03569af0",
  "firstname": "Marie",
  "Lastname": "Elman",
  "login": "Cocojamba12",
  "token": "C73E4154D83B7D11FA2613914F5556B9B31192C987153D4BADB5F6D416CD3BDE"
}
```
##### Auth
* loginRequest

```http request
POST {host}/auth/login
```

```json
{
    "email": "MarieEl09@mail.ru",
    "token": "C73E4154D83B7D11FA2613914F5556B9B31192C987153D4BADB5F6D416CD3BDE"
}
```

* loginResponse
```http response
200, OK
```
```json
{
  "id":"fceacd7f03569af0",
  "firstname": "Marie",
  "lastname": "Elman",
  "login": "Cocojamba12",
  "token": "C73E4154D83B7D11FA2613914F5556B9B31192C987153D4BADB5F6D416CD3BDE"
}
```