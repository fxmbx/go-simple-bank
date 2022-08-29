## EKS- Amazons Elastice Kubernetes Cluster

after creating an Elatice Container Registry and pushing your docker image there

created a kubernestes cluster to house the nodes we would be running our container on,
created a node group (named it same as project)

after creating the node group, installed kubectl locally
verified its configuration with

# $ kubectl cluster-info

updated the kubeclt config to work with the cluster created on aws

# $ aws eks update-kubeconfig --name <cluter-name> -region <region>

granted the user access to eks access creating a new ploicy and adding it to the user for this project

# $ ls -l ~/.kube

# $ cat ~/.kube/config

# $kubectl config use-context <cluster-arn>

<!-make sure the user is authourizeed to configure cluster ->
kubectl cluster-info

# $aws sts get-caller-identity //checks current user identiy

# $vi ~/.aws/credentials

[dault-profile]
ACCESS_KEY = newly generated access key that has access
ACESS_SECRET = newly generated secret

# $ kubectl cluster-info //checks the cluster info

to enable users have access usr the aws-auth.yaml file then run

# $ kubectl apply -f eks/aws-auth.yaml

then

# $export AWS_PROFILE = [profile-name]

and check the cluster-info using the newly added user

# $ kubectl cluster-info

install k9s for better interface
:namespace
:node

## to enable deployment to the cluster

follow the deployment.yaml file

# $ kubectl apply -f eks/deployment.eks

on k9s
:namespace //to check namespace
:deployment //to see lists of deployment
:pod //to see pods
//hot enter on a specific container when on the pods list
:service //to see service

## to allow trafic from the outside node to the pods we deploy kubernetes service

follow the service.yml file

## i order to expose a service to the outside

we specify the type of service in the .yml file

## ingress

allows set up a record only once but have several rules in the config to route traffic to services in case of multiple services
is also handle load balancing and setting up http is easier
