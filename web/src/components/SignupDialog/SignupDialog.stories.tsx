import type { Meta, StoryObj } from "@storybook/react";
import { SignupDialog } from "./SignupDialog";
import { withRouter } from "@decorators";

const meta: Meta<typeof SignupDialog> = {
  component: SignupDialog,
  decorators: [withRouter]
};
export default meta;

type Story = StoryObj<typeof SignupDialog>;

export const DefaultSignupDialog: Story = {};
