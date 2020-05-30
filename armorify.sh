#!/bin/bash
cat $1 |
tr -d " " |
sed "s/-----BEGINCERTIFICATE-----/-----BEGIN CERTIFICATE-----\\n/" |
sed "s/-----ENDCERTIFICATE-----/\\n-----END CERTIFICATE-----/" |
sed "s/-----BEGINRSAPRIVATEKEY-----/-----BEGIN RSA PRIVATE KEY-----\\n/" |
sed "s/-----ENDRSAPRIVATEKEY-----/\\n-----END RSA PRIVATE KEY-----/" > $1