<h1 align="center">
<a href="https://blog.horonlee.com">Visual Docker</a>
</h1>

<p align="center">
🐳 Manage Containers with Ease (Beta).
</p>

<pre align="center">
Makes it easier and faster to use docker
🧪 developing
</pre>

- **English** | [简体中文](./README.md)

## TODO:
- [x] You can view docker information through the console.
- [x] The console will terminate the application if it detects a Docker exception.
- [x] Displaying information through the web interface
- [x] Connect to Kubernetes cluster and show all pods through console.
- [ ] Multi-database support(SQLite MySQL)

## How to build

> Required Docker Client API Version >= 1.45

1. Go to the project directory and execute `go build`. 2.
2. Get the `VDController` binary and give it executable permissions `sudo chmod +x VDController`. 3.
3. Put `VDController` into a separate folder and put it into the project's webSrc folder.
4. Execute `. /VDController` to start the application

## Configuration files

> The configuration file is created in the same directory as the application after the first run and can be changed later.

- `WebEnable = true&false` Whether to automatically enable the web function after starting the program
- `ListeningPort = '0.0.0.0:8080'` The listening address and port for the web function
- `KubeEnable = true&false` Whether to automatically enable the Kubernetes function after starting the program
- `KubeconfigPath = '.kube/config file path'` The configuration file path for the Kubernetes function
    - If not specified, the default will be `$HOME/.kube/config`
- `DBType = 'sqlite&mysql'` Database type, defaults to sqlite. Currently, only sqlite and mysql are supported
- `DBPath = 'data.db'` Database file path, defaults to `data.db` in the current directory of the program
- `DBAddr = '127.0.0.1:3306'` Database address
- `DBUser = 'root'` Database username
- `DBPass = 'password'` Database password
- `DBName = 'test'` Database name

Example:
```toml
WebEnable = true
ListeningPort = '127.0.0.1:1024'
KubeEnable = true
KubeconfigPath = '/Users/horonlee/Downloads/k8s/config'
```

## Web usage

1. `IP:8080` is a default homepage (nothing)
2. `IP:8080/json/*` returns a variety of json information.
   1. `IP:8080/json/docker` docker
   2. `IP:8080/json/kube` kubernetes
3. `IP:8080/search?image=$IMAGE_NAME` Returns the running container for the specified image.

## Environment variable

- LOG_DIR Path to the log file `/var/log/vdcontroller`.

## Startup parameters

> Support to configure software settings via startup parameters, e.g.: `./VDController -kubeconfig="/home/user/document/k8s/config"

- `-kubeconfig` Kubernetes configuration file path