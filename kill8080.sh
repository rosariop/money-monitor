#/bin/bash
sudo lsof -t -i tcp:8080 -s tcp:listen | sudo xargs kill