# Jenkins

CI CD with jenkins pipeline 


## Installation (DigitalOcean) & Configuration

0. DigitalOcean [Droplet](https://cloud.digitalocean.com/droplets)

1. For Droplet :

        $ sudo nano /etc/ssh/sshd_config

            PermitRootLogin yes 
            PasswordAuthentication yes

        $ sudo service ssh restart
    
2. Access the droplet with ssh:

        $ ssh root@<droplet_ip>

        # sudo bash install-jenkins.sh

3. Access Jenkins though the browser:  `<droplet_ip>:8080`  ->  install suggested plugins

        # cat /var/jenkins_home/secrets/initialAdminPassword (Jenkins password)