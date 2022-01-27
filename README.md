# gRPC-Gin
Practice for building a microservices architecture

![Architecture](https://imgur.com/kFLQsPil.png)

user-api-gateway:  
&nbsp; handle the request with route:/v1/user/*  
&nbsp; pass the data to the responsible gRPC servers

post-api-gateway:  
&nbsp; handle the request with route:/v1/post/*  
&nbsp; pass the data to the responsible gRPC servers  
&nbsp;  
&nbsp;  

user-register-server:  
&nbsp; process request relate to register

user-auth-server:  
&nbsp; process request relate to login/ logout/ JWT Validation

post-post-server:  
&nbsp; process request relate to post insertion/ deletion

post-get-server:  
&nbsp; process request relate to post selection
