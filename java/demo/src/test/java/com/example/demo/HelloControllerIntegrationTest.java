package com.example.demo;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.reactive.AutoConfigureWebTestClient;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.reactive.server.WebTestClient;

// almost the same as SliceTest - but uses spring boot and starts the whole app
@RunWith(SpringRunner.class)
@SpringBootTest
@AutoConfigureWebTestClient
public class HelloControllerIntegrationTest {

    @Autowired
    private WebTestClient webClient;

    @Test
    public void shouldSayHello() {

        webClient.get().uri("/hello").accept(MediaType.TEXT_PLAIN)
                .exchange()
                .expectStatus().isOk()
                .expectBody(String.class)
                .isEqualTo("Hello");
    }
}
