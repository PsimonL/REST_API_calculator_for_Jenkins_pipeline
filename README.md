# REST API written in golang for needs of university classes.
Very simple project was needed within the course of DevOps concepts: https://github.com/InzynieriaOprogramowaniaAGH/MDO2023_INO.  
Aim was to build Dockerfiles for building, testing and deploying stages. Next step was to write Jenkins pipeline in Groovy 
to automatize whole process of CI/CD concept. 
<br>
<br>
Image with working project was published on DockerHub: https://hub.docker.com/r/srpl/go-deploy-img/tags
<br>
<br>
Example usage:  
curl -X POST -H "Content-Type: application/json" -d "[5, 1]" http://localhost:3001/add  
curl http://localhost:3001/hello
<br>
<br>
You could also use for instance Postman and send body as RAW with parameters: [5, 1].  

## "infra" dir
Bunch of files used for deployment of application; dockerfiles for building separate stages in pipeline written in groovy.
Also kubernetes simple deployment for local cluster minikube.
