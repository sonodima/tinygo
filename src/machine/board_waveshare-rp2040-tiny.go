//go:build waveshare_rp2040_tiny

// This file contains the pin mappings for the Waveshare RP2040-Tiny boards.
//
// Waveshare RP2040-Tiny is a microcontroller using the Raspberry Pi RP2040 chip.
//
// - https://www.waveshare.com/wiki/RP2040-Tiny
package machine

// Onboard crystal oscillator frequency, in MHz.
const xoscFreq = 12 // MHz

// Digital Pins
const (
	D0  Pin = GPIO0
	D1  Pin = GPIO1
	D2  Pin = GPIO2
	D3  Pin = GPIO3
	D4  Pin = GPIO4
	D5  Pin = GPIO5
	D6  Pin = GPIO6
	D7  Pin = GPIO7
	D8  Pin = GPIO8
	D9  Pin = GPIO9
	D10 Pin = GPIO10
	D11 Pin = GPIO11
	D12 Pin = GPIO12
	D13 Pin = GPIO13
	D14 Pin = GPIO14
	D15 Pin = GPIO15
	D16 Pin = GPIO16
	D26 Pin = GPIO26
	D27 Pin = GPIO27
	D28 Pin = GPIO28
	D29 Pin = GPIO29
)

// Analog pins
const (
	A0 Pin = GPIO26
	A1 Pin = GPIO27
	A2 Pin = GPIO28
	A3 Pin = GPIO29
)

// Onboard LEDs
const (
	NEOPIXEL = GPIO16
	WS2812   = GPIO16
)

// I2C pins
const (
	I2C0_SDA_PIN Pin = GPIO0
	I2C0_SCL_PIN Pin = GPIO1

	I2C1_SDA_PIN Pin = GPIO2
	I2C1_SCL_PIN Pin = GPIO3
)

// SPI pins
const (
	SPI0_SCK_PIN Pin = GPIO6
	SPI0_SDO_PIN Pin = GPIO3
	SPI0_SDI_PIN Pin = GPIO4

	SPI1_SCK_PIN Pin = GPIO10
	SPI1_SDO_PIN Pin = GPIO11
	SPI1_SDI_PIN Pin = GPIO12
)

// UART pins
const (
	UART0_TX_PIN = GPIO0
	UART0_RX_PIN = GPIO1
	UART_TX_PIN  = UART0_TX_PIN
	UART_RX_PIN  = UART0_RX_PIN
	UART1_TX_PIN = GPIO8
	UART1_RX_PIN = GPIO9
)

var DefaultUART = UART0

// USB CDC identifiers
const (
	usb_STRING_PRODUCT      = "RP2040-Tiny"
	usb_STRING_MANUFACTURER = "Waveshare"
)

var (
	usb_VID uint16 = 0x2e8a
	usb_PID uint16 = 0x000a
)
