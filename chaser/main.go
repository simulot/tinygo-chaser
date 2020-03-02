package main

import (
	"time"

	"github.com/simulot/drivers/shiftregister"
)

func main() {
	d := shiftregister.New(
		shiftregister.EIGHT_BITS,
		latch, // D12 Pin latch connected to ST_CP of 74HC595 (12)
		clock, // D11 Pin clock connected to SH_CP of 74HC595 (11)
		data,  // D10 Pin data connected to DS of 74HC595 (14)
	)
	d.Configure()

	for {
		for _, c := range chasers {
			runChaser(d, c, 100*time.Millisecond, 5*time.Second)
		}
	}
}

func runChaser(sr *shiftregister.Device, patterns []uint32, stepDuration time.Duration, totalDuration time.Duration) {
	stopAt := time.Now().Add(totalDuration)
	for time.Now().Before(stopAt) {
		for _, pattern := range patterns {
			sr.WriteMask(pattern)
			time.Sleep(stepDuration)
		}
	}
}

var chasers = [][]uint32{
	{
		0x55,
		0xAA,
	},
	{
		0b00000000,
		0b10001000,
		0b11001100,
		0b11101110,
		0b11111111,
	},
	{
		0b00000001,
		0b00000010,
		0b00000100,
		0b00001000,
		0b00010000,
		0b00100000,
		0b01000000,
		0b10000000,
		0b01000000,
		0b00100000,
		0b00010000,
		0b00001000,
		0b00000100,
		0b00000010,
		0b00000001,
	},
	{
		0b00000001,
		0b00000011,
		0b00000110,
		0b00001100,
		0b00011000,
		0b00110000,
		0b01100000,
		0b11000000,
		0b10000000,
	},
	{
		0b00010000,
		0b00101000,
		0b01000100,
		0b10000010,
		0b00000001,
	},
	{
		0b00000001,
		0b00000010,
		0b00000100,
		0b00001000,
		0b00010000,
		0b00100000,
		0b01000000,
		0b10000000,
		0b10000001,
		0b10000010,
		0b10000100,
		0b10001000,
		0b10010000,
		0b10100000,
		0b11000000,
		0b11000001,
		0b11000010,
		0b11000100,
		0b11001000,
		0b11010000,
		0b11100000,
		0b11100001,
		0b11100010,
		0b11100100,
		0b11101000,
		0b11110000,
		0b11110001,
		0b11110010,
		0b11110100,
		0b11111000,
		0b11111001,
		0b11111010,
		0b11111100,
		0b11111101,
		0b11111110,
		0b11111111,
		0b00000000,
		0b11111111,
		0b00000000,
		0b11111111,
		0b00000000,
	},
}
