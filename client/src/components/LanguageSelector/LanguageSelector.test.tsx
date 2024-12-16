import { composeStory } from "@storybook/react";
import { render, screen, waitFor } from "@testing-library/react";
import Meta, { DefaultLanguageSelector } from "./LanguageSelector.stories";

describe("Language selector", () => {
  beforeEach(() => {
    const ComposedLanguageSelector = composeStory(
      DefaultLanguageSelector,
      Meta
    );
    render(<ComposedLanguageSelector />);
  });

  it("Should display language selector trigger", async () => {
    await waitFor(() =>
      expect(screen.getByTestId("language-selector-trigger")).toBeVisible()
    );
  });
});
