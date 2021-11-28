## GoTP
#### Go realization to otp authentication

Stack:
1. Redis (Save OTP and token)
2. Go
3. SMS provider (Megafon)

Schema:
1. User send `phonenumber` (ex.79212345555)
2. `GoTP` validate it
3. Generate random code (4-6 digits)
4. Save `phone_number:otp` in `Redis`
5. Generate random `token`
6. Save `token:phone_number` in `Redis`
7. Send back `token` to user and send `async` sms code (use sms provider)
8. User send `token + sms code`
9. Get (Redis): `token -> phone_number -> code`
10. Compare `user_code` and `Redis code`
11. Send result to user
