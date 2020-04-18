package com.example.demo;

import com.example.demo.controller.HelloController;
import com.example.demo.service.HelloService;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.reactive.AutoConfigureWebTestClient;
import org.springframework.boot.test.autoconfigure.web.reactive.WebFluxTest;
import org.springframework.context.annotation.Import;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.reactive.server.WebTestClient;

// https://howtodoinjava.com/spring-webflux/webfluxtest-with-webtestclient/
// https://mkyong.com/spring-boot/spring-boot-junit-5-mockito/

@RunWith(SpringRunner.class)
@WebFluxTest(HelloController.class)
@Import(HelloService.class)
@AutoConfigureWebTestClient
public class HelloControllerSliceTest {

    @Autowired
    public WebTestClient webClient;

    @Test
    public void shouldSayHello() {
        webClient.get().uri("/hello").accept(MediaType.TEXT_PLAIN)
                .exchange()
                .expectStatus().isOk()
                .expectBody(String.class).isEqualTo("Hello");
    }
}
