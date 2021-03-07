# generate private key
openssl genrsa -out private.pem 4096
# generate public key
openssl rsa -in private.pem -outform PEM -pubout -out public.pem

# convert private key to pkcs1 form
openssl rsa -in private.pem -out private-pkcs1.pem

# print keys
cat private-pkcs1.pem
echo "\n\n"
cat public.pem

echo "make sure the \"peer\" does not have the private key. only the client is supposed to have it"