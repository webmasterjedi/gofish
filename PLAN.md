# Fish Game — Go Learning Project

## What We're Building
Terminal fishing game using BubbleTea (Elm Architecture TUI framework).
Player casts line, waits for bite, plays catch mini-game, collects fish.

## Teaching Philosophy
- You code, I teach the concepts behind each step
- Each phase introduces new Go concepts tied to real game features
- No hand-waving — understand every line you write

---

## Phase 1: Foundations — Hello BubbleTea
**Go concepts:** packages, imports, structs, interfaces, methods, return values
**Game:** Blank screen with title, quit on `q`

Steps:
1. Understand the BubbleTea Model interface (Init / Update / View)
2. Define your first `model` struct
3. Implement the 3 interface methods
4. Wire up `main()` with `tea.NewProgram`

Checkpoint: App launches, shows title, exits cleanly on `q`

---

## Phase 2: Game State — The Dock
**Go concepts:** enums (iota), custom types, switch statements, string formatting
**Game:** Player stands on dock. Press SPACE to cast.

Steps:
1. Define game states with `iota` (Idle, Casting, Waiting, Reeling, Caught)
2. Update `model` to hold current state + fish line position
3. Switch on state in `View()` to show different screens
4. Handle SPACE keypress in `Update()`

Checkpoint: App transitions between Idle → Casting on SPACE

---

## Phase 3: Time & Events — Waiting for a Bite
**Go concepts:** `tea.Cmd`, `time.Duration`, channels, goroutines (under the hood)
**Game:** After cast, random wait before fish bites

Steps:
1. Learn how BubbleTea handles async via `tea.Cmd`
2. Write a `waitForBite()` command using `time.Sleep` + channel
3. Handle the bite message in `Update()`
4. Show animated waiting state in `View()`

Checkpoint: Cast → wait random seconds → "Fish on the line!" message

---

## Phase 4: The Catch — Mini-Game
**Go concepts:** slices, maps, structs with multiple fields, randomness
**Game:** Timing-based button press to catch fish

Steps:
1. Define a `Fish` struct (name, rarity, weight)
2. Build a fish catalog as a `map[string]Fish`
3. Mini-game: press SPACE at the right moment (moving indicator)
4. Random fish selection weighted by rarity

Checkpoint: Full catch loop — cast → wait → mini-game → catch/miss result

---

## Phase 5: Collection & Stats
**Go concepts:** slices of structs, sorting, formatting, multiple files
**Game:** Fish log, best catches, session stats

Steps:
1. Add `[]Fish` inventory to model
2. Build a log view (scrollable list)
3. Track stats: total caught, rarest caught, total weight
4. Split into multiple `.go` files (model.go, fish.go, views.go)

Checkpoint: Full game loop with persistent session inventory

---

## Phase 6a: Polish — Lipgloss Styling
**Go concepts:** lipgloss styling, constants, more string manipulation
**Game:** Colors, borders, styled UI

Steps:
1. Apply lipgloss styles to each view
2. Add color coding by fish rarity (common=white, uncommon=blue, rare=gold)
3. Add borders and padding to views

Checkpoint: Looks like a real game

---

## Phase 6b: Animation — Harmonica Spring Physics
**Go concepts:** third-party library integration, tea.Tick frame loop, float64 → display mapping
**Library:** `github.com/charmbracelet/harmonica`
**Game:** Spring-physics bobber while waiting, spring-based catch indicator while reeling

Steps:
1. `go get github.com/charmbracelet/harmonica`
2. Add spring state to model: `bobberPos, bobberVel float64` and `indicatorPos, indicatorVel float64`
3. Wire `tea.Tick` in `Init()` to drive animation frames (~60fps)
4. Handle `tickMsg` in `Update()` — call `spring.Update(pos, vel, target)` each frame, store result
5. `StateWaiting` view: render bobber position as ASCII water line (bobber moves up/down)
6. `StateReeling` view: render `[====🎣   ]` indicator whose position is spring-driven toward a random target
7. SPACE at right moment (indicator in catch zone) = `StateCaught`, miss = fish escapes to `StateIdle`

New Go concepts:
- `tea.Tick(interval, func) tea.Cmd` — fires message on timer, enables frame loops
- Spring: `harmonica.NewSpring(harmonica.FPS(60), angularFreq, damping)` — tune feel with two floats
- Mapping float position to terminal characters (float → int → string index)

Checkpoint: Bobber bobs on water while waiting, spring indicator drives real timing-based catch mini-game

---

## Key BubbleTea Concepts (reference)
```
type model struct { ... }          // your game state lives here
func (m model) Init() tea.Cmd      // runs once on start
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd)  // handles events
func (m model) View() string       // renders current frame
```

The loop: BubbleTea calls Update() on every event → you return new model + optional Cmd → View() re-renders.
