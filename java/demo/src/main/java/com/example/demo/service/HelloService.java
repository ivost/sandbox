package com.example.demo.service;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

@Service
public class HelloService {

    @Value("${app.message}")
    private final String message = "Hello";

    public String sayHello() {
        return message;
    }
}