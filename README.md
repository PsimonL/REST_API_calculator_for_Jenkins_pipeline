# REST API written in golang for needs of university classes. AIM was to build Dockerfiles for building and testing, then writing Jenkins pipelines in Groovy to automatize whole process.

Example usage:  
curl -X POST -H "Content-Type: application/json" -d "[5, 1]" http://localhost:3001/add  
curl http://localhost:3001/hello  

You could also use postman and send body as RAW with parameters: [5, 1].  