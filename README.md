# Test-app



## Getting started
test assignment for devops

## Installation
For docker-compose:
    git clone https://github.com/wdemiurg/test_app.git
    cd test_app 
    docker-compose up -d
    
    curl localhost:8080 - main app  
    localhost:8080/metrics - metrics
    localhost:8080/ready   - readness probe
    localhost:8080/health - health

For kubernetes one yml for all containers and services:
    kubectl apply -f kubernetes/deploy.yaml
## Usage

localhost/ready check connect to site from environment  link 
