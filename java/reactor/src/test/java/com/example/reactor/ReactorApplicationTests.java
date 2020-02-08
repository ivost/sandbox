package com.example.reactor;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import reactor.core.publisher.Flux;

import java.util.ArrayList;
import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
class ReactorApplicationTests {

	@Test
	void contextLoads() {
		List<Integer> elements = new ArrayList<>();

		Flux.just(1, 2, 3, 4)
				.log()
				.subscribe(elements::add);

		assertThat(elements).containsExactly(1, 2, 3, 4);

	}

}
