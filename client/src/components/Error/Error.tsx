import React from "react";

export interface IErrorProps {
  text?: string;
}

export const Error: React.FC<IErrorProps> = ({ text }) => {
  return <div>{text ?? "Something weird happened"}</div>;
};
