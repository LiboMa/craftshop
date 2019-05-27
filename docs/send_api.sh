#!/bin/bash 


TS=$(date +%FT\%T)
curl "https://api.huobi.pro/v1/order/orders?AccessKeyId=$huobi_access_key&SignatureMethod=HmacSHA256&SignatureVersion=2&Timestamp=$TS"
