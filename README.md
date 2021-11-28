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

#### How to use
1. Generate token:
```GO
POST: `/generate/` {"phone_number":"PHONE_NUMBER"}
```
2. Verification:
```GO
POST `/verificate/` 
Payload: {
	    "token":"<TOKEN>", 
	    "otp": "<OTP_FROM_SMS>"}
Response: {
	"status": true/false,
	"message":"success message"}
```
