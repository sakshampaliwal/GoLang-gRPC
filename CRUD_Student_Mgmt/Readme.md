# Golang gRPC CRUD Service

## Introduction

The project is a demonstration of a CRUD (Create, Read, Update, Delete) service implemented in Go (Golang) utilizing gRPC (Google Remote Procedure Call) for communication between clients and servers. The service allows clients to perform basic CRUD operations on student records stored in memory.

## Overview:
The project implements a server-side application that exposes four gRPC methods:

- **Create:** Adds a new student record to the in-memory database.
- **Read:** Retrieves student information based on the provided roll number.
- **Update:** Modifies existing student records based on the provided roll number.
- **Delete:** Removes a student record from the database using the provided roll number.


## Prerequisites
1. Go (Golang) Installed: Ensure that you have Go (Golang) installed on your system. You can download and install it from the official Go website: https://golang.org/dl/.

2. Protocol Buffers Compiler (protoc): The project uses Protocol Buffers for defining the gRPC service and message types. You need to have the Protocol Buffers compiler (protoc) installed on your system.
To Download: https://grpc.io/docs/protoc-installation/

3. Quick Reference for Installation: https://grpc.io/docs/languages/go/quickstart/
## Installation
### Windows
- Click on this link to download Golang. https://go.dev/dl/

### Adding Go to Environment Variables on Windows

1. **Open System Properties**:
   - Right-click on "This PC" or "My Computer" and select "Properties".

2. **Access Environment Variables**:
   - Click on "Advanced system settings".
   - In the System Properties window, click on the "Environment Variables..." button.

3. **Edit Path Variable**:
   - In the Environment Variables window, under "System variables", find the "Path" variable and select it.
   - Click on the "Edit..." button.

4. **Add Go Path**:
   - Click on the "New" button and add the path to the directory where Go is installed. By default, it is `C:\Go\bin`.
   - Click "OK" on each window to save the changes.

### Linux
### Adding Go to Environment Variables on Linux

1. **Open Terminal**:
   - Open a terminal window.

2. **Edit Bash Profile**:
   - Run the following command to open your bash profile in a text editor:
     ```bash
     nano ~/.bashrc
     ```

3. **Add Go Path**:
   - Add the following line at the end of the file:
     ```bash
     export PATH=$PATH:/usr/local/go/bin
     ```

4. **Save and Exit**:
   - Press `Ctrl + X` to exit, then press `Y` to save the changes, and finally press `Enter` to confirm the file name.

5. **Reload Bash Profile**:
   - After saving the changes, run the following command to apply the changes:
     ```bash
     source ~/.bashrc
     ```
6. **To See it is added on path or not**:
   - ```bash
     echo $PATH
     ```


## Command to initalize the Project

**Maintain Dependency Integrity with `go mod tidy`**

To ensure the integrity and consistency of your project's dependencies, it's crucial to update the `go.mod` and `go.sum` files according to the actual usage of dependencies in your codebase. Running `go mod tidy` regularly is considered a good practice as it ensures that your project's dependencies are accurately represented and kept up-to-date.

```bash
go mod tidy
```

## To Run the Server
1. Navigate to the server directory:
    ```bash
    cd ./server
    ```

2. Run the server using `go run`:
    ```bash
    go run ./server/main.go
    ```
### Note
- Make sure to replace `server/main.go` with the correct path to your server file.
- Modify the command according to your file structure if necessary.
  
## Running the Client

To run the client and send a requests to the server:

1. Navigate to the client directory:
    ```bash
    cd ./client
    ```

2. Run the client using `go run`:
    ```bash
    go run ./client/main.go 
    ```

## **Minimal Workflow of Project**

1. **Start the Server:**
    - Run the main application to start the gRPC server.
    - The server listens on a specified port for incoming gRPC requests.

2. **Client Interaction:**
    - Use a gRPC client to interact with the server.
    - Clients can be other applications, scripts, or tools that communicate with the server to perform CRUD operations on student records.

3. **Create Operation:**
    - The client sends a gRPC request to the server to create a new student record.
    - The server receives the request, validates the data, and adds the new record to the in-memory database.
    - The server responds with a success message or an error if the operation fails.

4. **Read Operation:**
    - The client sends a gRPC request to the server to retrieve a student record by providing the roll number.
    - The server searches for the student record based on the provided roll number in the in-memory database.
    - If the record is found, the server responds with the student's information; otherwise, it returns an error message.

5. **Update Operation:**
    - The client sends a gRPC request to the server to update an existing student record.
    - The server receives the request, searches for the record based on the provided roll number, and updates the record with the new data.
    - The server responds with a success message or an error if the operation fails.

6. **Delete Operation:**
    - The client sends a gRPC request to the server to delete a student record by providing the roll number.
    - The server searches for the record based on the provided roll number and removes it from the in-memory database.
    - The server responds with a success message or an error if the operation fails.

7. **Logging:**
    - Throughout the workflow, the server logs relevant information using the Zap logging library.
    - Logging helps track the flow of data within the server, record important events, and diagnose issues.


## Working Diagram
![Architecture_Diagram](https://github.com/sakshampaliwal/GoLang-gRPC/assets/83136084/e157db9f-1fb1-4c25-aeea-6032333c1bc1)

## References
- https://go.dev/doc/
- https://pkg.go.dev/
- https://www.golangprograms.com/
- https://grpc.io/docs/languages/go/basics/
