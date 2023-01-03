docker-compose down && \
sudo rm -rf $HOME/PISTUFF && \
mkdir $HOME/PISTUFF && \
mkdir $HOME/PISTUFF/Ampgologs && \
touch $HOME/PISTUFF/Ampgologs/ampgo_setup_log.txt && \
mkdir $HOME/PISTUFF/Data && \
mkdir $HOME/PISTUFF/Data/db && \
mkdir $HOME/PISTUFF/Thumbnails && \

python3 create_pages.py && \
docker-compose up