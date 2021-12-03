## Go-OTP
#### Go realization to otp authentication

#### Stack:
1. Redis (Save OTP and token)
2. Go
3. SMS provider (Megafon as example)

#### Schema:
1. User send `phonenumber` (ex.79212345555)
2. `Go-OTP` validate it
3. `Go-OTP` generate random code (4-6 digits)
4. Save `phonenumber:otp` in `Redis`
5. Generate random `token`
6. Save `token:phonenumber` in `Redis`
7. Send back `token` to user and send `async` sms code (goroutine)
8. User send `token + otp`
9. Get (Redis): `token -> phonenumber -> code`
10. Compare both `otps`
11. Send result to user

#### Installation:
`go get https://github.com/a1k89/go-otp`

#### Docker-compose
1. Git clone 
2. `docker-compose build`
3. `docker-compose up -d`

#### Environments (example)

1. REDIS_HOST=localhost:6379
2. TRANSPORT_CRED_LOGIN
3. TRANSPORT_CRED_PASSWORD
4. TRANSPORT_CRED_FROM=79211234567
5. TRANSPORT_CRED_URL
6. OTP=1111 # When debug. Empty (or nil) in production

#### How to use
1. First step:
```GO
Method: POST
URL: `/generate/`
Payload: {"phone_number":"<phone_number>"}
Response: {"token": "<TOKEN>"}
```
2. Second step:
```GO
Method: POST
URL: `/verificate/` 
Payload: {"token":"<TOKEN>", "otp": "<OTP_FROM_SMS>"}
Response: {"status": true/false, "message":"success message"}
```
