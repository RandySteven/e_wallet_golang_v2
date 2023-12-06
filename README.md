# assignment-go-rest-api

## Register User A.P.I
- <b>/v1/register</b>
### Register Success
```
Response Code - 201
``` 

`Request Sample`
```json
{
  "name": "Test User",
  "email": "test.user@gmail.com",
  "password": "test_1234"
}
```

`Response Sample`
```json
{
    "message": "Success created user",
    "data": {
        "ID": 3,
        "name": "Test User",
        "email": "test.user@gmail.com"
    }
}
```
![Alt register-success](/img/register_success.png)

### Register Failed - Email Format
```
Response Code - 400
``` 

`Request Sample`
```json
{
  "name": "Matthew Alfredo",
  "email": "matthew.alfredo",
  "password": "test_1234"
}
```

`Response Sample`
```json
{
    "errors": [
        "Email field is email"
    ]
}
```
![Alt register-failed-email](/img/register_failed_email_validation.png)

### Register Failed - Password Required
```
Response Code - 400
``` 

`Request Sample`
```json
{
  "name": "Matthew Alfredo",
  "email": "matthew.alfredo@gmail.com"
}
```

`Response Sample`
```json
{
    "errors": [
        "Password field is required"
    ]
}
```
![Alt register-failed-email](/img/register_failed_password_required.png)

### Register Failed - Email Required
```
Response Code - 400
``` 

`Request Sample`
```json
{
  "name": "Matthew Alfredo",
  "password": "test_1234"
}
```

`Response Sample`
```json
{
    "errors": [
        "Email field is required"
    ]
}
```
![Alt register-failed-email](/img/register_failed_email_required.png)

### Register Failed - Name Required
```
Response Code - 400
``` 

`Request Sample`
```json
{
  "email": "matthew.alfredo@gmail.com",
  "password": "test_1234"
}
```

`Response Sample`
```json
{
    "errors": [
        "Name field is required"
    ]
}
```
![Alt register-failed-email](/img/register_failed_name_required.png)

### Register Failed - Email Already Exists
```
Response Code - 400
``` 

`Request Sample`
```json
{
  "name": "Alvin Fernando",
  "email": "alvin.fernando@gmail.com",
  "password": "test_1234"
}
```

`Response Sample`
```json
{
    "errors": [
        "email already exists"
    ]
}
```
![Alt register-failed-email](/img/register_failed_email_already_exists.png)