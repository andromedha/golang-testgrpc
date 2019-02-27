# golang-testgrpc

A small application to try out grpc with golang.

This repository contains a client and a server application.

You start one server and multiple clients.
Every client register and subscribe himself by the server.
Each time a client subscribe to the server, the server send data to each connected client back.
