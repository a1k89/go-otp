## Go-OTP
#### Go otp authentication (gRPC and RESTful)

1. Redis as database
2. Go
3. SMS provider (Megafon as example)
4. gRPC and RESTful support

#### Idea:
**First step**
1. User send `phonenumber` (ex.79212345555)
2. Generate random code
3. Save `phonenumber:otp` in `Redis`
4. Generate random `token`
5. Save `token:phonenumber` in `Redis`
6. Send back `token`

**Last step:**
1. Send `token + otp` (got from first step)
2. Fetch otp from redis: `token -> phonenumber -> code`
3. Compare otp
4. Send result to user (success:true/false)

#### Installation:
`go get https://github.com/a1k89/go-otp`

#### Or use docker-compose
1. `docker-compose build`
2. `docker-compose up -d`

#### Environments (example)

1. REDIS_HOST=localhost:6379
2. TRANSPORT_CRED_LOGIN
3. TRANSPORT_CRED_PASSWORD
4. TRANSPORT_CRED_FROM=79211234567
5. TRANSPORT_CRED_URL
6. OTP=1111 # When debug. Empty (or nil) in production

#### How to use 
#### RESTful
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
2. Generate gRPC server (for your language)
3. Use code from step2, make logic to get token and verify it

