package com.example.demo;

import com.example.demo.controller.HelloController;
import org.junit.Test;
import reactor.core.publisher.Mono;
import reactor.test.StepVerifier;

public class HelloControllerUnitTest {
    HelloController controller = new HelloController();

    @Test
    public void shouldSayHello() {
        Mono<String> result = controller.hello();
        StepVerifier.create(result)
                .expectNext("Hello").verifyComplete();
    }
}
