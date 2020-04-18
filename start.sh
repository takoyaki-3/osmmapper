cd /osm
wget http://download.geofabrik.de/asia/japan/kanto-latest.osm.pbf
wget http://m.m.i24.cc/osmconvert64
./osmconvert64 kanto-latest.osm.pbf > kanto.osm
