# Lazarus

Lazarus queries Prometheus Alertmanager API and will run a script in response to the alert.


![alt text](https://upload.wikimedia.org/wikipedia/commons/0/0b/%27The_Raising_of_Lazarus%27%2C_tempera_and_gold_on_panel_by_Duccio_di_Buoninsegna%2C_1310%E2%80%9311%2C_Kimbell_Art_Museum.jpg)

## Setup

Lazarus will look for yaml config files under /etc/lazarus/conf.d/ to determine what action to take to resolve the alert.

1. Download Release binary or run ```Make build``` to build yourself and place in /usr/local/bin/
2. Install systemd file into /etc/systemd/system/
2. Put all configs in /etc/lazarus/conf.d/ (See lazarus.yml for example)
3. Make sure script that lazarus will call is executable
5. systemctl start lazarus.service
6. Logs to syslog 
