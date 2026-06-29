

echo "Testing with client CA cert:"
curl --cert certs/client.crt --key certs/client.key \
     --cacert certs/ca.crt \
     https://localhost:8443/missions

echo ""

echo "Testing without client CA cert:"
curl --cacert certs/ca.crt https://localhost:8443/missions