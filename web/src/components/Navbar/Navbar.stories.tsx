import type { Meta, StoryObj } from "@storybook/react";
import { Navbar } from "./Navbar";
import { withRouter } from "@decorators";

const meta: Meta<typeof Navbar> = {
  component: Navbar,
  decorators: [withRouter]
};
export default meta;

type Story = StoryObj<typeof Navbar>;

export const DefaultNavbar: Story = {};
