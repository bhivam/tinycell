package main

type Cell struct {
	expr        Expr
	value       any
	calculating bool
}
