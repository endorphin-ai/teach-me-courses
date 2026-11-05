# Tutorial 1: Meet the Circuit and the ESP32

> **Time**: 60 min | **Difficulty**: Beginner
> **Prerequisites**: none

## What You'll Learn

- What a circuit is, in plain words: voltage, current and ground.
- What an ESP32 is and what its pins can do.
- How to write, compile and upload a sketch from your Mac.
- What to order for the robotic hand, so parts arrive in time.

## Why This Matters

:::hook {kind="problem"}
In four weeks you will say "close" to a plastic hand and its fingers will curl. Right now you have a green board and a cable. The gap is smaller than it looks: every finger movement starts as one pin on that board switching on and off.
Today you make one pin blink an LED. That tiny blink is the same trick the hand uses.
:::

**Key idea: an ESP32 is a tiny computer whose pins you can switch on and off from code, and everything in the hand, from servos to the microphone, hangs on those pins.**

Three points support it:
1. Electricity only does work around a complete loop (a circuit).
2. The ESP32's pins are the doors between your code and the real world.
3. A fixed routine (write, compile, upload) puts your code on the board.

## Walkthrough

### Step 1: A circuit is a loop

:::note {kind="key"}
Circuit: a closed loop in which electric current is pushed by voltage from a power source, through a part, and back to ground.
:::

:::analogy {concept="c01"}
title: A water loop with a pump
boundary: Water can be seen and leaks messily; electricity is invisible and a "leak" to the wrong place can burn a part.
---
Picture a loop of pipe with a pump. The pump's pressure is **voltage**, measured in volts (V). The amount of water flowing is **current**, measured in amps (A). The pipe that returns water to the pump is **ground** (GND). A water wheel in the loop is your LED or servo: it only turns if water flows through it.
:::

Your ESP32 runs at **3.3 V**. The servos later will want **5 V**. Mixing them up is the most common way to break a board, so we say the voltage out loud every time.

:::example {concept="c01" level="minimal"}
title: The loop on paper
---
```text
# led-loop.txt
3.3 V pin --> resistor --> LED --> GND pin
   (pump)     (narrow     (water   (return
               pipe)       wheel)   pipe)
```
Break any link and the LED goes dark. The resistor (about 220 ohms) is a narrow pipe that stops too much current from burning the LED.
:::

:::checkpoint {concept="c01"}
question: You wire 3.3 V to an LED but forget to connect the other end to GND. Does it light, and why?
expected: No. Current needs a complete loop back to ground; with the loop broken nothing flows.
---
:::

### Step 2: The ESP32 and its pins

:::note {kind="key"}
GPIO pin: a General Purpose Input/Output pin on the microcontroller that your program sets as an output (it drives HIGH = 3.3 V or LOW = 0 V) or reads as an input.
:::

A **microcontroller** is a small computer on one chip that runs one program forever. The ESP32 board has a microcontroller, a USB chip, memory and Wi-Fi/Bluetooth. Its pins are labelled by GPIO number, such as GPIO2 or GPIO13.

:::analogy {concept="c02"}
title: A row of light switches your program can flip
boundary: A real switch is flipped by a hand; a pin flips in millionths of a second and can also listen, which a wall switch cannot.
---
Each GPIO pin is a switch on a wall. Your code can flip it up (HIGH) or down (LOW), or it can leave the switch alone and just watch whether someone else flipped it (input).
:::

The pin of the small built-in LED varies by board (many use GPIO2, which is also a start-up "strapping" pin). To avoid surprises we use an external LED on GPIO4 below.

:::note {kind="warning"}
Use GPIO4 for the LED today. A few pins have special jobs while the board starts up, so the servo lessons name the safe ones to use.
:::

:::example {concept="c02" level="realistic"}
title: Wiring an external LED
---
On a breadboard: GPIO4 -> 220 ohm resistor -> LED long leg (the anode); LED short leg -> GND pin on the ESP32. Only one wire per row of the breadboard strip connects; the LED's two legs must be in different strips.
:::

:::checkpoint {concept="c02"}
question: What is the difference between a pin set as OUTPUT and one set as INPUT?
expected: An output is driven by the program to HIGH or LOW; an input is read by the program to see what voltage something else applies.
---
:::

