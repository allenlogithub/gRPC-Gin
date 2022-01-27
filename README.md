# gRPC-Gin
Practice for building a microservices architecture

![Architecture](https://imgur.com/Qe7UGvEl.png)

user-api-gateway:  
&nbsp; handles the request under routes:/v1/user/*  
&nbsp; passes the data to the responsible gRPC servers

post-api-gateway:  
&nbsp; handles the request under routes:/v1/post/*  
&nbsp; passes the data to the responsible gRPC servers  
&nbsp;  
&nbsp;  

user-register-server:  
&nbsp; processes the request related to register action

user-auth-server:  
&nbsp; processes the request related to login/ logout/ JWT Validation action

post-post-server:  
&nbsp; processes the request related to post insertion/ deletion action

post-get-server:  
&nbsp; processes the request related to post selection action

