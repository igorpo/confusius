build:
	go build .

clean: 
	rm ./confusius 

run: build 
	$(shell ./confucius)