### Step 3: Install the tools on your Mac

1. Install the **Arduino IDE 2** from arduino.cc.
2. Open Settings, and in "Additional boards manager URLs" paste this URL (check the Espressif guide linked below in case it has changed): `https://espressif.github.io/arduino-esp32/package_esp32_index.json`
3. Open Boards Manager, search `esp32`, install "esp32 by Espressif Systems". Expected result: it shows as "INSTALLED", and "ESP32 Dev Module" appears in the board list.
4. Plug in the board with a **data** USB cable. Choose Tools > Board > "ESP32 Dev Module" (or your exact board) and Tools > Port > the entry that appeared (it looks like `/dev/cu.usbserial-...` or `cu.SLAB_USBtoUART`).

If no port appears, your Mac may need the USB driver for the board's chip (CP210x or CH340); the board's seller page names which one.

### Step 4: Write, compile, upload

:::note {kind="key"}
Sketch workflow: you write a sketch in the editor, the IDE compiles it into machine code (catching typos), and uploads it into the board's memory where it runs by itself.
:::

:::analogy {concept="c03"}
title: Writing a recipe, translating it, handing it to the cook
boundary: A cook can ask questions; the board only follows what was uploaded, and forgets nothing until you upload again.
---
You write the recipe (sketch). A translator turns it into the cook's language (compile). You hand the translated recipe to the cook (upload). The cook then repeats it forever.
:::

Every sketch has two parts: `setup()` runs once at the start, and `loop()` runs again and again.

:::example {concept="c03" level="minimal"}
title: Blink
---
```cpp
// blink.ino
const int LED_PIN = 4;   // the pin the external LED is wired to

void setup() {
  pinMode(LED_PIN, OUTPUT);   // face the pin outward
}

void loop() {
  digitalWrite(LED_PIN, HIGH);  // 3.3 V: LED on
  delay(500);                   // wait half a second
  digitalWrite(LED_PIN, LOW);   // 0 V: LED off
  delay(500);
}
```
Press the check mark (compile), then the arrow (upload). The LED should blink once per second.
:::

:::example {concept="c03" level="challenging"}
title: Reading a compile error
---
```cpp
// blink_broken.ino
void setup() {
  pinMode(LED_PIN, OUTPUT)
}
```
This sketch is broken on purpose. The IDE shows red text such as `expected ';' before '}'` and `'LED_PIN' was not declared`. Read from the **first** error: fixing the missing semicolon and adding the `const int LED_PIN` line removes both.
:::

:::checkpoint {concept="c03"}
question: You change delay(500) to delay(100) in the editor. The LED still blinks slowly. What step did you forget?
expected: You must compile and upload again; the board still runs the old uploaded program.
---
:::

:::diagram
title: From your keyboard to a blinking LED
---
```mermaid
flowchart LR
  W[Write sketch]:::accent --> C[Compile]:::info
  C --> U[Upload by USB]:::info
  U --> B[ESP32 runs it]:::good
  B --> P[Pin 4 goes HIGH and LOW]:::good
  P --> L[LED blinks]:::good
```
:::

### Step 5: Plan your parts

The hand is the slow part: 3D-printed hands often take 1 to 3 weeks to arrive. Order it today: lesson 4 (the start of week 2) needs it, and until then you can use one servo and a paper strip as a stand-in finger.

| Part | Why | Rough price |
|---|---|---|
| 3D-printed tendon-driven robotic hand (ask the seller to include fishing line or cord) | The body | $20-40 |
| 5 micro servos, SG90 or MG90S (metal gear is sturdier) | One per finger | $10-15 |
| INMP441 I2S microphone breakout | The ears | $3-6 |
| 5 V power supply, at least 3 A, with a barrel jack to screw-terminal adapter | Powers the servos | $8-12 |
| Breadboard, jumper wires (male-to-male and male-to-female), LEDs, 220 ohm resistors | Wiring | $8 |
| 470-1000 microfarad capacitor, 10 V or more | Smooths power spikes | $1 |

