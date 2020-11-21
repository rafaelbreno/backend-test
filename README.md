# code-rockr-test
## Deploying app
- This app is 100% running on Docker, so it's easier to locally deploy it with:
- > $ sudo ./start.sh

## Problem
- API for social events
- More at [Coderockr](https://github.com/Coderockr/backend-test)

## Solution Description
- I'm studying Go latelly, so I'll be using
    - [GORM](https://gorm.io/) as my ORM
    - [Gonic-Gin](https://github.com/gin-gonic/gin) for routing
    - [godotenv](https://github.com/joho/godotenv) to not mess with _.env_ files reading
    - [jwt-go](https://github.com/dgrijalva/jwt-go) for jwt tokens

## Endpoints
### _Public Area_
- _/_
    - HTTP Method: _GET_
    - Request:
        - ```go
            {
                "filter":
            } 
          ```
    - Response:
        - ```go
            {
                "events": [
                    [
                    {
                        "name": "Party",
                        "date": "23/11/2019"
                    },
                    ]
                ]
            }
          ```
    - The events object is an array of arrays with 10 items for paginations
- _/:event-slug_
    - HTTP Method: _GET_
    - Request: ```{}```
    - Response:
    - ```json
{
    "data": {
        "event": {
            [..details..]
        }
    }
}
      ```
- _/login_
    - HTTP: _POST_
    - Request:
    -   ```json
            {
                "email": "john.doe@mock.com",
                "password": "*******"
            }
        ```
    - Response:
    - ```json
        {
            "token": "Bearer [...]"
        }
      ```
- _/register_
    - HTTP: _POST_
    - Request:
    -   ```json
            {
                "email": "john.doe@mock.com",
                "password_confirmation": "*******",
                "password": "*******",
                "biography": "......",
                "profile_picture": "/images/{{random_string}}",
                "city": "....",
                "state": "....."
            }
        ```
    - Response:
    - ```
        {
            "token": "Bearer [...]"
        }
      ```
