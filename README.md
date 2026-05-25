# Browser Navigator

A browser history navigation system that supports visiting pages, moving backward, and moving forward through browsing history.

---

## Problem

Implement the `BrowserNavigator` class with the following functionality:

- Start from a homepage.
- Navigate to new pages.
- Move backward through history.
- Move forward through history.
- Clear forward history when a new page is visited after going back.

The navigator should simulate the behavior of a real web browser.

---

## Methods

### Constructor

```go
Constructor(startPage string)
```

Initializes the browser with the starting page.

---

### hopTo

```go
hopTo(page string)
```

Navigates to a new page.

Behavior:
- Saves the current page into the back history.
- Updates the current page.
- Clears the forward history.

---

### backtrack

```go
backtrack(steps int) string
```

Moves backward up to `steps` times.

Behavior:
- If there are fewer available pages than requested, it moves back as far as possible.
- Returns the current page after navigation.

---

### leapForward

```go
leapForward(steps int) string
```

Moves forward up to `steps` times.

Behavior:
- If there are fewer available pages than requested, it moves forward as far as possible.
- Returns the current page after navigation.

---

# Solution Explanation

The solution uses two stacks:

- `backStack`
- `forwardStack`

And one variable:

- `current`

---

## Core Idea

The current page is stored separately.

### Visiting a New Page

When navigating to a new page:

1. Push the current page into `backStack`
2. Update the current page
3. Clear `forwardStack`

This simulates how real browsers remove forward history after opening a new page.

---

### Moving Backward

When calling `backtrack`:

1. Push the current page into `forwardStack`
2. Pop the latest page from `backStack`
3. Set it as the current page

This operation is repeated until:
- the requested steps are completed, or
- there is no more history available.

---

### Moving Forward

When calling `leapForward`:

1. Push the current page into `backStack`
2. Pop the latest page from `forwardStack`
3. Set it as the current page

This operation is repeated until:
- the requested steps are completed, or
- there is no forward history available.

---

# Data Structures

```go
type BrowserNavigator struct {
	current      string
	backStack    []string
	forwardStack []string
}
```

---

# Complexity Analysis

## Time Complexity

| Operation | Complexity |
|---|---|
| hopTo | O(1) |
| backtrack | O(k) |
| leapForward | O(k) |

Where `k` is the number of navigation steps.

---

## Space Complexity

```text
O(n)
```

Where `n` is the number of visited pages stored in history.

---

# Example

```text
home -> page1 -> page2 -> page3
```

After:

```text
backtrack(2)
```

Current page becomes:

```text
page1
```

Then:

```text
leapForward(1)
```

Current page becomes:

```text
page2
```

Then:

```text
hopTo(page4)
```

Forward history is cleared.

Final history:

```text
home -> page1 -> page2 -> page4
```

---

# Key Concepts

- Stack
- History Navigation
- State Management
- Simulation
- In-Memory Data Structures
```
