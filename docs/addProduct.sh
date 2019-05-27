#!/bin/bash

# Add records for product


for img_id in `seq 1 15`;
do
    IMAGE_URL="https://picsum.photos/id/11$img_id/600/300"

curl -H "Content-type: application/json" -X POST -v localhost:8080/api/products/ \
    -d "{\"name\":\"english_A$img_id\",\"model\":\"A$img_id\",\"price\":$((99+$img_id/10)),\"description\":\"english lessons for children age of 3-6\",\"image_url\":\"$IMAGE_URL\",\"video_url\":\"$IMAGE_URL\",\"Capacity\":99,\"create_at\":0,\"created_by\":\"\",\"modified_on\":0,\"modified_by\":\"\",\"labels\":\"\",\"state\":1}"
done
