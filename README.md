Multitier Docker Example
------------
This application stack is comprised out of three images
* HA Proxy
* Web Server written in Go (which increases a counter in Redis for every hit)
* Redis KV store

Required Installation Process
------------
1. Install Docker v1.12 experimental
    # curl -sSL https://experimental.docker.com/ | sh

2. Install Docker Compose
    # curl -L https://github.com/docker/compose/releases/download/1.8.0-rc2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose

3. Start Docker Service
    # sudo service docker start

4. Ensure that your user is in the Docker group!
    # sudo usermod -aG docker your-user


Docker Compose Tutorial
------------
1. Clone this repo
    # git clone https://github.com/Tueffy/multitier_docker/

2. Build the web image (XYZ your user id)
    # docker build -t XYZ/webgoredis .

3. Modify the docker-compose.yml and replace the web image with the one you have just build, e.g. replacing the user id

4. Start the stack
    # docker-compose up

5. Test the stack (HAProxy is listening on port 80) and see the counter increasing
    # http://yourip

6. Play around with scaling
    # docker-compose scale web=3

7. Hit further times in the web browser and you see changing host names

Docker Swarm Tutorial
------------


Docker CLoud Tutorial
------------

[![Deploy to Docker Cloud](https://files.cloud.docker.com/images/deploy-to-dockercloud.svg)](https://cloud.docker.com/stack/deploy/)

