import type { Meta, StoryObj } from "@storybook/react";
import { Menubar } from "./Menubar";
import { withRouter } from "@decorators";
const meta: Meta<typeof Menubar> = {
  component: Menubar,
  decorators: [withRouter]
};
export default meta;

type Story = StoryObj<typeof Menubar>;

export const DefaultMenubar: Story = {
  args: {
    menuLinks: [
      { label: "About", path: "/about" },
      { label: "Dashboard", path: "/user/dashboard" }
    ]
  }
};
