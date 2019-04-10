**_Tech bits_**

- [] Create a REST API in Go, using mux / gorilla?, with hardcoded data
- [] Run this locally within Minikube
- [] Add SQL database

**_User stories_**

- [] As a user, i need to add a 'Trigger' to the app
- [] As a user, i need to a add a 'Response' to the app, related to one or more triggers.

**_How to run this project_**

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
