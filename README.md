## Go-OTP
#### Go otp authentication (gRPC and RESTful)

1. Redis
2. Go 3
3. gRPC 
4. REST api

#### Idea:
**First step**
1. Send `phone_number` (ex.79212345555)
2. Generate random code
3. Save `phone_number:otp` in `Redis`
4. Generate random `token`
5. Save `token:phone_number` in `Redis`
6. Send back `token`

**Last step:**
1. Send `token + otp` (got from first step)
2. Fetch otp from redis: `token -> phone_number -> code`
3. Compare otp
4. Send result to user (success:true/false)

#### Installation:
`go get https://github.com/a1k89/go-otp`

#### Or use docker-compose
1. `OTP_PORT=8080 OTP_GRPC_PORT=3000 docker-compose build`
2. `OTP_PORT=8080 OTP_GRPC_PORT=3000 docker-compose up -d`

#### Environments (example)

1. REDIS_HOST=localhost:6379
2. TRANSPORT_CRED_LOGIN
3. TRANSPORT_CRED_PASSWORD
4. TRANSPORT_CRED_FROM=79211234567
5. TRANSPORT_CRED_URL
6. OTP=1111 # Only while debug

#### How to use 
#### REST api
1. First step:
```GO
Method: POST
URL: `/generate/`
Payload: {"phone_number":"<phone_number>"}
Response: {"token": "bla-bla-token"}
```
2. Second step:
```GO
Method: POST
URL: `/verificate/` 
Payload: {"token":"bla-bla-token", "otp": "1111"}
Response: {"status": true/false, "message":"success message"}
```
#### gRPC
1. `proto/otp.proto`
```
syntax = "proto3";
package proto;

option go_package = "github.com/monkrus/grpc-from0;grpc_from0";

service Payload{
  rpc Generate (PayloadGenerateRequest) returns (PayloadGenerateResponse){}
  rpc Verificate (PayloadVerificateRequest) returns (PayloadVerificateResponse){}
}

message PayloadGenerateRequest {
  string phone_number = 1;
}

message PayloadGenerateResponse {
  string token = 1;
}

message PayloadVerificateRequest {
  string token = 1;
  string otp = 2;
}

message PayloadVerificateResponse {
  bool success = 1;
  string message = 2;
}
```
2. Generate gRPC client (python/java/go and etc)
3. Write gRPC client logic (using step2 above)

