/*
 * Copyright (c) 2016 Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

#include <zephyr.h>
#include <device.h>
#include <drivers/gpio.h>

#define LED_PORT  DT_ALIAS_LED0_GPIOS_CONTROLLER
#define LED		  DT_ALIAS_LED0_GPIOS_PIN

/* 1000 msec = 1 sec */
#define SLEEP_TIME	500

void main(void)
{
	u32_t cnt = 0;
	struct device *dev;

	printk("START\n");

	dev = device_get_binding(LED_PORT);
	gpio_pin_configure(dev, LED, GPIO_OUTPUT);

	while (1) {
		printk("LOOP %d\n", cnt);
		gpio_pin_set(dev, LED, cnt % 2);
		cnt++;
		k_sleep(SLEEP_TIME);
	}
}
