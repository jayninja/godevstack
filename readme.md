# godevstack  
  
This is a skeleton im putting together to deploy my golang stuff.  
  
Todo:
- terraform to stand up ec2 instance to build image and push to ECR 
   - userdata.txt needs to get into the userdata of that instance to prep with docker  
   - Git clone dev/prod repo  
   - build repo  
   - tag image  
   - push image to ECR  
- Update ECS  
  
Idea:  
- TF -> spins up instance with user data to prep it for being a build machine.  
- Our build instance pulls code from github and builds the container image with our compiled code.  
- Our build machine pushes the container gets pushed to ECR.  
- Figure out how to deploy the image from ECR into something that will expose it to the internet.  
  
  
Current:  
- build.sh will generate the go.mod and go.sum in the source directory.  
-          it will then start the container build  

