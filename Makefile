run:
	@go run main.go

build:
	@go build -o jenkins-tracer
	@chmod +x jenkins-tracer
	@echo "Build app success -> 'jenkins-tracer'"

package:
	@docker build -t local/jenkins-tracer .

push:
	@docker push local/jenkins-tracer

clean:
	@rm -rf jenkins-tracer tmp default.log
	@echo "Clean up success"