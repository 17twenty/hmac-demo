import hashlib
import hmac
import json

secret = 'CLIENT_SECRET'
data = dict()
data['foo'] = 'bar'
print("Encoding:",json.dumps(data).encode('utf-8'))

signature = '6ee84d2105efe2f41f1e403295679505577c123401204021ac290e2a2a55f542'
signature_computed = hmac.new(
    key=secret.encode('utf-8'),
    msg=json.dumps(data,separators=(',', ':')).encode('utf-8'), # Get dictionary, as JSON and encode to utf-8
    digestmod=hashlib.sha256
).hexdigest()
if not hmac.compare_digest(signature, signature_computed):
    print("Expected>",signature)
    print("Got>",signature_computed)
    print("Invalid payload")
else:
    print("Matched!")