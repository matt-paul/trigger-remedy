**Tech bits**

- [] Create a REST API in Go, using mux / gorilla?, with hardcoded data
- [] Run this locally within Minikube
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
`docker build -t trigger-remedy-app .`

`docker run -p 3030:3001 -i --rm --name my-golang-app-run trigger-remedy-app`

`-i` Keep STDIN open even if not attached (--interactive)
`-p` Publish a container's port(s) to the host (--publish list )
`--rm` Automatically remove the container when it exits
`--name` Assign a name to the container

Server will be running on localhost:3030
