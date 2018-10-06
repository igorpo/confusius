build:
	docker build -t confusius . -f Dockerfile

run: 
	docker run -it confusius