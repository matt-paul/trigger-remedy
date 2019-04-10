**Tech bits**

- [x] Create a REST API with hardcoded data
- [x] Run this locally within Minikube
- [] Add SQL database

**User stories**

- [] As a user, i need to add a 'Trigger' to the app
- [] As a user, i need to a add a 'Response' to the app, related to one or more triggers.

**How to run this project**

Install Minikube

```
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 && \
chmod +x minikube && \
sudo mv minikube /usr/local/bin/
```

Install the driver

```
brew install docker-machine-driver-xhyve
sudo chown root:wheel $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
sudo chmod u+s $(brew --prefix)/opt/docker-machine-driver-xhyve/bin/docker-machine-driver-xhyve
```

Start the minikube cluster - xhyve specifies that you are using 'Docker for mac' hypervisor
`minikube start --vm-driver=xhyve`

Set kubernetes cli to use our Minikube cluster
`kubectl config use-context minikube`

Verify that kubectl can communicate with the cluster
`kubectl cluster-info`

**Docker**
Build the image, with an appropriate tag on the end
`docker build -t trigger-remedy-app:1.0 .`

Run the image, which will be found at localhost:3030
`docker run -p 3030:3001 -i --rm --name my-golang-app-run trigger-remedy-app:1.0`

`-i` Keep STDIN open even if not attached (--interactive)
`-p` Publish a container's port(s) to the host (--publish list )
`--rm` Automatically remove the container when it exits
`--name` Assign a name to the container

**Kubernetes**

`minikube start`

View the dashboard
`minikube dashboard`

Set docker env to use local docker image(this needs to be set within each terminal)
`eval $(minikube docker-env)`

Run in minikube
`kubectl run trigger-remedy-app --image=trigger-remedy-app:1.0`

Create a kubernetes deployment
`kubectl create deployment trigger-remedy-app --image=trigger-remedy-app:1.0`

Check that it's running
`kubectl get deployments`

View the pods
`kubectl get pods`

View cluster events
`kubectl get events`

View kubectl configuration
`kubectl config view`

Expose the pod to the public internet as a Service
`kubectl expose deployment trigger-remedy-app --type=LoadBalancer --port=3001`

View in the browser bia minikube
`minikube service trigger-remedy-app`
