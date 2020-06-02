package com.example.demo.controller;

import com.example.demo.service.HelloService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@RestController
@Slf4j
public class HelloController {

    final HelloService helloService;
    public HelloController() {
        this.helloService = new HelloService();
    }

    public HelloController(HelloService helloService) {
        this.helloService = helloService;
    }

    @GetMapping("/hello")
    public Mono<String> hello() {
        log.info("hello controller");
        return Mono.just(helloService.sayHello());
    }

}
