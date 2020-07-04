#include <zephyr.h>
#include <device.h>
#include <drivers/gpio.h>

#define LED_PORT  DT_ALIAS_LED0_GPIOS_CONTROLLER
#define LED		  DT_ALIAS_LED0_GPIOS_PIN
#define SLEEP_TIME	1000

struct device *dev;

void delay() 
{
    k_sleep(SLEEP_TIME);
}

void init() 
{
    delay();
    dev = device_get_binding(LED_PORT);
    for (int i=0; i<4; i++) 
    {
        gpio_pin_configure(dev, LED+i, GPIO_OUTPUT);
    }
    printk("Init 0.0.5.10\n");
}

void led(int idx, int val) 
{
    gpio_pin_set(dev, LED+idx, val);
    delay();
}

void main(void)
{
    init();
    while (1) {
        printk("loop\n");
        led(0, 1);
        led(1, 1);
        led(3, 1);
        led(2, 1);

        led(0, 0);
        led(1, 0);
        led(3, 0);
        led(2, 0);
    }
}
