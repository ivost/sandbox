package com.example.reactor;

import lombok.extern.slf4j.Slf4j;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import reactor.core.publisher.Flux;
import reactor.core.scheduler.Schedulers;

import java.time.Duration;
import java.util.ArrayList;

import static org.assertj.core.api.Assertions.assertThat;

/*
 https://www.baeldung.com/reactor-core
 */

@SpringBootTest
@Slf4j
class ReactorTest {

	@Test
	void delaying() {
		Flux.range(0, 5)
				.delayElements(Duration.ofMillis(1))
				.elapsed()
				.log()
				.subscribe(e -> log.info("Elapsed {} ms, {}", e.getT1(), e.getT2()))
		;
	}

	@Test
	void scheduling() {
		var list = new ArrayList<Integer>();
		var immediate = Schedulers.immediate();
		var stream = Flux.range(100, 5);
		stream
				.log()
				.subscribeOn(immediate)
				.subscribe(
						el -> list.add(el),
						err -> log.error(err.getMessage()),
						() -> log.info("Completed")
				);
		assertThat(list).hasSize(5);
	}

	@Test
	void buffering() {
		var immediate = Schedulers.immediate();
		var stream = Flux.range(1, 12)
				//.delayElements(Duration.ofMillis(100))
				//.publishOn(Schedulers.boundedElastic())
				//.publish()
				;

		stream
				//.buffer(Duration.ofSeconds(1))
				.log()
				//.subscribeOn(immediate)
				.subscribe(
						el -> log.info("element {}", el),
						err -> log.error(err.getMessage()),
						() -> log.info("Completed")
				);
	}

}
