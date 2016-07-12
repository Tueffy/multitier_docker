Multitier Docker Example
------------
This application stack is comprised out of three images
* HA Proxy
* Web Server written in Go (which increases a counter in Redis for every hit)
* Redis KV store

Required Installation Process
------------
1. Install Docker v1.12 experimental
~~~
curl -sSL https://experimental.docker.com/ | sh
~~~
2. Install Docker Compose
~~~
curl -L https://github.com/docker/compose/releases/download/1.8.0-rc2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
~~~
3. Start Docker Service
~~~
sudo service docker start
~~~
4. Ensure that your user is in the Docker group!
~~~
sudo usermod -aG docker your-user
~~~

Docker Compose Tutorial
------------
1. Clone this repo
~~~
git clone https://github.com/Tueffy/multitier_docker/
~~~
2. Build the web image (XYZ your user id)
~~~
docker build -t XYZ/webgoredis .
~~~
3. Modify the docker-compose.yml and replace the web image with the one you have just build, e.g. replacing the user id
4. Start the stack
~~~
docker-compose up
~~~
5. Test the stack (HAProxy is listening on port 80) and see the counter increasing
~~~
http://yourip
~~~
6. Play around with scaling
~~~
docker-compose scale web=3
~~~
7. Hit further times in the web browser and you see changing host names

Docker Swarm Tutorial
------------
You need three machines for this, all with Docker v1.12 experimental installed. Name the machines: manager, node1 and node2. In my example I used packet.net and got three machines of size Tiny with CentOS 7 on it.

1. On the manager initialize the Docker Swarm
~~~
docker swarm init --listen-addr <managerip>:2377
~~~
Note: ensure that port 2377 is opened in the firewall for accepting connections from the outside, therefore run on the manager
~~~
sudo firewall-cmd --zone=public --add-port=2377/tcp --permanent
~~~
2. Add the two nodes to the Docker Swarm
~~~
docker swarm join <managerip>:2377
~~~
3. Go back to the manager and view the nodes
~~~
docker node ls
~~~
4. If you have not done already in previous tutorial, build the web image and adopt the docker-compose.yml file with the correct name
5. In order to create a DAB (Distributed Application Bundle), get the images for redis and lb
~~~
docker-compose pull redis
docker-compose pull lb
~~~
6. And now we create the bundle:
~~~
docker-compose bundle --fetch-digests
~~~
7. Via the new deploy command, the dab file can be deployed to the Swarm and you can give it a name
~~~
docker deploy -f multitierdocker.dab mymultitiertest
~~~
8. After that you can have a look at the created services
~~~
docker service ls
~~~
9. For detailed information on a service, you can look at the tasks (aka deployed containers)
~~~
docker service tasks mymultitiertest_web
~~~
10. And also here you can scale easily
~~~
docker service scale mymultitiertest_web=3
~~~
 

Docker CLoud Tutorial
------------
That is the most easy one! Ensure that you have an account in cloud.docker.com. Create a node cluster (e.g. also via Packet or AWS) and then just hit the button below:

[![Deploy to Docker Cloud](https://files.cloud.docker.com/images/deploy-to-dockercloud.svg)](https://cloud.docker.com/stack/deploy/)

