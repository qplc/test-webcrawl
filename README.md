# test-webcrawl
Steps to execute binary:
1. Extract test-webcrawl.tar.gz
2. Issue below command
	cd test-webcrawl
3. Issue below command to start module
	on mac
	./test-webcrawl-mac
	or
	on linux
	./test-webcrawl-linux
4. Send http call to fetch web response	
e.g.
curl -X POST http://localhost:8082/webcrawl   -H 'Content-Type: application/json' -H 'cache-control: no-cache'   -d '{
    "urls": [
        "https://google.com",
        "https://github.com"
    ]
}'
