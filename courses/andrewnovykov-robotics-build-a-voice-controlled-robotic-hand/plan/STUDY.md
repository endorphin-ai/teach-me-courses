# Study notes
Sources: background knowledge of official docs (Espressif Arduino-ESP32, ESP32Servo library, Arduino language reference, Edge Impulse docs). No web pages were fetched in this run; lessons link the official docs so the student checks versions (Arduino-ESP32 core 3.x, ESP32Servo 3.x).

Key facts used:
- Hobby servos (SG90/MG90S) take a 50 Hz pulse train; pulse width roughly 0.5-2.4 ms maps to 0-180 degrees. The 3.3 V signal from the ESP32 works for most; power comes from a separate 5 V supply; grounds must be shared.
- Five servos can draw several amps when stalled: use a 5 V supply rated 3 A or more, never the USB port or the ESP32 5V pin for all five.
- Avoid boot-strapping pins (GPIO 0, 2, 5, 12, 15) for servos. Servo pins used: 13, 14, 25, 26, 27. Microphone pins: 32 (BCLK), 33 (WS), 34 (data in, input-only).
- INMP441 is an I2S digital MEMS microphone, 3.3 V, 24-bit data in a 32-bit slot; L/R pin to GND = left channel.
- Voice commands: a classic ESP32 cannot run the Espressif ESP-SR command recogniser (it needs an ESP32-S3 with PSRAM), so the course trains a small keyword-spotting model with Edge Impulse and runs it on the ESP32. ESP32-S3 boards are mentioned as an option.
- Typed-command mode over Serial is the safe fallback used to build and test the hand before voice exists.

Scope left out (per brief): custom 3D design, advanced AI, custom PCBs.
Safety: fingers pinch, power polarity, never connect 5 V to a 3.3 V pin.
Untrusted material: none supplied by the student.
