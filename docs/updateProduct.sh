#!/bin/bash 

curl -H "Content-type: application/json" -X PUT -v localhost:8080/api/products/1 \
    -d '{"name":"Art","model":"A1","price":2099,"description":"ARTlessons for children age of 3-6","image_url":"http://s3.edushop.com/static/images/art_a1.jepg","video_url":"http://s3.edushop.com/static/video/en_a1.mp4","Capacity":99,"create_at":0,"created_by":"","modified_on":0,"modified_by":"","labels":"addedbyMlb","state":0}'
