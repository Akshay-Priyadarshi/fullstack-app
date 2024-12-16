import type { Meta, StoryObj } from "@storybook/react";
import { LoginDialog } from "./LoginDialog";
import { withRouter } from "@decorators";

const meta: Meta<typeof LoginDialog> = {
  component: LoginDialog,
  decorators: [withRouter]
};
export default meta;

type Story = StoryObj<typeof LoginDialog>;

export const DefaultLoginDialog: Story = {};
