import type { Meta, StoryObj } from "@storybook/react";
import { Counter } from "./Counter";

const meta: Meta<typeof Counter> = {
  component: Counter,
  argTypes: {
    initialCount: {
      name: "Initial Count",
      control: "number",
      description: "Initial count of the counter",
      type: "number"
    }
  }
};
export default meta;

type Story = StoryObj<typeof Counter>;

export const DefaultCounter: Story = {
  args: {
    initialCount: 100
  }
};
