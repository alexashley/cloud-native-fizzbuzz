# cloud native fizzbuzz

badges go here

## Background

[FizzBuzz](https://en.wikipedia.org/wiki/Fizz_buzz) is a complex engineering problem that requires a combined mastery of advanced mathematical concepts and programming techniques, such as congruence relations and iteration. 
It's often used in the final stages of job interviews, after trivial screening questions about red-black trees and system architecture. 
Calculating the result requires immense computational power and cannot be done trivially; as such developers in interviews often turn to managed services in order to solve the problem -- either by writing code 
that interacts with these services directly or even just mentioning that a proper engineering solution would almost certainly require the use of such a service.

Unfortunately for developers, these FizzBuzz-as-a-service providers (hereafter FaaS) are often built around insecure, unscalable, and fragile legacy architectures. 
For example, a common approach is to precompute the answer for various numbers, store the result in a table, and use a stored procedure to query the table for any given range. 
This is often accompanied by a small SOAP web service or RPC interface that acts as a veneer to the underlying infrastructure. 
Consequently, these services are difficult to use, impossible to debug, and it's common for them to ignore authentication, authorization, observability. 

However, there is a light at the end of the tunnel. 
Thanks to the immense growth and proliferation of cloud native technologies, it's now possible to build a highly secure, infinitely scalable, and maintainable solution to this once thorny problem. 
This repository is an example of the use of these technologies, specifically Go, Kubernetes, Helm, Docker, Istio, NATS, and Open Policy Agent. 
It demonstrates how FizzBuzz can be split into separate, logical domains, each with their own microservices communicating via a service mesh. See the [architecture](#architecture) for more information. 

## Architecture

## Local Development