Specs to check before you buy: servos must be 4.8-6 V micro servos (SG90 or MG90S size); the supply must say 5 V and 3 A or more; the microphone must be the I2S type (INMP441), not an analog one; the capacitor must be rated 10 V or more. Total fits inside $50-$100.

**Try it now:** open your browser and put the six parts in a basket, then check the total against your budget before you order. Check that the hand you pick says it takes micro servos and gives its finger count.

:::note {kind="warning"}
Do not buy a hand that needs larger servos or gives no servo size. If unsure, message the seller and ask: "Does it use SG90/MG90S micro servos, one per finger?"
:::

## Common Mistakes

:::mistake {concept="c03" title="Using a charge-only USB cable"}
::wrong
The board powers on (a light glows) but no port shows up in the Tools > Port menu, so you keep reinstalling drivers.
::right
Swap to a cable that is known to carry data (one that has moved files from a phone), then re-open the Port menu.
::why
Many cheap cables only have the two power wires, so the Mac can power the board but cannot talk to it.
:::

:::mistake {concept="c02" title="Forgetting pinMode"}
::wrong
```cpp
// blink.ino
void setup() { }
void loop() { digitalWrite(4, HIGH); }
```
::right
```cpp
// blink.ino
void setup() { pinMode(4, OUTPUT); }
void loop() { digitalWrite(4, HIGH); }
```
::why
A pin must be told its direction before it is driven; skipping pinMode can leave it weakly driven or silent.
:::

:::mistake {concept="c01" title="Connecting an LED without a resistor"}
::wrong
GPIO4 straight to the LED, LED to GND.
::right
GPIO4 to a 220 ohm resistor to the LED to GND.
::why
The resistor limits current. Without it the LED, or the pin, can overheat.
:::

## Try It Yourself

:::challenge
title: Morse "S"
---
::task
Make the LED flash three short blinks (100 ms on, 100 ms off), then pause one second, forever.
::hint
Write the on/off pair once, then copy it three times inside loop(). Next lesson you will learn a shorter way.
::solution
```cpp
// morse.ino
const int LED_PIN = 4;

void setup() {
  pinMode(LED_PIN, OUTPUT);
}

void loop() {
  for (int i = 0; i < 3; i++) {   // repeat the next two lines 3 times
    digitalWrite(LED_PIN, HIGH);
    delay(100);
    digitalWrite(LED_PIN, LOW);
    delay(100);
  }
  delay(1000);
}
```
:::

1. Wire the LED and make it blink at a rhythm of your choice.
2. Open the Tools > Serial Monitor menu and just look: nothing prints yet, and that is fine.
3. Break your sketch on purpose in two different ways and write down the error text each gives.

:::quiz {ref="q1"}
:::

:::quiz {ref="q5"}
:::

## Key Takeaways

:::takeaways
- A circuit is a complete loop: voltage pushes current from a source through a part and back to ground.
- The ESP32 runs at 3.3 V, and its GPIO pins are switches your code controls.
- A sketch has setup() once and loop() forever; write, compile, upload, repeat.
- Order the hand early; it is the slow part.
:::

We come back to GPIO pins in [lesson 3: Servo Motion](../03-servo-motion/lesson.md), where one pin sends a signal that sets an angle instead of just on and off.

## What's Next

You can switch a pin, but you can only give it fixed instructions. In [lesson 2: Your First Code](../02-first-code/lesson.md) you learn the four building blocks of code that let the board remember, decide, reuse and talk to you.

:::bridge {to="02-first-code"}
Your board blinks the same way forever. Next you teach it to remember values, make decisions and answer you.
:::

:::homework {kind="build"}
Order the hand and the parts from Step 5 today (the hand has the longest shipping time). While you wait, finish the Morse challenge and keep the blink sketch: lesson 3 reuses the wiring habits.
:::

## Further Reading

:::docs
title: Arduino-ESP32 installation guide
url: https://docs.espressif.com/projects/arduino-esp32/en/latest/installing.html
---
The official steps for adding the ESP32 boards to the Arduino IDE; use the current board-manager URL it lists.
:::

:::docs
title: Arduino language reference
url: https://www.arduino.cc/reference/en/
---
Every function we use (pinMode, digitalWrite, delay) is explained here; this lesson adds the order to learn them.
:::
