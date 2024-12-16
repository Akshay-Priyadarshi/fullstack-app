import { composeStory } from "@storybook/react";
import { render, screen, waitFor } from "@testing-library/react";
import Meta, { DefaultCounter } from "./Counter.stories";

describe("Counter", () => {
  beforeEach(() => {
    const ComposedCounter = composeStory(DefaultCounter, Meta);
    render(<ComposedCounter />);
  });

  it("should display buttons", async () => {
    await waitFor(() => expect(screen.getByText("Increase")).toBeVisible());
    await waitFor(() => expect(screen.getByText("Reset")).toBeVisible());
    await waitFor(() => expect(screen.getByText("Decrease")).toBeVisible());
  });

  it("should display count", async () => {
    await waitFor(() =>
      expect(screen.getByTestId("counter-count")).toBeVisible()
    );
  });

  it("should increment count", async () => {
    const incrementButton = screen.getByText("Increase");
    const counterDisplay = screen.getByTestId("counter-count");

    const initialCount = Number(counterDisplay.textContent);
    const expectedCountAfterIncrement = initialCount + 1;

    incrementButton.click();

    await waitFor(() =>
      expect(counterDisplay).toHaveTextContent(
        String(expectedCountAfterIncrement)
      )
    );
  });

  it("should decrement count", async () => {
    const decrementButton = screen.getByText("Increase");
    const counterDisplay = screen.getByTestId("counter-count");

    const initialCount = Number(counterDisplay.textContent);
    const expectedCountAfterIncrement = initialCount + 1;

    decrementButton.click();

    await waitFor(() =>
      expect(counterDisplay).toHaveTextContent(
        String(expectedCountAfterIncrement)
      )
    );
  });
});
