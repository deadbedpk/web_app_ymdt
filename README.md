# web_app_ymdt
This is a simple web application is developed in golang. This is display current year, month, date and time. It runs on the port 6543 in localhost.

first install go modules in current operating system.

git clone https://github.com/deadbedpk/web_app_ymdt.git .

cd web_app_ymdt.

go run app.go 
//then open the browser type localhost:6543


-----------------

In docker environment 

1. docker build . -t <name> 
2. docker images // now it will listed 
3. docker run --name <name> <image id or container id > 
  
  
-----------------
  
In k8s environment 
  
1. kubectl create -f https://github.com/deadbedpk/web_app_ymdt/blob/main/finalreplica.yaml 
  // this will create 2 replica and go to browser in url 
  
  https://localhost:30008 // it will visible in browser 

 
  

