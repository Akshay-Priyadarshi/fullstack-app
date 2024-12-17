import type { Meta, StoryObj } from "@storybook/react";
import { Typography } from "./Typography";
import { withRouter } from "@decorators";

const meta: Meta<typeof Typography> = {
  component: Typography,
  decorators: [withRouter],
  args: {
    children: "Fullstack App"
  }
};
export default meta;

type Story = StoryObj<typeof Typography>;

export const DefaultTypography: Story = {};

export const Heading1Typography: Story = {
  args: {
    variant: "h1"
  }
};

export const Heading2Typography: Story = {
  args: {
    variant: "h2"
  }
};

export const Heading3Typography: Story = {
  args: {
    variant: "h3"
  }
};

export const Heading4Typography: Story = {
  args: {
    variant: "h4"
  }
};

export const Heading5Typography: Story = {
  args: {
    variant: "h5"
  }
};

export const Heading6Typography: Story = {
  args: {
    variant: "h6"
  }
};

export const BodyTypography: Story = {
  args: {
    variant: "body"
  }
};
