# gRPC-Gin
Practice for building a microservices architecture

![Architecture](https://imgur.com/dmnFnkyl.png)

user-api-gateway:  
&nbsp; handles the request under routes:/v1/user/*  
&nbsp; passes the data to the responsible gRPC servers

post-api-gateway:  
&nbsp; handles the request under routes:/v1/post/*  
&nbsp; passes the data to the responsible gRPC servers  
&nbsp;  
&nbsp;  

user-register-server:  
&nbsp; processes the request related to register actions

user-auth-server:  
&nbsp; processes the request related to login/ logout/ JWT Validation actions

user-get-server:  
&nbsp; processes the request related to user selection actions

user-post-server:  
&nbsp; processes the request related to user insertion actions

post-get-server:  
&nbsp; processes the request related to post selection actions

post-post-server:  
&nbsp; processes the request related to post insertion/ deletion actions
