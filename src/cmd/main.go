package main

func main() {
	// receive request - grpc / http -> FE generate typescript
	
	// authenticate

	// check the schema if in json
	// validate if all required provided -> middleware or interceptors?


	/// Goroutine 1
	//  settings to create infra to run modeler and solver

	/// Goroutine 2 (talks to 1 and later 3 for input)
	//  solver preparation on the infra is ready
	//  create required config maps and data
	//  should we need access to secrets and service accounts to connect and fetch data?

	/// Goroutine 3
	//  modeler template engines
}