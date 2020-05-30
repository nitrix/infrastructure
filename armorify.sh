#!/bin/bash
cat $1 |
tr -d " " |
sed "s/-----BEGINCERTIFICATE-----/-----BEGIN CERTIFICATE-----\\n/" |
sed "s/-----ENDCERTIFICATE-----/\\n-----END CERTIFICATE-----/" > $1