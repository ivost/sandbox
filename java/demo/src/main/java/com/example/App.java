package com.example;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ApplicationContextInitializer;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.context.support.GenericApplicationContext;

/*
https://blog.tratif.com/2018/03/08/when-two-beans-collide/
https://stackoverflow.com/questions/35217354/how-to-add-custom-applicationcontextinitializer-to-a-spring-boot-application

initializer not called
using app.yml
 */
@SpringBootApplication
public class App {
	public static void main(String[] args) {
		SpringApplication app = new SpringApplication(App.class);
		app.addInitializers(new CustomAppCtxInitializer());
		ConfigurableApplicationContext ctx = app.run();
	}

	private static class CustomAppCtxInitializer
			implements ApplicationContextInitializer<GenericApplicationContext> {

		@Override
		public void initialize(GenericApplicationContext applicationContext) {
			applicationContext
					.getDefaultListableBeanFactory()
					.setAllowBeanDefinitionOverriding(false);
		}
	}
}
