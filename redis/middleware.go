package main

import (
	"fmt"
	"log"
	"os"
	Store "redis-app/store"
)

type LoggingMiddleware struct {
	innerStorer Store.IStore //let inject so that we can implement ttl and kv
	logger 		*log.Logger
}

func NewLoggingMiddleware(inner Store.IStore) *LoggingMiddleware {
	return &LoggingMiddleware{
		innerStorer: inner,
		logger: log.New(os.Stdout, "Log ", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Set(key, value string) (error) {
	l.logger.Printf("Set %q", key) //we perform side effect by just logging
	return l.innerStorer.Set(key, value) //this will go into the engine of the key value store
}

func (l *LoggingMiddleware) Delete(key string) (error) {
	l.logger.Printf("Delete %q", key)
	_, done := l.innerStorer.Pop(key)
	if !done {
		return fmt.Errorf("key %q does not exist", key)
	}
	return nil
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	l.logger.Printf("Get %q", key) //we perform side effect by just logging
	val, err := l.innerStorer.Get(key)
	if err != nil {
		l.logger.Printf("Get (%v) miss %v", key, err)
	} else {
		l.logger.Printf("Get (%v) hit",key)
	}
	return val, err
}

func (l *LoggingMiddleware) Len() int {
	l.logger.Printf("Len")
	return l.innerStorer.Len()
}

func (l *LoggingMiddleware) Pop(key string) (string, bool) {
	l.logger.Printf("Len")
	return l.innerStorer.Pop(key)
}