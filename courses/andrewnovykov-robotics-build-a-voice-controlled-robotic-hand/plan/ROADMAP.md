# Roadmap: voice-controlled robotic hand (8 x 60 min; weeks 1-4, two lessons per week)
Project: a 5-servo hand on an ESP32 that opens, closes and points on spoken commands.

| # | Lesson | Week | Milestone toward the goal |
|---|--------|------|---------------------------|
| 1 | 01-meet-the-circuit | 1 | LED blinks from your sketch; parts ordered (hand, servos, mic, supply) |
| 2 | 02-first-code | 1 | Sketch answers typed words on the Serial Monitor |
| 3 | 03-servo-motion | 2 | One servo sweeps, powered safely |
| 4 | 04-the-hand | 2 | Hand assembled; open(), close(), point() work |
| 5 | 05-typed-commands | 3 | Type "point" and the hand does it, smoothly |
| 6 | 06-microphone | 3 | Mic level meter; a clap moves a finger |
| 7 | 07-keyword-spotting | 4 | Trained model hears open/close/point |
| 8 | 08-voice-hand | 4 | Say "close" and the hand closes: goal reached |

## Spiral
| Concept | First taught | Revisited in | What is added |
|---|---|---|---|
| GPIO pins (c02) | 01 | 03, 06 | pins as PWM output, then as I2S input; strapping-pin care |
| Serial (c07) | 02 | 05, 07, 08 | whole-line commands, then debug prints of confidence |
| Functions (c06) | 02 | 04, 08 | gesture functions, then one act() called by voice |
| if/else (c05) | 02 | 05, 06, 08 | command matching, threshold, confidence gate |
| Common ground (c09) | 03 | 04, 06 | five servos, mic power |
| Servo angle (c10) | 03 | 04, 05 | per-finger arrays, smooth stepping |
| Command table (c15) | 05 | 08 | voice labels use the same table |
| Loudness/threshold (c19) | 06 | 07, 08 | threshold becomes model confidence |
| Confidence (c22) | 07 | 08 | cooldown, tuning |

## Parts timeline
The 3D-printed hand ships slowly: order it in lesson 1 so it arrives by lesson 4. Lessons 1-3 need only the ESP32, a USB cable, an LED, a resistor and one servo.
