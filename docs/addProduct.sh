#!/bin/bash

curl -H "Content-type: application/json" -X POST -v localhost:8080/api/products/ -d '{"name":"english_A1","model":"A1","price":99,"description":"english lessons for children age of 3-6","image_url":"http://s3.edushop.com/static/images/en_a1.jepg","video_url":"http://s3.edushop.com/static/images/en_a1.jepg","Capacity":99,"create_at":0,"created_by":"","modified_on":0,"modified_by":"","labels":"","state":1}'

curl -H "Content-type: application/json" -X PUT -v localhost:8080/api/products/1 -d '{"name":"english_A1","model":"A1","price":99,"description":"english lessons for children age of 3-6","image_url":"http://s3.edushop.com/static/images/en_a1.jepg","video_url":"http://s3.edushop.com/static/images/en_a1.jepg","Capacity":99,"create_at":0,"created_by":"","modified_on":0,"modified_by":"","labels":"addedbyMlb","state":0}'
