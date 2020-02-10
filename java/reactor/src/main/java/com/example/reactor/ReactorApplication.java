package com.example.reactor;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import reactor.core.publisher.Flux;
import reactor.core.scheduler.Schedulers;

import java.util.ArrayList;
import java.util.List;

@SpringBootApplication
public class ReactorApplication {

	public static void main(String[] args) {
		SpringApplication.run(ReactorApplication.class, args);
	}

	// https://www.baeldung.com/reactor-core
	public void foo() {
		var s = "Hello";
		List<Integer> elements = new ArrayList<>();

		/*
		The Scheduler interface provides an abstraction around asynchronous code,
		for which many implementations are provided for us.
		Let's try subscribing to a different thread to main:
		 */

		Flux.just(1, 2, 3, 4)
				.log()
				.map(i -> i * 2)
				.subscribeOn(Schedulers.parallel())
				.subscribe(elements::add);
	}
}
